package backend

import (
	"github.com/gofiber/fiber/v2"

	"gitlab.com/kongrentian-groups/golang/tianyi/backend/api"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/api/middleware"
	"gitlab.com/kongrentian-groups/golang/tianyi/frontend"
)

func setRoutes(app *fiber.App) {
	// api routes
	api.SetRoutes(app.Group("/api", middleware.RequireLoginUnauthorized))

	// frontend handler
	// returns embedded files from ../frontend/dist
	// (the module has to be there, because `embed` does not allow to use
	// parent directories)
	// (the folder is created and populated by 'yarn build')
	// it is a Single Page Application, routing is done by Javascript, that's
	// why if there are no files for the request, e.g. /login,
	// it just returns ../frontend/dist/index.html
	app.Get("*", frontend.DistHandler)

	// catch all, just in case
	app.Use("*", api.MissingRouteHandler)
}
