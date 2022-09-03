package usecasePipeline

import "gitlab.com/kongrentian-group/tianyi/v1/entity"

type interactor struct{}

func New() Interactor {
	return &interactor{}
}

func (interactor *interactor) Create(
	pipeline *entity.PipelineConfigPipeline,
) *entity.Pipeline {
	return nil
}
