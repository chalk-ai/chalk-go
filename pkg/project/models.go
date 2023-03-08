package project

type environmentSettings struct {
	Runtime      *string `json:"runtime" yaml:"runtime,omitempty"`
	Requirements *string `json:"requirements" yaml:"requirements,omitempty"`
	Dockerfile   *string `json:"dockerfile" yaml:"dockerfile,omitempty"`
}

type projectSettings struct {
	Project        string                          `json:"project" yaml:"project"`
	Environments   map[string]*environmentSettings `json:"environments" yaml:"environments"`
	LocalDirectory string                          `json:"localDirectory" yaml:"localDirectory,omitempty"`
	ChalkIgnore    *string                         `json:"chalkIgnore" yaml:"chalkIgnore,omitempty"`
	Filename       string                          `yaml:"filename,omitempty"`
}
