package controller

import (
	"github.com/gofiber/fiber/v2"

	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/project/branch"
)

type Branch interface {
	Get(context *fiber.Ctx) error
	GetProjectBranches(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
}

type branchController struct {
	interactor        usecaseBranch.Interactor
	projectInteractor usecaseProject.Interactor
}

func NewBranch(
	branchInteractor usecaseBranch.Interactor,
	projectInteractor usecaseProject.Interactor,
) Branch {
	return &branchController{
		interactor:        branchInteractor,
		projectInteractor: projectInteractor,
	}
}

type (
	ResponseBranch   presenter.Response[entity.Branch]
	ResponseBranches presenter.Response[[]entity.Branch]
)

// get a project branch
// @Summary get a project branch
// @Description get a project branch
// @ID get-branch
// @Tags branch
// @Security ApiKeyAuth
//
// @Param   project_id  path     string  true  "project id"
// @Param   branch_name  path     string  true  "branch name"
//
// @Success 200 {object} ResponseBranch
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/projects/{project_id}/branches/{branch_name} [GET]
func (controller *branchController) Get(context *fiber.Ctx) error {
	id, err := getProjectID(context)
	if err != nil {
		return err
	}
	name, err := getBranchName(context)
	if err != nil {
		return err
	}
	branch, err := controller.interactor.GetProjectBranch(id, name)
	if err != nil {
		return presenter.CouldNotFindProjectBranch(err)
	}
	return presenter.Success(context, branch)
}

// get project branches
// @Summary get all project branches
// @Description get all branches
// @ID get-branches
// @Tags branch
// @Security ApiKeyAuth
//
// @Param   project_id  path     string  true  "project id"
//
// @Success 200 {object} ResponseBranches
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/projects/{project_id}/branches [GET]
func (controller *branchController) GetProjectBranches(
	context *fiber.Ctx,
) error {
	id, err := getProjectID(context)
	if err != nil {
		return err
	}
	branches, err := controller.interactor.GetProjectBranches(id)
	if err != nil {
		return presenter.CouldNotFindProjectBranch(err)
	}
	return presenter.Success(context, branches)
}

type RequestBranchCreate struct {
	Name string `json:"name" form:"name"  validate:"required,min=3,max=64"`
}

// create a branch
// @Summary create a branch
// @Description create a branch
// @ID create-branch
// @Tags branch
// @Security ApiKeyAuth
//
// @Param project_id   path string                 true  "project id"
// @Param Body body RequestBranchCreate true  "creation request body"
//
// @Success 200 {object} ResponseBranch
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/projects/{project_id}/branches [POST]
func (controller *branchController) Create(context *fiber.Ctx) error {
	id, err := getProjectID(context)
	if err != nil {
		return err
	}
	request := &RequestBranchCreate{}
	if err = parse(context, request); err != nil {
		return err
	}
	branch := &entity.Branch{ProjectID: id, Name: request.Name}
	if err := controller.interactor.Create(branch); err != nil {
		return presenter.CouldNotCreateProjectBranch(err)
	}
	return presenter.Success(context, branch)
}

// update a branch
// @Summary update a branch
// @Description update a branch
// @ID update-branch
// @Tags branch
// @Security ApiKeyAuth
//
// @Param   project_id  path     string  true  "project id"
// @Param branch_name path string                 true  "branch name"
//
// @Success 200 {object} ResponseBranch
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/projects/{project_id}/branches/{branch_name} [PUT]
func (controller *branchController) Update(context *fiber.Ctx) error {
	id, err := getProjectID(context)
	if err != nil {
		return err
	}
	branchName, err := getBranchName(context)
	if err != nil {
		return err
	}
	project, err := controller.projectInteractor.Get(id)
	if err != nil {
		return presenter.CouldNotFindProject(err)
	}
	branch, err := controller.interactor.UpdateBranchFromRemote(
		project, branchName,
	)
	if err != nil {
		return presenter.CouldNotUpdateProject(err)
	}
	return presenter.Success(context, branch)
}
