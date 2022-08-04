package frontend

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var (
	DistHandler func(context *fiber.Ctx) error
	// embeds directory into the variable on initialization
	//go:embed dist
	dist embed.FS
)

func init() {
	DistHandler = filesystem.New(filesystem.Config{
		Root:         http.FS(dist),
		PathPrefix:   "dist",
		NotFoundFile: "dist/index.html",
	})
}
