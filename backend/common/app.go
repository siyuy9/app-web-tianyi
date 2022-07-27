package common

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/db"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model"

	"github.com/gofiber/fiber/v2"
)

var App *AppType = newApp()

type AppType struct {
	// app config
	Config *model.ConfigType
	// database, useless until Database.Connect() is called
	Database *db.DatabaseService
	// fiber instance, nil until App.InitializeServer is called
	Fiber *fiber.App
	// session store, nil until App.InitializeServer is called
	Sessions *session.Store
}

func newApp() *AppType {
	config := model.NewConfigType()
	return &AppType{
		Config:   config,
		Database: db.NewDatabaseService(config.Database),
		Fiber:    nil,
		Sessions: nil,
	}
}

func (app *AppType) InitializeServer() {
	app.Database.Connect()
	app.Fiber = fiber.New(*app.Config.Server.Fiber)
	app.Sessions = session.New(session.Config{
		Storage: memory.New(memory.Config{
			GCInterval: 24 * time.Hour,
		}),
	})
}
