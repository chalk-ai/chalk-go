package auth

import (
	"github.com/cockroachdb/errors"
)

func GetProjectAuthConfig() (ProjectToken, error) {
	authConfigFromFile, configFilepath, err := loadAuthConfig()
	if err != nil {
		return ProjectToken{}, errors.Wrap(err, "getting project auth config")
	}
	if configFilepath == nil {
		return ProjectToken{}, errors.New("unexpected error getting auth config filepath from home directory")
	}
	if authConfigFromFile == nil {
		return ProjectToken{}, errors.New("unexpected error getting auth config file")
	}

	projectAuthConfig, err := getProjectAuthConfigForProjectRoot(*authConfigFromFile, *configFilepath)
	return projectAuthConfig, errors.Wrap(err, "error getting project auth config")
}
