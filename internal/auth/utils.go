package auth

import (
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
			return nil, err
		}
	}
	path := filepath.Join(configDir, authConfigFileName)
	return &path, nil
}

func LoadAuthConfig() ProjectTokens {
	path, err := GetConfigPath()
	if err != nil || path == nil {
		return getDefaultAuthConfig()
	}

	data, err := os.ReadFile(*path)

	if err != nil {
		return getDefaultAuthConfig()
	}

	config := ProjectTokens{}
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return getDefaultAuthConfig()
	}

	return config
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
		"New argument",
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
