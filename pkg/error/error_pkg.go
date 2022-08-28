package pkgError

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

// error wrapper
type Error interface {
	Unwrap() error
	StatusCode() int
	ErrorWithFrame
	ErrorWithoutStacktrace() string
}

const (
	FrameDefault = 2
)

type ErrorWithFrame interface {
	error
	Frame() *runtime.Frame
}

type customError struct {
	statusCode int
	err        error
	frame      *runtime.Frame
}

func New(err error, frame ...int) error {
	return NewWithCode(err, http.StatusInternalServerError, frame...)
}

func NewWithCode(err error, code int, frame ...int) error {
	if err == nil {
		return nil
	}
	var frameUse int
	if len(frame) != 0 {
		frameUse = frame[0]
	} else {
		frameUse = FrameDefault
	}
	return &customError{code, err, getFrame(frameUse)}
}

func (customError *customError) Error() string {
	return Stacktrace(customError)
}

func (customError *customError) ErrorWithoutStacktrace() string {
	return customError.err.Error()
}

func (customError *customError) Unwrap() error {
	return errors.Unwrap(customError.err)
}

func (customError *customError) Frame() *runtime.Frame {
	return customError.frame
}

func (customError *customError) StatusCode() int {
	return customError.statusCode
}

// return stacktrace
func Stacktrace(err error) string {
	var buffer bytes.Buffer
	for err := err; ; {
		if err == nil {
			return buffer.String()
		}
		putStacktrace(err, &buffer)
		err = errors.Unwrap(err)
	}
}

func putStacktrace(err error, buffer *bytes.Buffer) {
	fmt.Fprintf(buffer, "%v", err)
	errorWithFrame, ok := err.(ErrorWithFrame)
	var frame *runtime.Frame
	if ok {
		frame = errorWithFrame.Frame()
	}
	if frame != nil {
		fmt.Fprintf(buffer, "\n    %s:%d", frame.File, frame.Line)
	}
	fmt.Fprint(buffer, "\n")
}

func getFrame(calldepth int) *runtime.Frame {
	programCounter, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return nil
	}
	frame := &runtime.Frame{
		PC:   programCounter,
		File: file,
		Line: line,
	}
	if function := runtime.FuncForPC(programCounter); function != nil {
		frame.Func = function
		frame.Function = function.Name()
		frame.Entry = function.Entry()
	}
	return frame
}
