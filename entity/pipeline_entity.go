package entity

type PipelineConfig struct {
	Jobs      []PipelineConfigJob      `hcl:"job,block"`
	Pipelines []PipelineConfigPipeline `hcl:"pipeline,block"`
}

type PipelineConfigJob struct {
	Name string `hcl:"name,label"`
}

type PipelineConfigPipeline struct {
	Name string `hcl:"name,label"`
}
