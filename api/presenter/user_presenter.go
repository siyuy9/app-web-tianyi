package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
)

type UserLogin struct {
	Success
	Token string    `json:"token"`
	ID    uuid.UUID `json:"id"`
}

func NewUserLogin(context *fiber.Ctx, user *entity.User, token string) error {
	return context.Status(fiber.StatusOK).JSON(
		&UserLogin{*responseSuccess, token, user.ID},
	)
}
