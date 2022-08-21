package entity

type PipelineConfig struct {
	Jobs      []PipelineConfigJob      `hcl:"job,block" json:"jobs"`
	Pipelines []PipelineConfigPipeline `hcl:"pipeline,block" json:"pipelines"`
}

type PipelineConfigJob struct {
	Name        string            `hcl:"name,label" json:"name"`
	RequestType string            `hcl:"request_type, label" json:"request_type"`
	URL         string            `hcl:"url" json:"url"`
	Query       map[string]string `hcl:"query" json:"query"`
}

type PipelineConfigPipelineJob struct {
	Name string `hcl:"name,label" json:"name"`
}

type PipelineConfigPipeline struct {
	Name string                      `hcl:"name,label" json:"name"`
	Jobs []PipelineConfigPipelineJob `hcl:"job,block" json:"jobs"`
}
