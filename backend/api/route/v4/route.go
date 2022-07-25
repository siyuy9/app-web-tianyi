package v4

// API for gitlab-runner

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(group fiber.Router) {
	group.Get("/runners", runnersGetHandler)
}
