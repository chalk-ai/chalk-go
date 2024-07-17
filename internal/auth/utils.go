package auth

import (
	"errors"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var authConfigFileName = ".chalk.yml"

func loadProjectDirectory() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	currentDirectory := wd
	rootChecked := false
	for currentDirectory != "/" && !rootChecked {
		for _, filename := range []string{"chalk.yaml", "chalk.yml"} {
			chalkYamlExists := internal.FileExists(filepath.Join(currentDirectory, filename))
			if chalkYamlExists {
				return currentDirectory, nil
			}
		}

		if currentDirectory == "/" {
			rootChecked = true
		} else {
			currentDirectory = filepath.Dir(currentDirectory)
		}
	}

	return "", fmt.Errorf("cannot determine project root directory: "+
		"failed to find chalk.yml in the working directory '%s' or any of its parent directories", wd,
	)
}

func getConfigPath() (*string, error) {
	var err error
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir, err = os.UserHomeDir()
		if err != nil {
			return nil, errors.New("error getting home directory")
		}
	}
	path := filepath.Join(configDir, authConfigFileName)
	return &path, nil
}

func getProjectAuthConfigForProjectRoot(config ProjectTokens, configPath string) (ProjectToken, error) {
	projectRoot, err := loadProjectDirectory()
	if err != nil {
		return ProjectToken{}, fmt.Errorf("error loading auth config: %s", err)
	}

	if config.Tokens == nil {
		return ProjectToken{}, errors.New(
			fmt.Sprintf("'tokens' collection does not exist or is empty in the auth config file '%s' -- please try to 'chalk login' again", configPath))
	}

	var returnToken *ProjectToken = nil

	tokens := *config.Tokens
	if token, ok := tokens[projectRoot]; ok {
		returnToken = token
	}

	if token, ok := tokens["default"]; ok && returnToken == nil {
		returnToken = token
	}

	if returnToken == nil {
		return ProjectToken{}, errors.New(fmt.Sprintf("project root '%s' does not exist as a key in the collection 'tokens' in the config file '%s', and the fallback key 'default' is also missing. Please try to 'chalk login' again", projectRoot, configPath))
	}

	return *returnToken, nil
}

func loadAuthConfig() (*ProjectTokens, *string, error) {
	path, err := getConfigPath()
	if err != nil || path == nil {
		return nil, nil, err
	}

	data, err := os.ReadFile(*path)
	if err != nil {
		return nil, path, errors.New(fmt.Sprintf("Error reading auth config file from path '%s': %s", *path, err))
	}

	config := ProjectTokens{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, path, errors.New(fmt.Sprintf("Error parsing auth config file at path '%s'. Please make sure you have run 'chalk login' successfully. Error details: %s", *path, err))
	}

	return &config, path, nil
}

func GetFirstNonEmptyConfig(configs ...SourcedConfig) SourcedConfig {
	for _, config := range configs {
		if config.Value != "" {
			return config
		}
	}
	return SourcedConfig{Source: "value in '~/.chalk.yml' does not exist or is empty"}
}

func GetEnvVarConfig(key string) SourcedConfig {
	return SourcedConfig{
		os.Getenv(key),
		fmt.Sprintf("environment variable '%s'", key),
	}
}

func GetChalkClientArgConfig(value string) SourcedConfig {
	return SourcedConfig{
		value,
		"NewClient argument",
	}
}

func GetChalkYamlConfig(value string) SourcedConfig {
	var path string

	configPath, err := getConfigPath()
	if err != nil {
		path = "unknown"
	} else {
		path = *configPath
	}

	return SourcedConfig{
		value,
		fmt.Sprintf("config file %s", path),
	}
}

type SourcedConfig struct {
	Value  string
	Source string
}

type ResolvedConfig struct {
	EnvironmentId SourcedConfig
	ClientId      SourcedConfig
	ClientSecret  SourcedConfig
	ApiServer     SourcedConfig
}
