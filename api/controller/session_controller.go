package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	usecaseJWT "gitlab.com/kongrentian-group/tianyi/v1/usecase/jwt"
	usecaseSession "gitlab.com/kongrentian-group/tianyi/v1/usecase/session"
)

var UnauthorizedRoutes = map[string]bool{
	"/api/v1/users":       true,
	"/api/v1/users/login": true,
}

type Session interface {
	CheckSession(context *fiber.Ctx) error
}

type sessionController struct {
	interactor    usecaseSession.Interactor
	jwtController JWT
	jwtInteractor usecaseJWT.Interactor
}

func NewSession(
	interactor usecaseSession.Interactor, jwtController JWT,
	jwtInteractor usecaseJWT.Interactor,
) Session {
	return &sessionController{
		interactor:    interactor,
		jwtController: jwtController,
		jwtInteractor: jwtInteractor,
	}
}

func (controller *sessionController) CheckSession(context *fiber.Ctx) error {
	// allow registration and login without a token
	if _, defined := UnauthorizedRoutes[context.Path()]; defined {
		return nil
	}
	session, err := controller.interactor.Get(context)
	if err != nil {
		return err
	}
	log.Println(session.ID(), session.Keys())
	// if current session contains user_id, the user is logged in
	if session.Get(usecaseSession.UserID) != nil {
		return nil
	}
	// check if there is a JWT in the header
	if err := controller.jwtController.CheckJWT(context); err != nil {
		return err
	}
	claims, err := controller.jwtInteractor.GetClaims(context.Locals("user"))
	if err != nil {
		return err
	}
	session.Set(usecaseSession.UserID, claims.ID)
	return nil
}
