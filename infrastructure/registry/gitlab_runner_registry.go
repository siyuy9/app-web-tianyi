package registry

import (
	"gorm.io/gorm"

	"gitlab.com/kongrentian-group/tianyi/entity"
)

type GitlabRunnerRegistry interface {
	FindOne(condition *entity.GitlabRunner) (*entity.GitlabRunner, error)

	Save(runner *entity.GitlabRunner) error
}

type gitlabRunnerRegistry struct {
	database *gorm.DB
}

func NewGitlabRunnerRegistry(database *gorm.DB) GitlabRunnerRegistry {
	return &gitlabRunnerRegistry{database: database}
}

func (registry *gitlabRunnerRegistry) FindOne(
	condition *entity.GitlabRunner,
) (*entity.GitlabRunner, error) {
	var model *entity.GitlabRunner
	return model, registry.database.Where(condition).First(model).Error
}

func (registry *gitlabRunnerRegistry) Save(runner *entity.GitlabRunner) error {
	return registry.database.Save(runner).Error
}
