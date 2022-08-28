package presenter

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

func CouldNotFincPileline(context *fiber.Ctx, err error) error {
	return pkgError.NewWithCode(
		fmt.Errorf("could not find project pipeline(s): %w", err),
		http.StatusNotFound, 3,
	)
}
