package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/api/route/v1/auth"
	_ "gitlab.com/kongrentian-groups/golang/tianyi/docs"
)

func SetRoutes(group fiber.Router) {
	// auth routes
	auth.SetRoutes(group.Group("/auth"))

	group.Get("/runners", runnersGetHandler)
	group.Get("/swagger/*", swagger.HandlerDefault)
}
