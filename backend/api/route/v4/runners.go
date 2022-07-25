package v4

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/common"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model/db"
)

func runnersGetHandler(context *fiber.Ctx) error {
	runners := []db.GitlabRunner{}
	common.App.DB.Find(&runners)
	return context.Status(fiber.StatusOK).JSON(runners)
}
