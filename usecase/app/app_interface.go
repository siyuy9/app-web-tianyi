package usecaseApp

import (
	usecaseJWT "gitlab.com/kongrentian-group/tianyi/v1/usecase/jwt"
	usecaseLifecycle "gitlab.com/kongrentian-group/tianyi/v1/usecase/lifecycle"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/project/branch"
	usecaseSession "gitlab.com/kongrentian-group/tianyi/v1/usecase/session"
	usecaseUser "gitlab.com/kongrentian-group/tianyi/v1/usecase/user"
)

type Interactor struct {
	Lifecycle usecaseLifecycle.Interactor `validate:"required"`
	User      usecaseUser.Interactor      `validate:"required"`
	// Access    usecaseAccess.Interactor    `validate:"required"`
	JWT     usecaseJWT.Interactor     `validate:"required"`
	Project usecaseProject.Interactor `validate:"required"`
	Branch  usecaseBranch.Interactor  `validate:"required"`
	Session usecaseSession.Interactor `validate:"required"`
}
