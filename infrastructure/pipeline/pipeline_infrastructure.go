package infraPipeline

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecasePipeline "gitlab.com/kongrentian-group/tianyi/v1/usecase/pipeline"
)

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) usecasePipeline.Repository {
	return &repository{database: database}
}

func (repository *repository) GetAll() ([]entity.Pipeline, error) {
	return repository.Find()
}

func (repository *repository) FindOne(
	condition *entity.Pipeline,
) (*entity.Pipeline, error) {
	Pipeline := &entity.Pipeline{}
	return Pipeline, repository.database.First(&Pipeline, condition).Error
}

func (repository *repository) Get(id uuid.UUID) (*entity.Pipeline, error) {
	Pipeline := &entity.Pipeline{}
	return Pipeline, repository.database.First(&Pipeline, id).Error
}

func (repository *repository) Find(
	conditions ...interface{},
) ([]entity.Pipeline, error) {
	Pipelines := make([]entity.Pipeline, 0)
	err := repository.database.Find(&Pipelines, conditions...).Error
	return Pipelines, err
}

func (repository *repository) Save(Pipeline *entity.Pipeline) error {
	return repository.database.Save(Pipeline).Error
}

func (repository *repository) Create(Pipeline *entity.Pipeline) error {
	return repository.database.Create(Pipeline).Error
}

func (repository *repository) Delete(Pipeline *entity.Pipeline) error {
	return repository.database.Delete(Pipeline).Error
}

func (repository *repository) Migrate() error {
	return repository.database.AutoMigrate(&entity.Pipeline{})
}
