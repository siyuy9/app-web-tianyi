package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	v1 "gitlab.com/kongrentian-groups/golang/tianyi/backend/api/route/v1"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/common"
)

func SetupRoutes(api fiber.Router) {
	// api routes
	v1.SetRoutes(api.Group("/v1"))
	// redirect for requests from gitlab-runners
	api.Use("/v4", func(context *fiber.Ctx) error {
		return context.Redirect("/v1")
	})
	// catch all so you don't get html files on invalid api requests
	api.Use(MissingRouteHandler)
}

func MissingRouteHandler(context *fiber.Ctx) error {
	return common.BadRequest(context, fiber.StatusBadRequest, fmt.Sprintf(
		"Route '%s %s' does not exist",
		context.Method(),
		context.OriginalURL(),
	))
}
