package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// serve embedded frontend files
type FrontendController interface {
	Serve(context *fiber.Ctx) error
}

type frontendController struct {
	filesystemController func(context *fiber.Ctx) error
}

func NewFrontendController(
	frontendFilesystem http.FileSystem,
) FrontendController {
	return &frontendController{
		filesystemController: filesystem.New(filesystem.Config{
			Root:         frontendFilesystem,
			PathPrefix:   "dist",
			NotFoundFile: "dist/index.html",
		}),
	}
}

func (controller *frontendController) Serve(context *fiber.Ctx) error {
	return controller.filesystemController(context)
}
