package config

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"gopkg.in/yaml.v3"
	"os"
)

func GetProjectAuthConfig(ctx context.Context, configDir *string) (*ProjectToken, string, error) {
	path, err := getConfigPath(ctx, configDir)
	if err != nil {
		return nil, "", errors.Wrap(err, "getting project path for auth config")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", fmt.Errorf("reading auth config file from path '%s': %w", path, err)
	}

	config := ProjectTokens{}
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, "", errors.Wrapf(
			err,
			"parsing auth config file at path '%s': make sure you have run 'chalk login' successfully.",
			path,
		)
	}

	projectAuthConfig, err := getProjectAuthConfigForProjectRoot(&config, path)
	return projectAuthConfig, path, errors.Wrap(err, "error getting project auth config")
}
