package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecaseJWT "gitlab.com/kongrentian-group/tianyi/v1/usecase/jwt"
	usecaseSession "gitlab.com/kongrentian-group/tianyi/v1/usecase/session"
	usecaseUser "gitlab.com/kongrentian-group/tianyi/v1/usecase/user"
)

type User interface {
	GetAll(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Get(context *fiber.Ctx) error
	Login(context *fiber.Ctx) error
}

type userController struct {
	interactor        usecaseUser.Interactor
	jwtInteractor     usecaseJWT.Interactor
	sessionInteractor usecaseSession.Interactor
}

// NewUser: constructor, dependency injection from user service and firebase service
func NewUser(
	userInteractor usecaseUser.Interactor, jwtInteractor usecaseJWT.Interactor,
	sessionInteractor usecaseSession.Interactor,
) User {
	return &userController{
		interactor: userInteractor, jwtInteractor: jwtInteractor,
		sessionInteractor: sessionInteractor,
	}
}

type (
	ResponseUser  = presenter.Response[entity.User]
	ResponseUsers = presenter.Response[entity.User]
)

// Get all users
// @Summary get all users
// @Description get all users
// @ID get-users
// @Tags user
// @Security ApiKeyAuth
//
// @Success 200 {object} ResponseUsers
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/users [GET]
func (controller *userController) GetAll(context *fiber.Ctx) error {
	users, err := controller.interactor.GetAll()
	if err != nil {
		return presenter.CouldNotFindUser(err)
	}
	return presenter.Success(context, users)
}

type UserRequestCreate struct {
	Username string `json:"username" form:"username"  validate:"required,min=3,max=64"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password"  validate:"required,min=8,max=64"`
}

// Create user
// @Summary create a user
// @Description create a user
// @ID create-user
// @Tags user
//
// @Param Body body UserRequestCreate true "User body"
//
// @Success 200 {object} ResponseUser
// @Failure 500 {object} presenter.ResponseError
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
		return presenter.CouldNotCreateUser(err)
	}
	return presenter.Success(context, user)
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
// @Success 200 {object} ResponseUser
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/users/user/{user_id} [GET]
func (controller *userController) Get(context *fiber.Ctx) error {
	id, err := getUserID(context)
	if err != nil {
		return err
	}
	user, err := controller.interactor.Get(id)
	if err != nil {
		return presenter.CouldNotFindUser(err)
	}
	return presenter.Success(context, user)
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
// @Success 200 {object} ResponseUser
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/users/login [GET]
func (controller *userController) Login(context *fiber.Ctx) error {
	request := &UserRequestLogin{}
	if err := parse(context, request); err != nil {
		return err
	}
	user, err := controller.interactor.FindByUsername(request.Username)
	if err != nil {
		return presenter.InvalidLoginOrPassword()
	}
	matches, err := controller.interactor.PasswordHashCheck(
		user,
		request.Password,
	)
	if err != nil {
		return err
	}
	if !matches {
		return presenter.InvalidLoginOrPassword()
	}
	session, err := controller.sessionInteractor.Get(context)
	if err != nil {
		return presenter.CouldNotGetSession(err)
	}
	context.Cookie(&fiber.Cookie{
		Name: usecaseSession.SessionCookie, Value: session.ID(),
	})
	session.Set(usecaseSession.UserID, user.ID.String())
	if err := session.Save(); err != nil {
		return presenter.CouldNotSaveSession(err)
	}
	return presenter.Success(context, user)
}
