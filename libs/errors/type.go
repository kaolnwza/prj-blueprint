package errors

import "github.com/sirupsen/logrus"

type (
	errorConfig struct {
		IsJsonFormat bool
		StackTrace   bool
	}

	StackError interface {
		LogError(logger *logrus.Entry)
		ErrWithCaller() string
		ErrWithCallerJsonString() string
		ErrStackTrace() string
		ErrStackTraceJsonString() string
		Err() error
	}

	errCaller struct {
		eachTrace
		Error string `json:"error"`
	}

	errStackTrace struct {
		Error      string      `json:"error"`
		StackTrace []eachTrace `json:"stackTrace"`
	}

	eachTrace struct {
		Function string `json:"function"`
		File     string `json:"file"`
		Line     int    `json:"line"`
	}
)
