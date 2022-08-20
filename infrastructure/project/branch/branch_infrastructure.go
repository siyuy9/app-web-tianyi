package infraBranch

import (
	"io"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/google/uuid"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"gorm.io/gorm"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/project/branch"
)

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) usecaseBranch.Repository {
	return &repository{database: database}
}

func (repository *repository) GetAll() ([]entity.Branch, error) {
	return repository.Find()
}

func (repository *repository) FindOne(condition *entity.Branch) (
	*entity.Branch, error,
) {
	branch := &entity.Branch{}
	return branch, repository.database.Limit(1).Find(&branch, condition).Error
}

func (repository *repository) Get(id uuid.UUID) (*entity.Branch, error) {
	branch := &entity.Branch{}
	return branch, repository.database.First(&branch, id).Error
}

func (repository *repository) Find(conditions ...interface{}) (
	[]entity.Branch, error,
) {
	branchs := make([]entity.Branch, 0)
	err := repository.database.Find(&branchs, conditions...).Error
	return branchs, err
}

func (repository *repository) Update(branch *entity.Branch) error {
	return repository.database.Updates(branch).Error
}

func (repository *repository) Save(branch *entity.Branch) error {
	return repository.database.Save(branch).Error
}

func (repository *repository) Create(branch *entity.Branch) error {
	return repository.database.Create(branch).Error
}

func (repository *repository) Delete(branch *entity.Branch) error {
	return repository.database.Delete(branch).Error
}

func (repository *repository) Migrate() error {
	return repository.database.AutoMigrate(&entity.Branch{})
}

func (repository *repository) GetRemotePipelineConfig(
	source string, branch string, filePath string,
) (config *entity.PipelineConfig, err error) {
	filesystem := memfs.New()
	_, err = git.Clone(memory.NewStorage(), filesystem, &git.CloneOptions{
		URL:           source,
		SingleBranch:  true,
		Depth:         1,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
	})
	if err != nil {
		return nil, err
	}
	fileObject, err := filesystem.Open(filePath)
	if err != nil {
		return nil, err
	}
	fileBytes, err := io.ReadAll(fileObject)
	if err != nil {
		return nil, nil
	}
	config = &entity.PipelineConfig{}
	err = hclsimple.Decode(filePath, fileBytes, nil, config)
	return config, err
}
