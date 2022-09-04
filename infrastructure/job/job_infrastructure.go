package infraJob

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecaseJob "gitlab.com/kongrentian-group/tianyi/v1/usecase/job"
)

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) usecaseJob.Repository {
	return &repository{database: database}
}

func (repository *repository) GetAll() ([]entity.Job, error) {
	return repository.Find()
}

func (repository *repository) FindOne(
	condition *entity.Job,
) (*entity.Job, error) {
	job := &entity.Job{}
	return job, repository.database.First(&job, condition).Error
}

func (repository *repository) Get(id uuid.UUID) (*entity.Job, error) {
	job := &entity.Job{}
	return job, repository.database.First(&job, id).Error
}

func (repository *repository) Find(
	conditions ...interface{},
) ([]entity.Job, error) {
	jobs := make([]entity.Job, 0)
	return jobs, repository.database.Find(&jobs, conditions...).Error
}

func (repository *repository) Save(job *entity.Job) error {
	return repository.database.Save(job).Error
}

func (repository *repository) Create(job *entity.Job) error {
	return repository.database.Create(job).Error
}

func (repository *repository) Delete(job *entity.Job) error {
	return repository.database.Delete(job).Error
}

func (repository *repository) Migrate() error {
	return repository.database.AutoMigrate(&entity.Job{})
}
