package project

type EnvironmentSettings struct {
	Runtime      *string `json:"runtime" yaml:"runtime,omitempty"`
	Requirements *string `json:"requirements" yaml:"requirements,omitempty"`
	Dockerfile   *string `json:"dockerfile" yaml:"dockerfile,omitempty"`
}

type ProjectSettings struct {
	Project        string                          `json:"project" yaml:"project"`
	Environments   map[string]*EnvironmentSettings `json:"environments" yaml:"environments"`
	LocalDirectory string                          `json:"localDirectory" yaml:"localDirectory,omitempty"`
	ChalkIgnore    *string                         `json:"chalkIgnore" yaml:"chalkIgnore,omitempty"`
	Filename       string                          `yaml:"filename,omitempty"`
}
