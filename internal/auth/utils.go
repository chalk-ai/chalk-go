package auth

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var authConfigFileName = ".chalk.yml"

func getDefaultAuthConfig() ProjectTokens {
	m := make(map[string]*ProjectToken)
	return ProjectTokens{
		Tokens: &m,
	}
}

func GetConfigPath() (*string, error) {
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

func LoadAuthConfig() (*ProjectTokens, *string, error) {
	path, err := GetConfigPath()
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
	return SourcedConfig{Source: "default"}
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

	configPath, err := GetConfigPath()
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
