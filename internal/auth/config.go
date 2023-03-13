package auth

import (
	"errors"
)

func GetProjectAuthConfig() (ProjectToken, error) {
	authConfigFromFile, configFilepath, loadConfigFileErr := loadAuthConfig()
	if loadConfigFileErr != nil {
		return ProjectToken{}, loadConfigFileErr
	}
	if configFilepath == nil {
		return ProjectToken{}, errors.New("unexpected error getting auth config filepath from home directory")
	}
	if authConfigFromFile == nil {
		return ProjectToken{}, errors.New("unexpected error getting auth config file")
	}

	projectAuthConfig, projectAuthConfigErr := getProjectAuthConfigForWD(*authConfigFromFile, *configFilepath)
	if projectAuthConfigErr != nil {
		return ProjectToken{}, projectAuthConfigErr
	}

	return projectAuthConfig, nil
}
