package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/project/branch"
)

type BranchController interface {
	Get(context *fiber.Ctx) error
	GetProjectBranches(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
}

type branchController struct {
	interactor        usecaseBranch.Interactor
	projectInteractor usecaseProject.Interactor
}

func NewbranchController(
	branchInteractor usecaseBranch.Interactor,
	projectInteractor usecaseProject.Interactor,
) BranchController {
	return &branchController{
		interactor:        branchInteractor,
		projectInteractor: projectInteractor,
	}
}

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
// @Success 200 {object} entity.Branch
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects/{project_id}/branches/{branch_name} [GET]
func (controller *branchController) Get(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("project_id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	name := context.Params("branch_name")
	if name == "" {
		return pkg.NewErrorBadRequest("name is empty")
	}
	branch, err := controller.interactor.GetProjectBranch(id, name)
	if err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(branch)
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
// @Success 200 {object} []entity.Branch
// @Failure 500 {object} pkg.Error
// @Router /api/v1/projects/{project_id}/branches [GET]
func (controller *branchController) GetProjectBranches(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("project_id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	branches, err := controller.interactor.GetProjectBranches(id)
	if err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(branches)
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
// @Success 200 {object} entity.Branch
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects/{project_id}/branches [POST]
func (controller *branchController) Create(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("project_id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	request := &RequestBranchCreate{}
	if err = parse(context, request); err != nil {
		return err
	}
	branch := &entity.Branch{ProjectID: id, Name: request.Name}
	if err := controller.interactor.Create(branch); err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(branch)
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
// @Success 200 {object} entity.Branch
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects/{project_id}/branches/{branch_name} [PUT]
func (controller *branchController) Update(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("project_id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	branchName := context.Params("branch_name")
	if branchName == "" {
		return pkg.NewErrorBadRequest(err)
	}
	project, err := controller.projectInteractor.Get(id)
	if err != nil {
		return err
	}
	branch, err := controller.interactor.UpdateBranchFromRemote(project, branchName)
	if err != nil {
		return nil
	}
	return context.Status(http.StatusOK).JSON(branch)
}
