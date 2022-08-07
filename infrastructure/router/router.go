package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"gitlab.com/kongrentian-group/tianyi/api/controller"
)

func New(config *fiber.Config) *fiber.App {
	config.ErrorHandler = errorHandler
	return fiber.New(*config)
}

/*
setup routes

@title Fiber Example API
@version 1.0
@description This is a sample swagger for Fiber
@termsOfService http://swagger.io/terms/
@contact.name API Support
@contact.email fiber@swagger.io
@license.name Apache 2.0
@license.url http://www.apache.org/licenses/LICENSE-2.0.html
@host localhost:8080
@BasePath /
*/
func Setup(router *fiber.App, app *controller.AppController) {
	// use global middlewares
	// https://docs.gofiber.io/api/middleware
	router.Use(
		cors.New(),
		compress.New(),
		etag.New(),
		csrf.New(),
		recover.New(),
		logger.New(),
		// limiter.New(),
	)

	// api routes
	// apiGroup := router.Group("/api")

	//router.Use("/v4", func(context *fiber.Ctx) error {
	//	return context.Redirect("/api/v1")
	//})

	// frontend controller
	//
	// returns embedded files from ../infrastructure/ui/web2/dist
	// (the folder is created and populated by 'yarn build')
	//
	// it is a Single Page Application, routing is done by Javascript, that's
	// why it has to return index.html on missing files
	router.Get("*", app.Frontend.Serve)

	// catch all, just in case
	// router.Use(api.MissingRouteHandler)
}
