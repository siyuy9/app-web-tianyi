package entity

type PipelineConfig struct {
	Jobs      map[string]interface{} `hcl:"jobs"`
	Pipelines map[string]interface{} `hcl:""`
}
