package common

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/db"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

var App *AppType = newApp()

type AppType struct {
	// app config
	Config *model.ConfigType
	// database
	DB *gorm.DB
	// fiber instance
	Fiber *fiber.App
	// session store
	Sessions *session.Store
}

func newApp() *AppType {
	config := model.NewConfigType()
	return &AppType{
		Config: config,
		DB:     db.ConnectDatabase(config),
		Fiber: fiber.New(fiber.Config{
			ServerHeader: "tianyi",
			AppName:      "tianyi",
		}),
		Sessions: session.New(session.Config{
			Storage: memory.New(memory.Config{
				GCInterval: 24 * time.Hour,
			}),
		}),
	}
}
