package db

import (
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model/request"
	"gorm.io/gorm"
)

// https://gitlab.com/gitlab-org/gitlab/-/blob/master/db/structure.sql#L13130
type GitlabRunner struct {
	gorm.Model
	Token        string `json:"token,omitempty"`
	Name         string `json:"name,omitempty"`
	Version      string `json:"version,omitempty"`
	Revision     string `json:"revision,omitempty"`
	Platform     string `json:"platform,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Executor     string `json:"executor,omitempty"`
	Shell        string `json:"shell,omitempty"`
	request.FeaturesInfo
	request.ConfigInfo
}
