package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
	usecaseUser "gitlab.com/kongrentian-group/tianyi/v1/usecase/user"
)

type UserController interface {
	GetAll(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Get(context *fiber.Ctx) error
	Login(context *fiber.Ctx) error
}

type userController struct {
	interactor usecaseUser.Interactor
}

// NewUserController: constructor, dependency injection from user service and firebase service
func NewUserController(
	userInteractor usecaseUser.Interactor,
) UserController {
	return &userController{interactor: userInteractor}
}

// Get all users
// @Summary get all users
// @Description get all users
// @ID get-users
// @Tags user
// @Security ApiKeyAuth
//
// @Success 200 {object} []entity.User
// @Failure 500 {object} pkg.Error
// @Router /api/v1/users [GET]
func (controller *userController) GetAll(context *fiber.Ctx) error {
	users, err := controller.interactor.GetAll()
	if err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(users)
}

type UserRequestCreate struct {
	Username string `json:"username" form:"username"  validate:"required,min=3,max=64"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password"  validate:"required,min=8,max=64"`
}

var invalidAuthError = pkg.NewErrorForbidden("invalid login or password")

// Create user
// @Summary create a user
// @Description create a user
// @ID create-user
// @Tags user
//
// @Param Body body UserRequestCreate true "User body"
//
// @Success 200 {object} entity.User
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/users [POST]
func (controller *userController) Create(context *fiber.Ctx) error {
	request := &UserRequestCreate{}
	if err := parse(context, request); err != nil {
		return err
	}
	user := &entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	if err := controller.interactor.Create(user); err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(user)
}

// get a user
// @Summary get a user
// @Description get a user
// @ID get-user
// @Tags user
// @Security ApiKeyAuth
//
// @Param   user_id  path     string  true  "account id"
//
// @Success 200 {object} entity.User
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/users/user/{user_id} [GET]
func (controller *userController) Get(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("user_id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	user, err := controller.interactor.Get(id)
	if err != nil {
		return err
	}
	return context.Status(fiber.StatusOK).JSON(user)
}

type UserRequestLogin struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// login
// @Summary login
// @Description login
// @ID login-user
// @Tags user
//
// @Success 200 {object} presenter.UserLogin
// @Failure 500 {object} pkg.Error
// @Failure 403 {object} pkg.Error
// @Router /api/v1/users/login [GET]
func (controller *userController) Login(context *fiber.Ctx) error {
	request := &UserRequestLogin{}
	if err := parse(context, request); err != nil {
		return err
	}
	user, err := controller.interactor.FindByUsername(request.Username)
	if err != nil {
		return invalidAuthError
	}
	matches, err := controller.interactor.PasswordHashCheck(
		user,
		request.Password,
	)
	if !matches {
		return invalidAuthError
	}
	if err != nil {
		return err
	}
	token, err := controller.interactor.JWT().New(user)
	if err != nil {
		return err
	}
	return presenter.NewUserLogin(context, user, token)
}
