package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// serve embedded frontend files
type Frontend interface {
	Serve(context *fiber.Ctx) error
	ServeSwagger(context *fiber.Ctx) error
}

type frontendController struct {
	filesystemController func(context *fiber.Ctx) error
	swaggerController    func(context *fiber.Ctx) error
}

func NewFrontend(
	frontendFilesystem http.FileSystem,
	swaggerFilesystem http.FileSystem,
) Frontend {
	return &frontendController{
		filesystemController: filesystem.New(filesystem.Config{
			Root:         frontendFilesystem,
			PathPrefix:   "dist",
			NotFoundFile: "dist/index.html",
		}),
		swaggerController: filesystem.New(filesystem.Config{
			Root: swaggerFilesystem,
		}),
	}
}

// get the current OpenAPI schema
// @Summary get the current OpenAPI schema
// @Description get the current OpenAPI schema
// @ID get-openapi
// @Tags openapi
// @Security ApiKeyAuth
//
// @Success 200 {object} map[string]any
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/swagger/swagger.json [PUT]
func (controller *frontendController) ServeSwagger(context *fiber.Ctx) error {
	return controller.swaggerController(context)
}

func (controller *frontendController) Serve(context *fiber.Ctx) error {
	return controller.filesystemController(context)
}
