package presenter

import (
	"github.com/gofiber/fiber/v2"
	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

func Error(context *fiber.Ctx, err pkgError.Error) error {
	return context.Status(err.StatusCode()).JSON(
		&ResponseError{Error: err.Error()},
	)
}
