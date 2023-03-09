package auth

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var AuthConfigFileName = ".chalk.yml"

func getDefaultAuthConfig() authConfig {
	m := make(map[string]*projectAuthConfig)
	return authConfig{
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
	path := filepath.Join(configDir, AuthConfigFileName)
	return &path, nil
}

func LoadAuthConfig() authConfig {
	path, err := GetConfigPath()
	if err != nil || path == nil {
		return getDefaultAuthConfig()
	}

	data, err := os.ReadFile(*path)

	if err != nil {
		return getDefaultAuthConfig()
	}

	config := authConfig{}
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return getDefaultAuthConfig()
	}

	return config
}
