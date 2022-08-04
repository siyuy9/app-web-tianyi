package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/common"
)

func requireLogin(
	context *fiber.Ctx,
	unauthorized func(context *fiber.Ctx) error,
) error {
	currentSession, err := common.App.SessionStore.Get(context)
	if err != nil {
		return err
	}
	user := currentSession.Get("User")
	defer currentSession.Save()

	if user == nil {
		// This request is from a user that is not logged in.
		return unauthorized(context)
	}

	// If we got this far, the request is from a logged-in user.
	// Continue on to other middleware or routes.
	return context.Next()
}

func RequireLoginRedirect(context *fiber.Ctx) error {
	return requireLogin(context, func(context *fiber.Ctx) error {
		return context.Redirect("/login")
	})
}

func RequireLoginUnauthorized(context *fiber.Ctx) error {
	return requireLogin(context, func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "unauthorized",
		})
	})
}
