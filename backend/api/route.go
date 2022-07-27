package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	v1 "gitlab.com/kongrentian-groups/golang/tianyi/backend/api/route/v1"
	v4 "gitlab.com/kongrentian-groups/golang/tianyi/backend/api/route/v4"
)

func SetRoutes(api fiber.Router) {
	v1.SetRoutes(api.Group("/v1"))
	v4.SetRoutes(api.Group("/v4"))
	api.All("/", MissingRouteHandler)
}

func MissingRouteHandler(context *fiber.Ctx) error {
	errorMessage := fmt.Sprintf(
		"Route '%s %s' does not exist",
		context.Method(),
		context.OriginalURL(),
	)
	return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"status":  "fail",
		"message": errorMessage,
	})
}
