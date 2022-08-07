package registry

import (
	"gorm.io/gorm"

	"gitlab.com/kongrentian-group/tianyi/entity"
)

type UserRegistry interface {
	// return all users in the database
	FindAll() ([]entity.User, error)
	// return a user based on the condition
	// repository.FindOne(&model.User{Username: "username0"})
	FindOne(condition *entity.User) (*entity.User, error)
	// return an array of users based on the conditions
	// https://gorm.io/docs/query.html
	Find(conditions ...interface{}) ([]entity.User, error)

	// save a given user (if it doesn't exist, create)
	Save(user *entity.User) error
	// delete a given user
	Delete(user *entity.User) error

	// automatically migrate userRegistry tables
	// https://gorm.io/docs/migration.html
	// NOTE: AutoMigrate will create tables, missing foreign keys, constraints,
	// columns and indexes
	// It will change existing column’s type if its size, precision, nullable
	// changed
	// It WON’T delete unused columns to protect your data
	Migrate() error
}

type userRegistry struct {
	database *gorm.DB
}

func NewUserRegistry(database *gorm.DB) UserRegistry {
	return &userRegistry{database: database}
}

func (repository *userRegistry) FindAll() ([]entity.User, error) {
	return repository.Find()
}

func (repository *userRegistry) FindOne(
	condition *entity.User,
) (user *entity.User, err error) {
	err = repository.database.First(&condition).Error
	return
}

func (repository *userRegistry) Find(
	conditions ...interface{},
) ([]entity.User, error) {
	users := make([]entity.User, 0)
	err := repository.database.Find(users, conditions...).Error
	return users, err
}

func (repository *userRegistry) Save(user *entity.User) error {
	return repository.database.Save(user).Error
}

func (repository *userRegistry) Delete(user *entity.User) error {
	return repository.database.Delete(user).Error
}

func (repository *userRegistry) Migrate() error {
	return repository.database.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Permission{},
		&entity.Capability{},
	)
}
