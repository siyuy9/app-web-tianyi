package usecasePipeline

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecaseJob "gitlab.com/kongrentian-group/tianyi/v1/usecase/job"
)

type interactor struct {
	repository    Repository
	jobInteractor usecaseJob.Interactor
}

func New(
	repository Repository, jobInteractor usecaseJob.Interactor,
) Interactor {
	return &interactor{repository, jobInteractor}
}

func (interactor *interactor) Create(
	pipeline *entity.PipelineConfigPipeline,
) *entity.Pipeline {
	return nil
}

func (interactor *interactor) RunJob(job *entity.Job) error {
	request, err := http.NewRequest(
		job.Config.RequestType, job.Config.URL, nil,
	)
	if err != nil {
		return err
	}
	query := request.URL.Query()
	for key, value := range job.Config.Query {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	var responseMap map[string]interface{}
	if json.NewDecoder(response.Body).Decode(&responseMap); err != nil {
		return err
	}
	return nil
}

func (interactor *interactor) JobErrorHandler(
	job *entity.Job, err error,
) error {
	if err == nil {
		return nil
	}
	job.Log += fmt.Sprintf("an error has occured:\n%+v", err)
	job.Result = false
	job.Status = entity.JobError
	return interactor.jobInteractor.Repository().Save(job)
}

func (interactor *interactor) Repository() Repository {
	return interactor.repository
}
