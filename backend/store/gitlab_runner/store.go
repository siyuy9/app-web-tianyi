package gitlab_runner

import (
	"gorm.io/gorm"

	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model/db"
)

type Store interface {
	// return a model based on the condition
	// 	gitlabRunnerModel, err := FindOneUser(&GitlabRunnerModel{Username: "username0"})
	Find(condition interface{}) (*db.GitlabRunnerModel, error)

	// save update value in database
	// if the value doesn't have a primary key, it will be inserted
	// 	err := SaveOne(&userModel)
	Save(data interface{}) error
}

type store struct {
	database *gorm.DB
}

func NewStore(database *gorm.DB) Store {
	return &store{database: database}
}

func (store *store) Find(
	condition interface{},
) (*db.GitlabRunnerModel, error) {
	var model *db.GitlabRunnerModel
	return model, store.database.Where(condition).First(model).Error
}

func (store *store) Save(data interface{}) error {
	return store.database.Save(data).Error
}
