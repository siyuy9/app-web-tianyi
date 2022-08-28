package presenter

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

type Response[Template any] struct {
	Data Template `json:"data"`
}

type ResponseError struct {
	Error string `json:"error"`
}

type SuccessModel struct {
	Status string `json:"status" example:"success"`
}

var SuccessModelDefault = &Response[SuccessModel]{
	Data: SuccessModel{Status: "success"},
}

func Success(context *fiber.Ctx, response any, code ...int) error {
	status := http.StatusOK
	if len(code) != 0 {
		status = code[0]
	}
	return context.Status(status).JSON(&Response[any]{Data: response})
}

func SuccessDefault(context *fiber.Ctx, code ...int) error {
	return Success(context, SuccessModelDefault, code...)
}

func InvalidRequestBodyFormat(err error) error {
	return pkgError.NewWithCode(
		fmt.Errorf("invalid request body format: %w", err),
		http.StatusBadRequest,
	)
}

func InvalidRequestBodyContent(err error) error {
	return pkgError.NewWithCode(
		fmt.Errorf("invalid request body content: %w", err),
		http.StatusBadRequest,
	)
}

func RouteDoesNotExist(context *fiber.Ctx) error {
	return pkgError.NewWithCode(
		fmt.Errorf(
			"route '%s %s' does not exist",
			context.Method(), context.OriginalURL(),
		),
		http.StatusNotFound,
	)
}
