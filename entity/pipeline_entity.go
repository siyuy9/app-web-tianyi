package entity

import "github.com/google/uuid"

type Pipeline struct {
	CommonFields
	ProjectID uuid.UUID `gorm:"type:uuid;not null;" json:"project_id"`
	BranchID  uuid.UUID `gorm:"type:uuid;not null;" json:"branch_id"`
}

type Job struct {
	CommonFields
	PipelineID uuid.UUID `gorm:"type:uuid;not null;" json:"pipeline_id"`
	Result     bool      `json:"result"`
	Log        string    `json:"log"`
}

type PipelineConfig struct {
	Jobs      []PipelineConfigJob      `hcl:"job,block" json:"jobs"`
	Pipelines []PipelineConfigPipeline `hcl:"pipeline,block" json:"pipelines"`
}

type PipelineConfigJob struct {
	Name        string            `hcl:"name,label" json:"name"`
	RequestType string            `hcl:"request_type,label" json:"request_type"`
	URL         string            `hcl:"url" json:"url"`
	Query       map[string]string `hcl:"query" json:"query"`
}

type PipelineConfigPipelineJob struct {
	Name     string    `hcl:"name,label" json:"name"`
	Job      string    `hcl:"job,label" json:"job"`
	Requires *[]string `hcl:"requires" json:"requires"`
}

type PipelineConfigPipeline struct {
	Name string                      `hcl:"name,label" json:"name"`
	Jobs []PipelineConfigPipelineJob `hcl:"job,block" json:"jobs"`
}
