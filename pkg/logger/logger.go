package logger

import (
	"context"
	"net/http"
	"os"

	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/libs/constants"
	"github.com/kaolnwza/proj-blueprint/libs/utils"
	"github.com/sirupsen/logrus"
)

var (
	logger   *logrus.Logger
	logLevel logrus.Level
	logConf  config.LogConfig
)

const (
	correlationKey    = "X-Correlation-ID" // Unique request key in NEXT
	formatJSONKey     = "JSON"
	prefixLog         = "==="
	prefixRequestLog  = ">>>"
	prefixResponseLog = "<<<"
	logFormat         = "[<time>] [<lvl>] [<requestID>]"
)

func Setup(conf config.LogConfig) {
	logLevel = logrus.DebugLevel
	logger = newLogger(prefixLog, logFormat)

	// configErrorLog(conf)
	updateLogLevel(conf.Level)

	logConf = conf
}

// func configErrorLog(conf config.LogConfig) {
// 	errors.ConfigError(conf.JsonFormat, conf.Error.StackTrace)
// }

func updateLogLevel(level string) {
	switch level {
	case "debug":
		logLevel = logrus.DebugLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	case "fatal":
		logLevel = logrus.FatalLevel
	case "info":
		logLevel = logrus.InfoLevel
	case "trace":
		logLevel = logrus.TraceLevel
	case "panic":
		logLevel = logrus.PanicLevel
	default:
		logLevel = logrus.DebugLevel
	}
	logger.SetLevel(logLevel)
	logger.SetFormatter(&customFormatter{LogFormat: getLogFormat(), prefixLog: prefixLog})
}

// call default logger from logrus with config format.
func New() *logrus.Entry {
	return logrus.NewEntry(logger)
}

type Entry interface {
	ID() string
	LogEntry() *logrus.Entry
	LogRequest()
	LogResponse(response string, status int)
}

type entry struct {
	*logrus.Entry
	id      string
	request *http.Request
}

func getLogFormat() string {
	if logConf.JsonFormat {
		return formatJSONKey
	}
	return logFormat
}

func newLogger(prefix, format string) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&customFormatter{LogFormat: getLogFormat(), prefixLog: prefix})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logLevel)
	return logger
}

func getLogEntry(id, prefix, format string) *logrus.Entry {
	logger := newLogger(prefix, format)
	return logrus.NewEntry(logger).WithField(constants.RequestIdKey, id)
}

func GetById(id string) *logrus.Entry {
	return getLogEntry(id, prefixLog, getLogFormat())
}

// this fn will include requestId on logging.
func GetByContext(ctx context.Context) *logrus.Entry {
	requestID, err := utils.GetContext[string](ctx, constants.CtxRequestId)
	if err != nil {
		GetById("").Errorf("Failed to get request ID from context, err = %v", err)
	}

	return GetById(requestID)
}
