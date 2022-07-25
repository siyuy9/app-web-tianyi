package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(group fiber.Router) {
	group.Post("/login", loginHandler)
}
