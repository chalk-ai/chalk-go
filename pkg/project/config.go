package project

import (
	"errors"
	"github.com/chalk-ai/chalk-go/pkg/utils"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var AuthConfigFileName = ".chalk.yml"
var DEFAULT_REQUIREMENTS = "requirements.txt"
var CHALKIGNORE = ".chalkignore"

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

func getConfigPath() (*string, error) {
	var err error
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir, err = os.UserHomeDir()
		if err != nil {
			return nil, err
		}
	}
	path := filepath.Join(configDir, AuthConfigFileName)
	return &path, nil
}

func checkDirectory(directory string, filename string) (*ProjectSettings, error) {
	configFilename := filepath.Join(directory, filename)
	if !utils.Exists(configFilename) {
		return nil, nil
	}

	hasDefaultRequirements := false
	defaultRequirementsFilename := filepath.Join(directory, DEFAULT_REQUIREMENTS)
	if utils.Exists(defaultRequirementsFilename) {
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
			env.Requirements = &DEFAULT_REQUIREMENTS
		}
	}

	chalkignoreFilename := filepath.Join(directory, CHALKIGNORE)
	if utils.Exists(chalkignoreFilename) {
		settings.ChalkIgnore = &chalkignoreFilename
	}

	return &settings, err
}
