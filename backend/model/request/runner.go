package request

// https://gitlab.com/gitlab-org/gitlab-runner/-/blob/main/common/network.go
type RegisterRunnerParameters struct {
	Description     string `json:"description,omitempty"`
	MaintenanceNote string `json:"maintenance_note,omitempty"`
	Tags            string `json:"tag_list,omitempty"`
	RunUntagged     bool   `json:"run_untagged"`
	Locked          bool   `json:"locked"`
	AccessLevel     string `json:"access_level,omitempty"`
	MaximumTimeout  int    `json:"maximum_timeout,omitempty"`
	Paused          bool   `json:"paused"`
}

type RegisterRunnerRequest struct {
	RegisterRunnerParameters
	Info  VersionInfo `json:"info,omitempty"`
	Token string      `json:"token,omitempty"`
}

type VersionInfo struct {
	Name         string       `json:"name,omitempty"`
	Version      string       `json:"version,omitempty"`
	Revision     string       `json:"revision,omitempty"`
	Platform     string       `json:"platform,omitempty"`
	Architecture string       `json:"architecture,omitempty"`
	Executor     string       `json:"executor,omitempty"`
	Shell        string       `json:"shell,omitempty"`
	Features     FeaturesInfo `json:"features"`
	Config       ConfigInfo   `json:"config,omitempty"`
}

type FeaturesInfo struct {
	Variables               bool `json:"variables"`
	Image                   bool `json:"image"`
	Services                bool `json:"services"`
	Artifacts               bool `json:"artifacts"`
	Cache                   bool `json:"cache"`
	Shared                  bool `json:"shared"`
	UploadMultipleArtifacts bool `json:"upload_multiple_artifacts"`
	UploadRawArtifacts      bool `json:"upload_raw_artifacts"`
	Session                 bool `json:"session"`
	Terminal                bool `json:"terminal"`
	Refspecs                bool `json:"refspecs"`
	Masking                 bool `json:"masking"`
	Proxy                   bool `json:"proxy"`
	RawVariables            bool `json:"raw_variables"`
	ArtifactsExclude        bool `json:"artifacts_exclude"`
	MultiBuildSteps         bool `json:"multi_build_steps"`
	TraceReset              bool `json:"trace_reset"`
	TraceChecksum           bool `json:"trace_checksum"`
	TraceSize               bool `json:"trace_size"`
	VaultSecrets            bool `json:"vault_secrets"`
	Cancelable              bool `json:"cancelable"`
	ReturnExitCode          bool `json:"return_exit_code"`
	ServiceVariables        bool `json:"service_variables"`
}

type ConfigInfo struct {
	Gpus string `json:"gpus"`
}
