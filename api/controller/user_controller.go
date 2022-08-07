package controller

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/entity"
	"gitlab.com/kongrentian-group/tianyi/pkg"
	"gitlab.com/kongrentian-group/tianyi/usecase/interactor"
)

type userController struct {
	interactor interactor.UserInteractor
	validator  *validator.Validate
}
type UserController interface {
	GetUsers(context *fiber.Ctx) error
	AddUser(context *fiber.Ctx) error
	Validate() error
}

// NewUserController: constructor, dependency injection from user service and firebase service
func NewUserController(
	userInteractor interactor.UserInteractor,
	validator *validator.Validate,
) UserController {
	return &userController{interactor: userInteractor, validator: validator}
}

func (controller *userController) GetUsers(context *fiber.Ctx) error {
	users, err := controller.interactor.FindAll()
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(err)
		return nil
	}
	return context.Status(http.StatusOK).JSON(users)
}

func (controller *userController) AddUser(context *fiber.Ctx) error {
	var user *entity.User
	if err := context.BodyParser(user); err != nil {
		return pkg.NewBadRequestError(err)
	}
	if err := controller.validator.Struct(user); err != nil {
		return err
	}
	return context.Status(fiber.StatusOK).JSON(user)
}

func (controller *userController) Validate() error {
	return errors.New("")
}
