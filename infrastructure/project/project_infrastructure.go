package infraProject

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
)

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) usecaseProject.Repository {
	return &repository{database: database}
}

func (repository *repository) GetAll() ([]entity.Project, error) {
	return repository.Find()
}

func (repository *repository) FindOne(condition *entity.Project) (
	*entity.Project, error,
) {
	project := &entity.Project{}
	return project, repository.database.First(&project, condition).Error
}

func (repository *repository) Get(id uuid.UUID) (*entity.Project, error) {
	project := &entity.Project{}
	return project, repository.database.First(&project, id).Error
}

func (repository *repository) Find(conditions ...interface{}) (
	[]entity.Project, error,
) {
	projects := make([]entity.Project, 0)
	err := repository.database.Find(&projects, conditions...).Error
	return projects, err
}

func (repository *repository) Update(project *entity.Project) error {
	return repository.database.Updates(project).Error
}

func (repository *repository) Create(project *entity.Project) error {
	return repository.database.Create(project).Error
}

func (repository *repository) Save(project *entity.Project) error {
	return repository.database.Save(project).Error
}

func (repository *repository) Delete(project *entity.Project) error {
	return repository.database.Delete(project).Error
}

func (repository *repository) Migrate() error {
	return repository.database.AutoMigrate(
		&entity.Namespace{},
		&entity.Project{},
		&entity.Branch{},
	)
}
