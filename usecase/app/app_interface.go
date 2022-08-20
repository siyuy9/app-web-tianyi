package usecaseApp

import (
	usecaseAccess "gitlab.com/kongrentian-group/tianyi/v1/usecase/access"
	usecaseJWT "gitlab.com/kongrentian-group/tianyi/v1/usecase/jwt"
	usecaseLifecycle "gitlab.com/kongrentian-group/tianyi/v1/usecase/lifecycle"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/project/branch"
	usecaseUser "gitlab.com/kongrentian-group/tianyi/v1/usecase/user"
)

type Interactor struct {
	Lifecycle usecaseLifecycle.Interactor
	User      usecaseUser.Interactor
	Access    usecaseAccess.Interactor
	JWT       usecaseJWT.Interactor
	Project   usecaseProject.Interactor
	Branch    usecaseBranch.Interactor
}
