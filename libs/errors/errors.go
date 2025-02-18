package errors

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const ()

func ConfigError(isJsonFormat bool, stackTrace bool) {
	config = &errorConfig{
		IsJsonFormat: isJsonFormat,
		StackTrace:   stackTrace,
	}
}

func newStackErr(err error, data map[string]interface{}) error {
	if err == nil {
		return nil
	}

	if err, ok := err.(*stackErr); ok {
		return err
	}

	return &stackErr{
		err:    err,
		frames: getFrames(2),
		data:   data,
	}
}
func NewStackErr(err error) error {
	return newStackErr(err, nil)
}

func NewStackErrWithFields(err error, data map[string]interface{}) error {
	return newStackErr(err, data)
}

func UnwrapStackErr(err error) error {
	if err, ok := err.(*stackErr); ok {
		return err.Err()
	}
	return err
}

func IsStackErr(err error) StackError {
	if err, ok := err.(StackError); ok {
		return err
	}
	return nil
}

var (
	config = &errorConfig{
		IsJsonFormat: false,
		StackTrace:   false,
	}
)

type stackErr struct {
	frames []runtime.Frame
	err    error
	data   map[string]interface{}
}

func (e *stackErr) LogError(logger *logrus.Entry) {
	if config.IsJsonFormat {
		errCaller := e.ErrWithCallerJson()
		if errCaller == nil {
			return
		}
		errField := map[string]interface{}{
			"file":       errCaller.File,
			"line":       errCaller.Line,
			"function":   errCaller.Function,
			"stackTrace": e.ErrStackTraceJson().StackTrace,
		}

		payload := make(map[string]interface{})
		if e.data != nil {
			payload = e.data
		}
		payload[logrus.ErrorKey] = errField
		logger.WithFields(payload).Error("error: ", e.err)
		return
	}

	if config.StackTrace {
		logger.Error(e.ErrStackTraceJson())
		return
	}
	logger.Error(e.ErrWithCaller())
}

// Get Error in stack
func (e *stackErr) Err() error {
	return e.err
}

// Get Error in stack
func (e *stackErr) Error() string {
	return e.err.Error()
}

// Get Error With Caller Non-Json
func (e *stackErr) ErrWithCaller() string {
	if len(e.frames) < 1 {
		return ""
	}
	frame := e.frames[0]
	msg := fmt.Sprintf("[%s] [%s:%d] %s", frame.Function, frame.File, frame.Line, e.Error())
	return msg
}

// Get Error With Caller Json
func (e *stackErr) ErrWithCallerJson() *errCaller {
	if len(e.frames) < 1 {
		return nil
	}

	frame := e.frames[0]
	return &errCaller{
		eachTrace: eachTrace{
			Function: frame.Function,
			File:     frame.File,
			Line:     frame.Line,
		},
		Error: e.Err().Error(),
	}
}

// Get Error With Caller JsonString
func (e *stackErr) ErrWithCallerJsonString() string {
	errJson := e.ErrWithCallerJson()
	if errJson != nil {
		return ""
	}

	jsonString, err := json.Marshal(errJson)
	if err != nil {
		return ""
	}
	return string(jsonString)
}

// Get Stack Trace Non-Json
func (e *stackErr) ErrStackTrace() string {
	stackTrace := []string{}
	for i, f := range e.frames {
		if i == 0 {
			stackTrace = append(stackTrace, fmt.Sprintf("\n%v\n", caller(f)))
		} else {
			stackTrace = append(stackTrace, fmt.Sprintf("%v\n", caller(f)))
		}
	}
	return fmt.Sprintf("%v %v", stackTrace, e.Err().Error())
}

// Get Stack Trace Json
func (e *stackErr) ErrStackTraceJson() *errStackTrace {
	stacks := []eachTrace{}
	for _, f := range e.frames {
		stacks = append(stacks, eachTrace{
			Function: f.Function,
			File:     f.File,
			Line:     f.Line,
		})
	}

	return &errStackTrace{
		StackTrace: stacks,
		Error:      e.Err().Error(),
	}
}

// Get Stack Trace JsonString
func (e *stackErr) ErrStackTraceJsonString() string {
	errStackTrace := e.ErrStackTraceJson()
	jsonString, err := json.Marshal(errStackTrace)
	if err != nil {
		return ""
	}
	return string(jsonString)
}

func getFrames(skipFrames int) []runtime.Frame {
	frameList := []runtime.Frame{}
	targetFrameIndex := skipFrames + 2
	programCounters := make([]uintptr, 10)
	n := runtime.Callers(0, programCounters)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return frameList
	}

	programCounters = programCounters[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(programCounters)
	frameIndex := 0
	for {
		frame, more := frames.Next()

		if frameIndex < targetFrameIndex {
			// Skip Process.
			frameIndex++
			continue
		}

		if rootPath, err := os.Getwd(); err == nil {
			frame.File = strings.Replace(frame.File, rootPath, "", 1)
		}

		s := strings.Split(frame.Function, "/")
		if len(s) > 0 {
			frame.Function = s[len(s)-1]
		}

		frameList = append(frameList, frame)
		if strings.Contains(frame.File, "handlers") || !more {
			break
		}

		frameIndex++
	}

	return frameList
}
func caller(frame runtime.Frame) string {
	return fmt.Sprintf("at %s() \n\t\t %s:%d", frame.Function, frame.File, frame.Line)
}
