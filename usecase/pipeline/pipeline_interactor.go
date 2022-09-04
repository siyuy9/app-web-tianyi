package usecasePipeline

import (
	"errors"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
)

func GetFromMap(input map[string]interface{}) (Interactor, error) {
	interactor, ok := input[InteractorKey]
	if !ok {
		return nil, errors.New("missing pipeline interactor")
	}
	interactorAsserted, ok := interactor.(Interactor)
	if !ok {
		return nil, errors.New("invalid pipeline interactor format")
	}
	return interactorAsserted, nil
}

type interactor struct{}

func New() Interactor {
	return &interactor{}
}

func (interactor *interactor) Create(
	pipeline *entity.PipelineConfigPipeline,
) *entity.Pipeline {
	return nil
}
