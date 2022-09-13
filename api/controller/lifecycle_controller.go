package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	useLifecycle "gitlab.com/kongrentian-group/tianyi/v1/usecase/lifecycle"
)

type Lifecycle interface {
	Migrate(ctx *fiber.Ctx) error
}

type lifecycleController struct {
	interactor useLifecycle.Interactor
}

func NewLifecycle(interactor useLifecycle.Interactor) Lifecycle {
	return &lifecycleController{interactor: interactor}
}

// migrate databases
// @Summary migrate databases
// @Description migrate databases
// @ID database-migrate
// @Tags database
// @Security ApiKeyAuth
//
// @Success 200 {object} presenter.SuccessModel
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/database/migrate [POST]
func (c *lifecycleController) Migrate(ctx *fiber.Ctx) error {
	if err := c.interactor.Migrate(); err != nil {
		return presenter.CouldNotMigrateDatabase(err)
	}
	return presenter.SuccessDefault(ctx)
}
