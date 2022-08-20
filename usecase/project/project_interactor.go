package usecaseProject

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/project/branch"
)

type interactor struct {
	repository       Repository
	branchInteractor usecaseBranch.Interactor
}

func New(
	repository Repository,
	branchInteractor usecaseBranch.Interactor,
) Interactor {
	return &interactor{
		repository:       repository,
		branchInteractor: branchInteractor,
	}
}

func (interactor *interactor) Validate(project *entity.Project) error {
	return pkg.ValidateStruct(project)
}

func (interactor *interactor) Create(project *entity.Project) error {
	if err := pkg.ValidateStruct(project); err != nil {
		return err
	}
	if project.NamespaceID == uuid.Nil {
		project.Path = strings.ToLower(project.Name)
	} else {
		panic("namespaces are not implemented yet")
		// project.Path = project.Namespace.Path + "/" + project.Name
	}
	config, err := interactor.branchInteractor.GetRemotePipelineConfig(
		project.Source, project.DefaultBranch, ".tianyi/config.hcl",
	)
	if err != nil {
		return err
	}
	project.Branches = []entity.Branch{{Config: config}}
	return interactor.repository.Create(project)
}

func (interactor *interactor) Update(project *entity.Project) error {
	if err := pkg.ValidateStruct(project); err != nil {
		return err
	}
	return interactor.repository.Update(project)
}

func (interactor *interactor) GetAll() ([]entity.Project, error) {
	return interactor.repository.GetAll()
}

func (interactor *interactor) Get(id uuid.UUID) (*entity.Project, error) {
	return interactor.repository.Get(id)
}

func (interactor *interactor) GetByPath(path string) (*entity.Project, error) {
	return interactor.repository.FindOne(
		&entity.Project{Path: path},
	)
}

func (interactor *interactor) GetByName(name string) (
	*entity.Project, error,
) {
	return interactor.repository.FindOne(
		&entity.Project{Name: name},
	)
}
