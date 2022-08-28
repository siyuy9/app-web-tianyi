package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
)

const (
	PathUserID       = "user_id"
	PathProjectID    = "project_id"
	PathBranchName   = "branch_name"
	PathPipelineName = "pipeline_name"
)

// parse request body into target, then validate it
// returns pkgError.Error
func parse(context *fiber.Ctx, target interface{}) error {
	if err := context.BodyParser(target); err != nil {
		return presenter.InvalidRequestBodyFormat(err)
	}
	if err := pkg.ValidateStruct(target); err != nil {
		return presenter.InvalidRequestBodyContent(err)
	}
	return nil
}

// get project id from path parameters
// returns pkgError.Error
func getProjectID(context *fiber.Ctx) (uuid.UUID, error) {
	id, err := uuid.Parse(context.Params(PathProjectID))
	if err != nil {
		return uuid.Nil, presenter.InvalidProjectID(err)
	}
	return id, nil
}

// get branch name from path parameters
// returns pkgError.Error
func getBranchName(context *fiber.Ctx) (string, error) {
	name := context.Params(PathBranchName)
	if name == "" {
		return name, presenter.InvalidBranchName(errors.New("empty"))
	}
	return name, nil
}

// get user id from path parameters
// returns pkgError.Error
func getUserID(context *fiber.Ctx) (uuid.UUID, error) {
	id, err := uuid.Parse(context.Params(PathUserID))
	if err != nil {
		return uuid.Nil, presenter.InvalidUserID(err)
	}
	return id, nil
}
