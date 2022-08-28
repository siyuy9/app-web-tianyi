package presenter

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

func CouldNotMigrateDatabase(context *fiber.Ctx, err error) error {
	return pkgError.NewWithCode(
		fmt.Errorf("could not migrate database: %w", err),
		http.StatusInternalServerError,
	)
}
