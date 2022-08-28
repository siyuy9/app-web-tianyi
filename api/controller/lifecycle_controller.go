package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	usecaseLifecycle "gitlab.com/kongrentian-group/tianyi/v1/usecase/lifecycle"
)

type Lifecycle interface {
	Migrate(context *fiber.Ctx) error
}

type lifecycleController struct {
	interactor usecaseLifecycle.Interactor
}

func NewLifecycle(
	interactor usecaseLifecycle.Interactor,
) Lifecycle {
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
func (controller *lifecycleController) Migrate(context *fiber.Ctx) error {
	err := controller.interactor.Migrate()
	if err != nil {
		return presenter.CouldNotMigrateDatabase(context, err)
	}
	return presenter.SuccessDefault(context)
}
