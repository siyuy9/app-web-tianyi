package usecaseBranch

import (
	"github.com/google/uuid"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
)

type interactor struct {
	repository Repository
}

func New(repository Repository) Interactor {
	return &interactor{
		repository: repository,
	}
}

func (interactor *interactor) Validate(branch *entity.Branch) error {
	return pkg.ValidateStruct(branch)
}

func (interactor *interactor) Create(branch *entity.Branch) error {
	if err := pkg.ValidateStruct(branch); err != nil {
		return err
	}
	return interactor.repository.Save(branch)
}

func (interactor *interactor) Update(branch *entity.Branch) error {
	if err := pkg.ValidateStruct(branch); err != nil {
		return err
	}
	return interactor.repository.Update(branch)
}

func (interactor *interactor) GetAll() ([]entity.Branch, error) {
	return interactor.repository.GetAll()
}

func (interactor *interactor) Get(id uuid.UUID) (*entity.Branch, error) {
	return interactor.repository.Get(id)
}

func (interactor *interactor) GetProjectBranch(projectID uuid.UUID, name string) (
	*entity.Branch, error,
) {
	return interactor.repository.FindOne(
		&entity.Branch{Name: name, ProjectID: projectID},
	)
}

func (interactor *interactor) GetProjectBranches(projectID uuid.UUID) (
	[]entity.Branch, error,
) {
	return interactor.repository.Find(&entity.Branch{ProjectID: projectID})
}

func (interactor *interactor) GetRemotePipelineConfig(
	source string, branch string, filePath string,
) (config *entity.PipelineConfig, err error) {
	return interactor.repository.GetRemotePipelineConfig(
		source, branch, filePath,
	)
}
