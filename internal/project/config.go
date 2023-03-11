package project

import (
	"errors"
	"github.com/chalk-ai/chalk-go/internal"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var defaultRequirements = "requirements.txt"
var chalkIgnore = ".chalkignore"

func LoadProjectConfig() (*ProjectSettings, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	currentDirectory := wd
	rootChecked := false
	for currentDirectory != "/" && !rootChecked {
		for _, filename := range []string{"chalk.yaml", "chalk.yml"} {
			settings, err := checkDirectory(currentDirectory, filename)
			if settings != nil && err == nil {
				return settings, err
			}
		}

		if currentDirectory == "/" {
			rootChecked = true
		} else {
			currentDirectory = filepath.Dir(currentDirectory)
		}
	}

	return nil, errors.New("Failed to find chalk.yml in any directory.")
}

func checkDirectory(directory string, filename string) (*ProjectSettings, error) {
	configFilename := filepath.Join(directory, filename)
	if !internal.FileExists(configFilename) {
		return nil, nil
	}

	hasDefaultRequirements := false
	defaultRequirementsFilename := filepath.Join(directory, defaultRequirements)
	if internal.FileExists(defaultRequirementsFilename) {
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
			env.Requirements = &defaultRequirements
		}
	}

	chalkignoreFilename := filepath.Join(directory, chalkIgnore)
	if internal.FileExists(chalkignoreFilename) {
		settings.ChalkIgnore = &chalkignoreFilename
	}

	return &settings, err
}