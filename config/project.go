package config

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/chalk-ai/chalk-go/internal/ptr"
	"gopkg.in/yaml.v3"
)

func getDefaultConfig(project string) ProjectSettings {
	return ProjectSettings{
		Project: project,
		Environments: map[string]*EnvironmentSettings{
			"default": {
				Runtime:      ptr.New("python312"),
				Dockerfile:   nil,
				Requirements: ptr.New("requirements.txt"),
			},
		},
	}
}

var DefaultRequirements = "requirements.txt"
var ChalkIgnore = ".chalkignore"

func checkDirectory(directory string, filename string) (*ProjectSettings, error) {
	configFilename := filepath.Join(directory, filename)
	if _, err := os.Stat(configFilename); err != nil {
		return nil, nil
	}

	hasDefaultRequirements := false
	defaultRequirementsFilename := filepath.Join(directory, DefaultRequirements)
	if _, err := os.Stat(defaultRequirementsFilename); err == nil {
		hasDefaultRequirements = true
	}

	yfile, err := os.ReadFile(configFilename)
	if err != nil {
		return nil, err
	}

	settings := ProjectSettings{
		LocalDirectory: directory,
		Filename:       configFilename,
	}
	err = yaml.Unmarshal(yfile, &settings)
	if err != nil {
		return nil, err
	}
	for _, env := range settings.Environments {
		if (env.Requirements == nil || *env.Requirements == "") && hasDefaultRequirements {
			env.Requirements = &DefaultRequirements
		}
	}

	chalkIgnoreFilename := filepath.Join(directory, ChalkIgnore)
	if _, err := os.Stat(chalkIgnoreFilename); err == nil {
		settings.ChalkIgnore = &chalkIgnoreFilename
	}

	return &settings, err
}

func LoadProjectConfig(ctx context.Context) (*ProjectSettings, error) {
	currentDirectory, err := EnvironmentGetterFromContext(ctx).Getwd()
	if err != nil {
		return nil, err
	}
	for lastCheckedDirectory := ""; lastCheckedDirectory != currentDirectory; currentDirectory = filepath.Dir(currentDirectory) {
		for _, filename := range []string{"chalk.yaml", "chalk.yml"} {
			if settings, err := checkDirectory(currentDirectory, filename); settings != nil && err == nil {
				return settings, nil
			}
		}
		lastCheckedDirectory = currentDirectory
	}
	return nil, errors.New("no chalk.yaml or chalk.yml in this directory or any parent directory")
}
