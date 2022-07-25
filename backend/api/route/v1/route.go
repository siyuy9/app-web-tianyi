package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/api/route/v1/auth"
)

func SetRoutes(group fiber.Router) {
	auth.SetRoutes(group.Group("/auth"))
}
