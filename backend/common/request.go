package common

import "github.com/gofiber/fiber/v2"

func BadRequest(context *fiber.Ctx, status int, message string) error {
	return context.Status(status).JSON(&fiber.Map{
		"status":  "fail",
		"message": message,
	})
}
