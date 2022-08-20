package pkg

import "github.com/go-playground/validator/v10"

var validate *validator.Validate = validator.New()

func ValidateStruct(data interface{}) error {
	return NewErrorBadRequest(validate.Struct(data))
}
