package backend

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/api"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/common"
	_ "gitlab.com/kongrentian-groups/golang/tianyi/docs"
	"gitlab.com/kongrentian-groups/golang/tianyi/frontend"
)

func Run() {
	// initialize stuff
	common.App.InitializeServer()
	// setup routes
	setupRoutes(common.App.Fiber)
	// listen in another goroutine
	go common.App.Listen()
	// channel for the interruption signal
	channel := make(chan os.Signal, 1)
	// when an interruption or termination signal is sent, shutdown
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	<-channel
	common.App.Shutdown(0)
}

/*
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
func setupRoutes(app *fiber.App) {
	// api routes
	api.SetupRoutes(app.Group("/api"))

	// frontend handler
	//
	// returns embedded files from ../frontend/dist
	// (the module has to be there, because `embed` does not allow to use
	// parent directories)
	// (the folder is created and populated by 'yarn build')
	//
	// it is a Single Page Application, routing is done by Javascript, that's
	// why it returns index.html on missing files
	//
	// for some reason route "/" does not work, have to specify greedy "*"
	app.Get("*", frontend.DistHandler)

	// catch all, just in case
	app.Use(api.MissingRouteHandler)
}
