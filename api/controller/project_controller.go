package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
)

type ProjectController interface {
	GetByID(context *fiber.Ctx) error
	Get(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
}

type projectController struct {
	interactor usecaseProject.Interactor
}

func NewProjectController(
	projectInteractor usecaseProject.Interactor,
) ProjectController {
	return &projectController{interactor: projectInteractor}
}

// get a project
// @Summary get a project
// @Description get a project
// @ID get-project
// @Tags project
// @Security ApiKeyAuth
//
// @Param   id  path     int  true  "project id"
//
// @Success 200 {object} entity.Project
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects/{id} [GET]
func (controller *projectController) GetByID(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	project, err := controller.interactor.Get(id)
	if err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(project)
}

// get projects
// @Summary get projects
// @Description get projects
// @ID get-projects
// @Tags project
// @Security ApiKeyAuth
//
// @Param        path    query     string  false  "search by path"
//
// @Success 200 {object} []entity.Project
// @Failure 500 {object} pkg.Error
// @Router /api/v1/projects [GET]
func (controller *projectController) Get(context *fiber.Ctx) error {
	path := context.Query("path")
	if path == "" {
		projects, err := controller.interactor.GetAll()
		if err != nil {
			return err
		}
		return context.Status(http.StatusOK).JSON(projects)
	}
	project, err := controller.interactor.GetByPath(path)
	if err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON([]entity.Project{*project})
}

type ProjectRequestCreate struct {
	Name          string `json:"name" validate:"required"`
	Source        string `json:"source" validate:"required"`
	DefaultBranch string `json:"default_branch"`
}

// create a project
// @Summary create a project
// @Description create a project
// @ID create-project
// @Tags project
// @Security ApiKeyAuth
//
// @Param Body body ProjectRequestCreate true "Project body"
//
// @Success 200 {object} entity.Project
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects [POST]
func (controller *projectController) Create(context *fiber.Ctx) error {
	request := &ProjectRequestCreate{}
	if err := parse(context, request); err != nil {
		return err
	}
	project := &entity.Project{
		Name:          request.Name,
		Source:        request.Source,
		DefaultBranch: request.DefaultBranch,
	}
	if err := controller.interactor.Create(project); err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(project)
}

// update a project
// @Summary update a project
// @Description update a project
// @ID update-project
// @Tags project
// @Security ApiKeyAuth
//
// @Param   id  path     int  true  "project id"
// @Param Body body entity.Project true "Project body"
//
// @Success 200 {object} entity.Project
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects/{id} [PUT]
func (controller *projectController) Update(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	project := &entity.Project{}
	if err = parse(context, project); err != nil {
		return err
	}
	project.ID = id
	if err = controller.interactor.Update(project); err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(project)
}
