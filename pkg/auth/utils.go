package auth

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func getDefaultAuthConfig() AuthConfig {
	m := make(map[string]*TokenConfig)
	return AuthConfig{
		Tokens:       &m,
		VersionCheck: nil,
	}
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

func LoadAuthConfig() AuthConfig {
	path, err := getConfigPath()
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
