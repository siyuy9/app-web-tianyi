package presenter

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
)

type Success struct {
	Status string `json:"status" example:"success"`
}

var responseSuccess = &Success{Status: "success"}

func NewSuccess(context *fiber.Ctx, code ...int) error {
	status := http.StatusOK
	if len(code) != 0 {
		status = code[0]
	}
	return context.Status(status).JSON(responseSuccess)
}

func RouteDoesNotExist(context *fiber.Ctx) error {
	return pkg.NewError(
		fmt.Sprintf(
			"Route '%s %s' does not exist",
			context.Method(),
			context.OriginalURL(),
		),
		http.StatusNotFound,
	)
}
