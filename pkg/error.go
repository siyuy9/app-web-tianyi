package pkg

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// error wrapper
type Error struct {
	Errors map[string]interface{} `json:"errors"`
	Type   string                 `json:"type"`
	Status string                 `json:"status"`
	Code   int                    `json:"code"`
	err    error
}

func (errorStruct *Error) Error() string {
	return errorStruct.err.Error()
}

// creates a new error
// if input is nil, return nil
// if input is a string, make a new Error
// if input is an Error, return it
// if input is a validator.ValidationErrors, parse it
// if input is an error, make a new Error
func NewError(err interface{}, code int) *Error {
	if err == nil {
		return nil
	}
	errorType := reflect.TypeOf(err).String()
	errorsMap := make(map[string]interface{})
	switch errorAsserted := err.(type) {
	case string:
		return NewError(errors.New(errorAsserted), code)

	case Error:
		return &errorAsserted
	case validator.ValidationErrors:
		code = http.StatusBadRequest
		for _, value := range errorAsserted {
			errorsMap[value.Field()] = fmt.Sprintf("%v", value.Tag())
		}
	case error:
		errorsMap["body"] = errorAsserted.Error()

	default:
		log.Panicf(fmt.Sprintf("invalid '%s' error: %s", errorType, err))
	}

	result := &Error{
		Errors: errorsMap,
		Type:   errorType,
		Status: "fail",
		Code:   code,
		err:    err.(error),
	}
	return result
}

func NewBadRequestError(err interface{}) *Error {
	return NewError(err, http.StatusBadRequest)
}
