package useJob

import "gitlab.com/kongrentian-group/tianyi/v1/entity"

type interactor struct {
	repository Repository
}

func New(repository Repository) Interactor {
	return &interactor{repository}
}

func (interactor *interactor) Repository() Repository {
	return interactor.repository
}

func (interactor *interactor) Create(
	pipeline *entity.PipelineConfigPipeline,
) *entity.Pipeline {
	return nil
}
