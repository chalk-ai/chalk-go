package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chalk-ai/chalk-go/internal"
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

func getConfigPath(ctx context.Context, configDir *string) (string, error) {
	var dir string
	var err error
	getter := EnvironmentGetterFromContext(ctx)

	if configDir != nil {
		dir = *configDir
	} else {
		dir = getter.Getenv("XDG_CONFIG_HOME")
		if dir == "" {
			dir, err = os.UserHomeDir()
			if err != nil {
				return "", errors.New("error getting home directory")
			}
		}
	}

	// Check for both chalk.yml and chalk.yaml
	for _, filename := range []string{".chalk.yml", ".chalk.yaml"} {
		path := filepath.Join(dir, filename)
		if internal.FileExists(path) {
			return path, nil
		}
	}

	// Default to .chalk.yml if neither exists
	path := filepath.Join(dir, authConfigFileName)
	return path, nil
}

func getProjectAuthConfigForProjectRoot(config *ProjectTokens, configPath string) (*ProjectToken, error) {
	projectRoot, err := loadProjectDirectory()
	if err != nil {
		return nil, fmt.Errorf("error loading auth config: %s", err)
	}

	if config.Tokens == nil {
		return nil, fmt.Errorf(
			"'tokens' collection does not exist or is empty in the auth config file "+
				"'%s' -- please try to 'chalk login' again",
			configPath,
		)
	}

	var returnToken *ProjectToken

	tokens := config.Tokens
	if token, ok := tokens[projectRoot]; ok {
		returnToken = token
	}

	if token, ok := tokens["default"]; ok && returnToken == nil {
		returnToken = token
	}

	if returnToken == nil {
		return nil, fmt.Errorf(
			"project root '%s' does not exist as a key in the collection 'tokens'"+
				" in the config file '%s', and the fallback key 'default' is also missing. "+
				"Please try to 'chalk login' again",
			projectRoot,
			configPath,
		)
	}

	return returnToken, nil
}
