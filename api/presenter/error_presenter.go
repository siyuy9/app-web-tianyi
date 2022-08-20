package presenter

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
)

// global fiber error handler
func Error(context *fiber.Ctx, err error) error {
	// if it is not a custom error, make it one
	errNormalized, ok := err.(*pkg.Error)
	if !ok {
		errNormalized = pkg.NewError(
			err, http.StatusInternalServerError,
		).(*pkg.Error)
	}

	// send error response
	err = context.Status(errNormalized.Code).JSON(errNormalized)
	if err == nil {
		return nil
	}
	// in case a serialization error occurs
	// string representation info - https://stackoverflow.com/a/16332828
	newError := pkg.NewError(
		fmt.Errorf(
			"could not serialize the error %#v because of '%s'",
			errNormalized,
			err.Error(),
		),
		http.StatusInternalServerError,
	)
	return context.Status(http.StatusInternalServerError).JSON(newError)
}
