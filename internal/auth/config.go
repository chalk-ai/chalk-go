package auth

import (
	"errors"
	"fmt"
	"os"
)

func GetProjectAuthConfig() (ProjectToken, error) {
	authConfigFromFile, configFilepath, loadConfigFileErr := LoadAuthConfig()
	if loadConfigFileErr != nil {
		return ProjectToken{}, loadConfigFileErr
	}
	if configFilepath == nil {
		return ProjectToken{}, errors.New("unexpected error getting auth config filepath from home directory")
	}
	if authConfigFromFile == nil {
		return ProjectToken{}, errors.New("unexpected error getting auth config file")
	}

	projectAuthConfig, projectAuthConfigErr := GetProjectAuthConfigForWD(*authConfigFromFile, *configFilepath)
	if projectAuthConfigErr != nil {
		return ProjectToken{}, projectAuthConfigErr
	}

	return projectAuthConfig, nil
}

func GetProjectAuthConfigForWD(config ProjectTokens, configPath string) (ProjectToken, error) {
	getwd, err := os.Getwd()
	if err != nil {
		return ProjectToken{}, errors.New(fmt.Sprintf("error determining working directory: %s", err))
	}

	if config.Tokens == nil {
		return ProjectToken{}, errors.New(
			fmt.Sprintf("'tokens' collection does not exist or is empty in the auth config file '%s' -- please try to 'chalk login' again", configPath))
	}

	var returnToken *ProjectToken = nil

	tokens := *config.Tokens
	if token, ok := tokens[getwd]; ok {
		returnToken = token
	}

	if token, ok := tokens["default"]; ok && returnToken == nil {
		returnToken = token
	}

	if returnToken == nil {
		return ProjectToken{}, errors.New(fmt.Sprintf("working directory '%s' does not exist as a key in the collection 'tokens' in the config file '%s' The key 'default' is also missing. Please try to 'chalk login' again", getwd, configPath))
	}

	return *returnToken, nil
}
