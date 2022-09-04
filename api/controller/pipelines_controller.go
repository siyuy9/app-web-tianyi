package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/branch"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
)

type Pipeline interface {
	Create(context *fiber.Ctx) error
}

type pipelineController struct {
	branchInteractor  usecaseBranch.Interactor
	projectInteractor usecaseProject.Interactor
}

func NewPipeline(
	branchInteractor usecaseBranch.Interactor,
	projectInteractor usecaseProject.Interactor,
) Pipeline {
	return &pipelineController{
		branchInteractor:  branchInteractor,
		projectInteractor: projectInteractor,
	}
}

type (
	ResponsePipeline = presenter.Response[entity.PipelineConfig]
)

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
// @Success 200 {object} map[string]any
// @Failure 500 {object} presenter.ResponseError
// @Router /api/v1/projects/{project_id}/branches/{branch_name}/pipelines/{pipeline_name} [POST]
func (controller *pipelineController) Create(context *fiber.Ctx) error {
	id, err := getProjectID(context)
	if err != nil {
		return err
	}
	name, err := getBranchName(context)
	if err != nil {
		return err
	}
	branch, err := controller.branchInteractor.GetProjectBranch(id, name)
	if err != nil {
		return presenter.CouldNotFindProjectBranch(err)
	}
	for _, pipeline := range branch.Config.Pipelines {
		if pipeline.Name != context.Params("pipeline_name") {
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
				return errors.New(
					"missing declaration for job " + pipelineJob.Name,
				)
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
			defer response.Body.Close()
			var responseMap map[string]interface{}
			err = json.NewDecoder(response.Body).Decode(&responseMap)
			if err != nil {
				return err
			}
			return context.Status(response.StatusCode).JSON(responseMap)
		}
	}
	return errors.New("no jobs to run")
}
