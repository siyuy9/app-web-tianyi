package common

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	middlewareRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/storage/redis"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/db"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/store/user"

	"gitlab.com/kongrentian-groups/golang/tianyi/backend/store/gitlab_runner"

	"github.com/gofiber/fiber/v2"
)

var App *AppType = &AppType{Config: newAppConfig()}

type AppType struct {
	Config            *model.ConfigType
	Fiber             *fiber.App
	SessionStore      *session.Store
	UserStore         user.Store
	GitlabRunnerStore gitlab_runner.Store
}

func (app *AppType) InitializeServer() {
	// handle initialization errors
	// database connection failure, for example
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		log.Println("Could not initialize server:\n", err)
		app.Shutdown(1)
	}()

	database := db.Connect(app.Config.Database)

	app.UserStore = user.NewStore(database)
	app.GitlabRunnerStore = gitlab_runner.NewStore(database)

	app.Fiber = fiber.New(*app.Config.Fiber)
	app.SessionStore = session.New(session.Config{
		Storage: redis.New(*app.Config.Redis),
	})

	// use global middlewares
	// https://docs.gofiber.io/api/middleware
	app.Fiber.Use(
		cors.New(),
		compress.New(),
		etag.New(),
		csrf.New(),
		middlewareRecover.New(),
		logger.New(),
		// limiter.New(),
	)
}

func (app *AppType) Listen() {
	err := app.Fiber.Listen(":" + app.Config.Server.Port)
	if err == nil {
		return
	}
	log.Println("Runtime error:\n", err)
	app.Shutdown(1)
}

func (app *AppType) Shutdown(code int) {
	log.Println("Shutting down...")
	err := app.Fiber.Shutdown()
	if err != nil {
		log.Panic("Error during the shutdown process:\n", err)
	}
	os.Exit(code)
}

func (app *AppType) FillConfigFromEnvironment() {
	fillConfigFromEnvironment(app.Config)
}
