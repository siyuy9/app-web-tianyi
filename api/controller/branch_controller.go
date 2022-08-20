package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/project/branch"
)

type BranchController interface {
	Get(context *fiber.Ctx) error
	GetProjectBranches(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
}

type branchController struct {
	interactor usecaseBranch.Interactor
}

func NewbranchController(
	branchInteractor usecaseBranch.Interactor,
) BranchController {
	return &branchController{interactor: branchInteractor}
}

// get a project branch
// @Summary get a project branch
// @Description get a project branch
// @ID get-branch
// @Tags branch
// @Security ApiKeyAuth
//
// @Param   id  path     int  true  "project id"
// @Param   name  path     int  true  "branch name"
//
// @Success 200 {object} entity.Branch
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects/{id}/branches/{name} [GET]
func (controller *branchController) Get(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	name := context.Params("name")
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
// @Param   id  path     string  true  "project id"
//
// @Success 200 {object} []entity.Branch
// @Failure 500 {object} pkg.Error
// @Router /api/v1/projects/{id}/branches [GET]
func (controller *branchController) GetProjectBranches(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("id"))
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
// @Param id   path string                 true  "project id"
// @Param Body body entity.PipelineConfig true  "pipeline config"
//
// @Success 200 {object} entity.Branch
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects/{id}/branches [POST]
func (controller *branchController) Create(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	name := context.Params("name")
	if name != "" {
		return pkg.NewErrorBadRequest("name is empty")
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
// @Param   id  path     string  true  "project id"
// @Param name path string                 true  "branch name"
//
// @Success 200 {object} entity.Branch
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects/{id}/branches/{name} [PUT]
func (controller *branchController) Update(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	branch := &entity.Branch{}
	if err = parse(context, branch); err != nil {
		return err
	}
	branch.ID = id
	if err = controller.interactor.Update(branch); err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(branch)
}
