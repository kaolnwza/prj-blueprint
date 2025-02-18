package logger

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/libs/consts"
	"github.com/sirupsen/logrus"
)

type logJson struct {
	ID        string        `json:"id"`
	Level     string        `json:"severity"`
	Timestamp string        `json:"timestamp"`
	Message   string        `json:"message"`
	Version   string        `json:"version"`
	Data      logrus.Fields `json:"data,omitempty"`
	Error     logrus.Fields `json:"error,omitempty"`
}

type customFormatter struct {
	LogFormat string
	prefixLog string
}

func (f *customFormatter) AppendFormat(additional string) *customFormatter {
	f.LogFormat = f.LogFormat + " " + additional
	return f
}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if f.LogFormat == "" {
		return (&logrus.TextFormatter{}).Format(entry)
	}

	if strings.EqualFold(f.LogFormat, formatJSONKey) {
		return f.jsonFormat(entry)
	}
	return f.textFormat(entry)
}

func (f *customFormatter) jsonFormat(entry *logrus.Entry) ([]byte, error) {
	id := ""
	if val, exist := entry.Data[consts.RequestIdKey]; exist {
		id = val.(string)
		delete(entry.Data, consts.RequestIdKey)
	}

	err := map[string]interface{}{}
	if val, exist := entry.Data[logrus.ErrorKey]; exist {
		if s, ok := val.(map[string]interface{}); ok {
			err = s
		}
		delete(entry.Data, logrus.ErrorKey)
	}

	if entry.Level <= logrus.FatalLevel {
		entry.Level = logrus.FatalLevel
	}

	logMsg := fmt.Sprintf("[%s] %s", id, entry.Message)
	format := logJson{
		ID:        id,
		Level:     strings.ToUpper(entry.Level.String()),
		Timestamp: entry.Time.Format(time.RFC3339),
		Version:   config.VERSION,
		Message:   logMsg,
		Data:      entry.Data,
	}

	if err != nil || len(err) == 0 {
		format.Error = err
	}

	return withNewLine(json.Marshal(&format))
}

func (f *customFormatter) textFormat(entry *logrus.Entry) ([]byte, error) {
	logOutput := f.LogFormat
	logOutput = strings.Replace(logOutput, "<lvl>", fmt.Sprintf("%v", strings.ToUpper(entry.Level.String())), 1)
	logOutput = strings.Replace(logOutput, "<time>", entry.Time.Format(time.RFC3339), 1)

	if val, exist := entry.Data[consts.RequestIdKey]; exist {
		logOutput = strings.Replace(logOutput, "<requestID>", val.(string), 1)
	} else {
		logOutput = strings.Replace(logOutput, " [<requestID>]", "", 1)
	}

	if f.prefixLog != "" {
		logOutput = fmt.Sprintf("%s %s", f.prefixLog, logOutput)
	}

	if entry.Message != "" {
		logOutput = fmt.Sprintf("%s %s", logOutput, entry.Message)
	}

	for k, val := range entry.Data {
		if k == consts.RequestIdKey {
			continue
		}
		if !strings.Contains(logOutput, "<"+k+">") {
			logOutput = fmt.Sprintf("%v %s=\"%v\"", logOutput, k, val)
			continue
		}
		switch v := val.(type) {
		case string:
			logOutput = strings.Replace(logOutput, "<"+k+">", v, 1)
		case int:
			s := strconv.Itoa(v)
			logOutput = strings.Replace(logOutput, "<"+k+">", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			logOutput = strings.Replace(logOutput, "<"+k+">", s, 1)
		}
	}
	return withNewLine([]byte(logOutput), nil)
}
