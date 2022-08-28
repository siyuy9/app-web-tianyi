package pkgError

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

const (
	// default number of frames to backtrack
	FrameDefault = 2
)

// error wrapper
type Error interface {
	// interface for errors.Unwrap
	Unwrap() error
	StatusCode() int
	// error interface
	// generates stacktrace of itself and all wrapped errors when called
	Error() string
	HasFrame
	// returns regular error string
	ErrorWithoutStacktrace() string
}

type HasFrame interface {
	// stacktrace is generated using that frame
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
	frameUse := FrameDefault
	if len(frame) != 0 {
		frameUse = frame[0]
	}
	return &customError{code, err, getFrame(frameUse)}
}

func (customError *customError) Error() string {
	return Stacktrace(customError)
}

func (customError *customError) ErrorWithoutStacktrace() string {
	assertedError, ok := customError.err.(Error)
	if !ok {
		return customError.err.Error()
	}
	return assertedError.ErrorWithoutStacktrace()
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
	withFrame, ok := err.(HasFrame)
	var frame *runtime.Frame
	if ok {
		frame = withFrame.Frame()
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
