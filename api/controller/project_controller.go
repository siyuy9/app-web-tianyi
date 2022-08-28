package controller

import (
	"github.com/gofiber/fiber/v2"

	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
)

type Project interface {
	GetByID(context *fiber.Ctx) error
	Get(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
}

type (
	ResponseProject  = presenter.Response[entity.Project]
	ResponseProjects = presenter.Response[[]entity.Project]
)

type projectController struct {
	interactor usecaseProject.Interactor
}

func NewProject(
	projectInteractor usecaseProject.Interactor,
) Project {
	return &projectController{interactor: projectInteractor}
}

// get a project
// @Summary get a project
// @Description get a project
// @ID get-project
// @Tags project
// @Security ApiKeyAuth
//
// @Param   project_id  path     int  true  "project id"
//
// @Success 200 {object} ResponseProject
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/projects/{project_id} [GET]
func (controller *projectController) GetByID(context *fiber.Ctx) error {
	id, err := getProjectID(context)
	if err != nil {
		return err
	}
	project, err := controller.interactor.Get(id)
	if err != nil {
		return presenter.CouldNotFindProject(err)
	}
	return presenter.Success(context, project)
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
// @Success 200 {object} ResponseProjects
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/projects [GET]
func (controller *projectController) Get(context *fiber.Ctx) error {
	path := context.Query("path")
	if path == "" {
		projects, err := controller.interactor.GetAll()
		if err != nil {
			return presenter.CouldNotFindProject(err)
		}
		return presenter.Success(context, projects)
	}
	project, err := controller.interactor.GetByPath(path)
	if err != nil {
		return presenter.CouldNotFindProject(err)
	}
	return presenter.Success(context, []entity.Project{*project})
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
// @Success 200 {object} ResponseProject
// @Failure 500 {object} presenter.ResponseError
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
		return presenter.CouldNotCreateProject(err)
	}
	return presenter.Success(context, project)
}

// update a project
// @Summary update a project
// @Description update a project
// @ID update-project
// @Tags project
// @Security ApiKeyAuth
//
// @Param   project_id  path     string  true  "project id"
// @Param Body body entity.Project true "Project body"
//
// @Success 200 {object} ResponseProject
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/projects/{project_id} [PUT]
func (controller *projectController) Update(context *fiber.Ctx) error {
	id, err := getProjectID(context)
	if err != nil {
		return err
	}
	project := &entity.Project{}
	if err = parse(context, project); err != nil {
		return err
	}
	project.ID = id
	if err = controller.interactor.Update(project); err != nil {
		return presenter.CouldNotUpdateProject(err)
	}
	return presenter.Success(context, project)
}
