package config

import (
	"time"

	artifactsv1 "github.com/chalk-ai/chalk-go/gen/chalk/artifacts/v1"
)

type ClientId string
type ClientSecret string
type EnvironmentId string

// ~/.chalk.yml contents

type JWT struct {
	Token      string    `yaml:"value"`
	ValidUntil time.Time `yaml:"validUntil"`
}

type ProjectToken struct {
	Name              string       `yaml:"name"`
	ClientId          ClientId     `yaml:"clientId"`
	ClientSecret      ClientSecret `yaml:"clientSecret"`
	ValidUntil        string       `yaml:"validUntil"`
	ApiServer         string       `yaml:"apiServer"`
	ActiveEnvironment string       `yaml:"activeEnvironment"`
	JWT               JWT          `yaml:"jwt,omitempty"`
}

type ProjectTokens struct {
	Tokens map[string]*ProjectToken `yaml:"tokens,omitempty"`
}

// Project-level chalk.yml contents

type EnvironmentSettings struct {
	Runtime         *string `json:"runtime,omitempty" yaml:"runtime,omitempty"`
	PlatformVersion *string `json:"platform_version,omitempty" yaml:"platform_version,omitempty"`
	Requirements    *string `json:"requirements,omitempty" yaml:"requirements,omitempty"`
	Dockerfile      *string `json:"dockerfile,omitempty" yaml:"dockerfile,omitempty"`
}

func (e *EnvironmentSettings) ToProto(id string) *artifactsv1.EnvironmentSettings {
	return &artifactsv1.EnvironmentSettings{
		Id:              id,
		Runtime:         e.Runtime,
		Requirements:    e.Requirements,
		Dockerfile:      e.Dockerfile,
		PlatformVersion: e.PlatformVersion,
	}
}

type ProjectSettings struct {
	Project        string                          `json:"project" yaml:"project"`
	Environments   map[string]*EnvironmentSettings `json:"environments" yaml:"environments"`
	LocalDirectory string                          `json:"localDirectory" yaml:"localDirectory,omitempty"`
	ChalkIgnore    *string                         `json:"chalkIgnore,omitempty" yaml:"chalkIgnore,omitempty"`
	Filename       string                          `yaml:"filename,omitempty"`
}

func (settings *ProjectSettings) ToProto() *artifactsv1.ProjectSettings {
	var environments []*artifactsv1.EnvironmentSettings
	for name, env := range settings.Environments {
		environments = append(environments, env.ToProto(name))
	}
	return &artifactsv1.ProjectSettings{
		Project:      settings.Project,
		Environments: environments,
		// TODO: Add validations
	}
}
