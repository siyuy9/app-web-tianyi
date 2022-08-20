package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
)

// parse request body into target, then validate it
func parse(context *fiber.Ctx, target interface{}) error {
	if err := context.BodyParser(target); err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	return pkg.ValidateStruct(target)
}
