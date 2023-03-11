package auth

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var authConfigFileName = ".chalk.yml"

func getDefaultAuthConfig() AuthConfig {
	m := make(map[string]*ProjectAuthConfig)
	return AuthConfig{
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

func LoadAuthConfig() AuthConfig {
	path, err := GetConfigPath()
	if err != nil || path == nil {
		return getDefaultAuthConfig()
	}

	data, err := os.ReadFile(*path)

	if err != nil {
		return getDefaultAuthConfig()
	}

	config := AuthConfig{}
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return getDefaultAuthConfig()
	}

	return config
}
