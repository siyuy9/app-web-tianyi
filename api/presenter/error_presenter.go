package presenter

import (
	"log"

	"github.com/gofiber/fiber/v2"
	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

func Error(context *fiber.Ctx, err pkgError.Error) error {
	log.Println("ERROR", err)
	return context.Status(err.StatusCode()).JSON(&ResponseError{Error: err})
}
