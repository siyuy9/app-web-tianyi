package usecaseApp

import (
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/branch"
	usecaseJob "gitlab.com/kongrentian-group/tianyi/v1/usecase/job"
	usecaseJWT "gitlab.com/kongrentian-group/tianyi/v1/usecase/jwt"
	usecaseLifecycle "gitlab.com/kongrentian-group/tianyi/v1/usecase/lifecycle"
	usecasePipeline "gitlab.com/kongrentian-group/tianyi/v1/usecase/pipeline"
	usecasePool "gitlab.com/kongrentian-group/tianyi/v1/usecase/pool"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
	usecaseSession "gitlab.com/kongrentian-group/tianyi/v1/usecase/session"
	usecaseUser "gitlab.com/kongrentian-group/tianyi/v1/usecase/user"
)

type Interactor struct {
	Lifecycle usecaseLifecycle.Interactor `validate:"required"`
	User      usecaseUser.Interactor      `validate:"required"`
	// Access    usecaseAccess.Interactor    `validate:"required"`
	JWT      usecaseJWT.Interactor      `validate:"required"`
	Project  usecaseProject.Interactor  `validate:"required"`
	Branch   usecaseBranch.Interactor   `validate:"required"`
	Session  usecaseSession.Interactor  `validate:"required"`
	Pool     usecasePool.Interactor     `validate:"required"`
	Job      usecaseJob.Interactor      `validate:"required"`
	Pipeline usecasePipeline.Interactor `validate:"required"`
}
