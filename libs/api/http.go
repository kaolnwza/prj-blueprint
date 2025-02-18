package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/libs/errors"
	"github.com/kaolnwza/proj-blueprint/pkg/logger"
)

const (
	keyLogResponse = "response"
	keyLogRequest  = "request"
)

type httpClient struct {
	logConf config.LogConfig
	// httpConf config.HttpConfig
}

type HttpClient interface {
	NewRequest(ctx context.Context, cli http.Client, req *http.Request, dest any) error
	BuildEndpoint(apiUrl ApiUrl) (*url.URL, error)
}

func New(logConf config.LogConfig, httpConf config.HttpConfig) HttpClient {
	return httpClient{
		logConf: logConf,
		// httpConf: httpConf,
	}
}

func (c httpClient) NewRequest(ctx context.Context, cli http.Client, req *http.Request, dest any) error {
	logHeader := req.Header.Clone()
	logHeader.Del("Authorization")
	log := logger.GetByContext(ctx)
	var reqBody []byte
	if req.Body != nil {
		reqBody, err := io.ReadAll(req.Body)
		if err == nil {
			req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}
	}

	loggerUrl := urlMessageFormat(req.Method, req.URL.String())
	payload := map[string]interface{}{
		"header": logHeader,
		"url":    loggerUrl,
	}

	c.addResponseLog(payload, reqBody)
	log.WithFields(payload).Infof("Sending request to %v %v", req.Method, req.URL.String())

	// resp, err := c.client.Do(req)
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	responseData, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseData, dest); err != nil {
		payload := map[string]interface{}{
			"url":    loggerUrl,
			"header": logHeader,
			"status": resp.Status,
		}

		c.addBodyLogOnErr(payload, keyLogRequest, reqBody)
		c.addBodyLogOnErr(payload, keyLogResponse, responseData)

		log.WithFields(payload).Errorf("failed to read responseData, err = %v", err)

		return err
	}

	if isErrorStatus(resp.StatusCode) {
		msg := fmt.Sprintf("%v %v failed with %v", req.Method, req.URL.String(), resp.Status)
		payload := map[string]interface{}{
			"url":    loggerUrl,
			"header": logHeader,
			"status": resp.Status,
		}

		c.addBodyLogOnErr(payload, keyLogRequest, reqBody)
		c.addBodyLogOnErr(payload, keyLogResponse, responseData)

		log.WithFields(payload).Error(msg)

		return err
	}

	payload = map[string]interface{}{
		"status": resp.Status,
		"url":    loggerUrl,
	}
	c.addResponseLog(payload, responseData)

	log.WithFields(payload).Infof("Completed request %v %v successfully", req.Method, req.URL.String())
	return errors.NewStackErr(err)
}

func urlMessageFormat(httpMethod, url string) string {
	return fmt.Sprintf("[%s] %s", httpMethod, url)
}

func (c httpClient) addResponseLog(m map[string]interface{}, b []byte) {
	if c.logConf.ShowBody {
		m[keyLogResponse] = byteToString(b)
	}
}

func (c httpClient) addBodyLogOnErr(m map[string]interface{}, k string, b []byte) {
	if c.logConf.ShowBodyOnError {
		m[k] = byteToString(b)
	}
}

func byteToString(b []byte) (s string) {
	if err := json.Unmarshal([]byte(b), &s); err != nil {
		return fmt.Sprintf("fail to decode byte into string, err: %v, byte: %v", err, b)
	}

	return s
}

func isErrorStatus(statusCode int) bool {
	return statusCode >= 400 && statusCode < 502
}

func SerializeObject(objectBody interface{}) (io.ReadCloser, error) {
	body, err := json.Marshal(objectBody)
	return io.NopCloser(bytes.NewReader(body)), err
}
