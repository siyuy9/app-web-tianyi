package docs

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var (
	// returns swagger.json or swagger.yaml
	DistHandler func(context *fiber.Ctx) error
	// embeds files into the variable on initialization
	//go:embed swagger.*
	dist embed.FS
)

func init() {
	DistHandler = filesystem.New(filesystem.Config{
		Root: http.FS(dist),
	})
}
