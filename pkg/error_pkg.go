package pkg

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// error wrapper
type Error struct {
	Type   string            `json:"type" example:"*fiber.Error"`
	Status string            `json:"status" example:"fail"`
	Code   int               `json:"code" example:"500"`
	Errors map[string]string `json:"errors"`
}

func (errorStruct *Error) Error() string {
	return fmt.Sprintf("{ Errors: %v }", errorStruct.Errors)
}

// creates a new error
// default status code - http.StatusInternalServerError
// if input is nil, return nil
// if input is a string, make a new Error
// if input is *fiber.Error, convert it to Error
// if input is map[string]error, convert it to Error (if all errors are nil, return nil)
// if input is *Error, return it
// if input is a validator.ValidationErrors, parse it
// if input is an error, make a new Error
func NewError(err interface{}, code ...int) error {
	if err == nil {
		return nil
	}
	var statusCode int
	if len(code) != 0 {
		statusCode = code[0]
	} else {
		statusCode = http.StatusInternalServerError
	}
	errorType := reflect.TypeOf(err).String()
	errorsMap := make(map[string]string)
	switch errorAsserted := err.(type) {
	case string:
		return NewError(errors.New(errorAsserted), statusCode)
	case *fiber.Error:
		return NewError(errorAsserted.Message, errorAsserted.Code)
	case map[string]error:
		for key, value := range errorAsserted {
			if value == nil {
				continue
			}
			errorsMap[key] = value.Error()
		}
	case *Error:
		return errorAsserted
	case validator.ValidationErrors:
		statusCode = http.StatusBadRequest
		for _, value := range errorAsserted {
			errorsMap[value.Field()] = value.Tag()
		}
	case error:
		errorsMap["body"] = errorAsserted.Error()

	default:
		log.Panicf(fmt.Sprintf("invalid '%s' error: %s", errorType, err))
	}

	if len(errorsMap) == 0 {
		return nil
	}
	return &Error{errorType, "fail", statusCode, errorsMap}
}

func NewErrorBadRequest(err interface{}) error {
	return NewError(err, http.StatusBadRequest)
}

func NewErrorForbidden(err interface{}) error {
	return NewError(err, http.StatusForbidden)
}

func NewErrorUnauthorized(err interface{}) error {
	return NewError(err, http.StatusUnauthorized)
}
