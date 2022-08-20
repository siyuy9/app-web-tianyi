package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	usecaseLifecycle "gitlab.com/kongrentian-group/tianyi/v1/usecase/lifecycle"
)

type LifecycleController interface {
	Migrate(context *fiber.Ctx) error
}

type lifecycleController struct {
	interactor usecaseLifecycle.Interactor
}

func NewLifecycleController(
	interactor usecaseLifecycle.Interactor,
) LifecycleController {
	return &lifecycleController{interactor: interactor}
}

// migrate databases
// @Summary migrate databases
// @Description migrate databases
// @ID database-migrate
// @Tags database
// @Security ApiKeyAuth
//
// @Success 200 {object} presenter.Success
// @Failure 500 {object} pkg.Error
// @Router /api/v1/database/migrate [POST]
func (controller *lifecycleController) Migrate(context *fiber.Ctx) error {
	err := controller.interactor.Migrate()
	if err != nil {
		return err
	}
	return presenter.NewSuccess(context)
}
