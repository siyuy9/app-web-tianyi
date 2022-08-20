package entity

type PipelineConfig struct {
	Jobs      map[string]interface{} `hcl:"job,block"`
	Pipelines map[string]interface{} `hcl:"pipeline,block"`
}
