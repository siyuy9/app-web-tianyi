package controller

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/project/branch"
)

type PipelineController interface {
	Create(context *fiber.Ctx) error
}

type pipelineController struct {
	branchInteractor  usecaseBranch.Interactor
	projectInteractor usecaseProject.Interactor
}

func NewPipelineController(
	branchInteractor usecaseBranch.Interactor,
	projectInteractor usecaseProject.Interactor,
) PipelineController {
	return &pipelineController{
		branchInteractor:  branchInteractor,
		projectInteractor: projectInteractor,
	}
}

// create a pipeline
// @Summary create a pipeline
// @Description create a pipeline
// @ID create-pipeline
// @Tags pipeline
// @Security ApiKeyAuth
//
// @Param   project_id  path     string  true  "project id"
// @Param   branch_name  path     string  true  "branch name"
// @Param   pipeline_name  path     string  true  "pipeline name"
//
// @Success 200
// @Failure 500 {object} pkg.Error
// @Failure 400 {object} pkg.Error
// @Router /api/v1/projects/{project_id}/branches/{branch_name}/pipelines/{pipeline_name} [POST]
func (controller *pipelineController) Create(context *fiber.Ctx) error {
	id, err := uuid.Parse(context.Params("project_id"))
	if err != nil {
		return pkg.NewErrorBadRequest(err)
	}
	name := context.Params("branch_name")
	if name == "" {
		return pkg.NewErrorBadRequest("name is empty")
	}
	branch, err := controller.branchInteractor.GetProjectBranch(id, name)
	if err != nil {
		return err
	}
	for _, pipeline := range branch.Config.Pipelines {
		if pipeline.Name != context.Params("branch_name") {
			continue
		}
		for _, pipelineJob := range pipeline.Jobs {
			var jobConfig *entity.PipelineConfigJob
			for _, jobConfigTemp := range branch.Config.Jobs {
				if jobConfigTemp.Name == pipelineJob.Name {
					jobConfig = &jobConfigTemp
					break
				}
			}
			if jobConfig == nil {
				break
			}
			request, err := http.NewRequest(
				jobConfig.RequestType, jobConfig.URL, nil,
			)
			if err != nil {
				return err
			}
			query := request.URL.Query()
			for key, value := range jobConfig.Query {
				query.Add(key, value)
			}
			request.URL.RawQuery = query.Encode()
			response, err := http.DefaultClient.Do(request)
			if err != nil {
				return err
			}
			context.Status(http.StatusOK).JSON(response)
		}
	}
	return errors.New("no jobs to run")
}
