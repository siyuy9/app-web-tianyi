package controller

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

// global fiber error handler
func Error(ctx *fiber.Ctx, err error) error {
	// if it is not a custom error for some reason, make it one
	errNormalized, ok := err.(pkgError.Error)
	if !ok {
		errNormalized = pkgError.New(
			fmt.Errorf("an unexpected error has occured: %w", err),
		).(pkgError.Error)
	}

	// send error response
	errSerialization := presenter.Error(ctx, errNormalized)
	if errSerialization == nil {
		return nil
	}
	// in case a serialization error occurs
	return presenter.Error(
		ctx,
		pkgError.NewWithCode(
			fmt.Errorf(
				"could not serialize error %+v: %w", err, errSerialization,
			),
			http.StatusInternalServerError,
		).(pkgError.Error),
	)
}
