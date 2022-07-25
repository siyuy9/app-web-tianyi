package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/api/middleware"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/api/route"
	"gitlab.com/kongrentian-groups/golang/tianyi/frontend"
)

func SetRoutes(app *fiber.App) {
	route.SetRoutes(app.Group("/api", middleware.RequireLoginUnauthorized))

	app.Use("/dashboard", middleware.RequireLoginUnauthorized, monitor.New())

	app.Use(
		"/favicon.ico",
		filesystem.New(filesystem.Config{
			Root:       http.FS(frontend.Favicon),
			PathPrefix: "dist",
			Index:      "favicon.ico",
		}))
	app.Use(
		"/static",
		filesystem.New(filesystem.Config{
			Root: http.FS(frontend.Static),
			// `embed` embeds the whole folder, we have to add prefix
			PathPrefix: "dist/static",
		}))
	app.Use(
		"/",
		middleware.RequireLoginRedirect,
		filesystem.New(filesystem.Config{
			Root:       http.FS(frontend.Index),
			PathPrefix: "dist",
		}))

	app.All("*", func(context *fiber.Ctx) error {
		errorMessage := fmt.Sprintf(
			"Route '%s' does not exist",
			context.OriginalURL(),
		)

		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})
}
