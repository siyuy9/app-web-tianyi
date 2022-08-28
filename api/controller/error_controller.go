package controller

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

// global fiber error handler
func Error(context *fiber.Ctx, err error) error {
	// if it is not a custom error for some reason, make it one
	errNormalized, ok := err.(pkgError.Error)
	if !ok {
		errNormalized = pkgError.New(fmt.Errorf(
			"an unexpected error has occured: %w", err,
		)).(pkgError.Error)
	}

	// send error response
	err = presenter.Error(context, errNormalized)
	if err == nil {
		return nil
	}
	// in case a serialization error occurs
	return presenter.Error(
		context,
		pkgError.NewWithCode(
			fmt.Errorf(
				"could not serialize error %#v: %w", errNormalized, err,
			),
			http.StatusInternalServerError,
		).(pkgError.Error),
	)
}
