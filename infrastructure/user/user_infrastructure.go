package infraUser

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	usecaseUser "gitlab.com/kongrentian-group/tianyi/v1/usecase/user"
)

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) usecaseUser.Repository {
	return &repository{database: database}
}

func (repository *repository) GetAll() ([]entity.User, error) {
	return repository.Find()
}

func (repository *repository) FindOne(
	condition *entity.User,
) (*entity.User, error) {
	user := &entity.User{}
	return user, repository.database.First(&user, condition).Error
}

func (repository *repository) Get(id uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	return user, repository.database.First(&user, id).Error
}

func (repository *repository) Find(
	conditions ...interface{},
) ([]entity.User, error) {
	users := make([]entity.User, 0)
	err := repository.database.Find(&users, conditions...).Error
	return users, err
}

func (repository *repository) Save(user *entity.User) error {
	return repository.database.Save(user).Error
}

func (repository *repository) Create(user *entity.User) error {
	return repository.database.Create(user).Error
}

func (repository *repository) Delete(user *entity.User) error {
	return repository.database.Delete(user).Error
}

func (repository *repository) Migrate() error {
	return repository.database.AutoMigrate(
		&entity.Capability{},
		&entity.Permission{},
		&entity.Role{},
		&entity.User{},
	)
}
