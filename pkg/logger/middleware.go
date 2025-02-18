package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/kaolnwza/proj-blueprint/libs/constants"
	"github.com/sirupsen/logrus"
)

// NewRequestLogger will include requestID in logs. Recommended to use during request handling.
func NewRequestLogger(request *http.Request) Entry {
	requestID := request.Header.Get(constants.RequestIdKey)
	if requestID == "" {
		requestID = uuid.New().String()
		request.Header.Add(constants.RequestIdKey, requestID)
	}

	return &entry{
		Entry:   logrus.NewEntry(logger).WithField(constants.RequestIdKey, requestID),
		id:      requestID,
		request: request,
	}
}

func (e *entry) ID() string {
	return e.id
}

func (e *entry) LogEntry() *logrus.Entry {
	return e.Entry
}

func (e *entry) RequestEntry() *logrus.Entry {
	return getLogEntry(e.ID(), prefixRequestLog, getLogFormat())
}

func (e *entry) ResponseEntry() *logrus.Entry {
	return getLogEntry(e.ID(), prefixResponseLog, getLogFormat())
}

func (e *entry) LogRequest() {
	header, err := jsonHeader(e.request.Header)
	if err != nil {
		e.Errorf("Failed to read request headers, err = %v", err)
	}

	var reqBody []byte

	if e.request.Body != nil {
		reqBody, err = io.ReadAll(e.request.Body)
		if err != nil {
			e.Errorf("Failed to read request body, err = %v", err)
		}
	}
	e.request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	body, err := jsonBody(reqBody)
	if err != nil {
		e.Errorf("Failed to marshal request body, err = %v", err)
	}

	if logConf.JsonFormat {
		payload := map[string]interface{}{
			"url":    fmt.Sprintf("[%s] %s", e.request.Method, e.request.URL.String()),
			"header": toInterfaceHeader(string(header)),
		}

		if showBodyLogger(logConf.ShowBody, logConf.ShowBodyOnError, nil) {
			payload["body"] = toInterfaceBody(body)
		}

		e.RequestEntry().WithFields(payload).Infof("%s %v %v", prefixRequestLog, e.request.Method, e.request.URL)
		return
	}

	e.RequestEntry().Infof("%v %v header:\"%s\" body:\"%s\"",
		e.request.Method,
		e.request.URL,
		header,
		body,
	)
}

func (e *entry) LogResponse(response string, status int) {
	header, err := jsonHeader(e.request.Header)
	if err != nil {
		e.Errorf("Failed to read response headers, err = %v", err)
	}

	if logConf.JsonFormat {
		payload := map[string]interface{}{
			"url":    fmt.Sprintf("[%s] %s", e.request.Method, e.request.URL.String()),
			"header": toInterfaceHeader(string(header)),
			"status": status,
		}

		if showBodyLogger(logConf.ShowBody, logConf.ShowBodyOnError, &status) {
			payload["body"] = toInterface(response)
		}

		e.ResponseEntry().WithFields(payload).Infof("%s [%d] %v %v", prefixResponseLog, status, e.request.Method, e.request.URL)
		return
	}

	e.ResponseEntry().Infof("%v %v header:\"%s\" body:\"%s\" status:\"%v\"",
		e.request.Method,
		e.request.URL,
		header,
		response,
		status,
	)
}

func showBodyLogger(logEnabled, logEnabledOnErr bool, httpStatus *int) bool {
	if httpStatus != nil && *httpStatus != http.StatusOK && logEnabledOnErr {
		return true
	}

	if logEnabled {
		return true
	}

	return false
}

func jsonHeader(headers http.Header) ([]byte, error) {
	res := map[string]interface{}{}
	for k, v := range headers {
		if len(v) == 1 {
			res[k] = v[0]
		} else {
			res[k] = v
		}
	}
	if logConf.Pretty {
		return json.MarshalIndent(res, "", "  ")
	}
	return json.Marshal(res)
}

func jsonBody(response []byte) ([]byte, error) {
	if string(response) == "" {
		return response, nil
	}
	var res map[string]interface{}
	var resp []byte
	if err := json.Unmarshal(response, &res); err != nil {
		return resp, err
	}
	if logConf.Pretty {
		return json.MarshalIndent(res, "", "  ")
	}
	return json.Marshal(res)
}

// for remove \" in json string
func toInterface(s string) (i interface{}) {
	if err := json.Unmarshal([]byte(s), &i); err != nil {
		return s
	}
	return i
}

func withNewLine(b []byte, err error) ([]byte, error) {
	return []byte(string(b) + "\n"), err
}

func toInterfaceHeader(s string) (i interface{}) {
	i = toInterfaceNew(s)
	if i == nil {
		return s
	}
	return i
}

// for remove \" in json string
func toInterfaceNew(s string) (i map[string]interface{}) {
	if err := json.Unmarshal([]byte(s), &i); err != nil {
		return nil
	}
	delete(i, "Authorization")
	return i
}

func toInterfaceBody(b []byte) (i interface{}) {
	if err := json.Unmarshal(b, &i); err != nil {
		return b
	}
	return i
}
