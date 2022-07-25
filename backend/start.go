package backend

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/api"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/common"
)

func Start() {
	// use global middlewares
	// https://docs.gofiber.io/api/middleware
	common.App.Fiber.Use(
		cors.New(),
		compress.New(),
		etag.New(),
		csrf.New(),
		recover.New(),
		logger.New(),
	)
	// setup routes
	api.SetRoutes(common.App.Fiber)

	// listen
	err := common.App.Fiber.Listen(
		common.App.Config.Server.Host +
			":" +
			common.App.Config.Server.Port)
	if err != nil {
		log.Panic("Runtime error: ", err)
	}
}
