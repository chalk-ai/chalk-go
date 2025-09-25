package config

import (
	"context"
	"fmt"

	"github.com/chalk-ai/chalk-go/envfs"
	"github.com/cockroachdb/errors"
	"gopkg.in/yaml.v3"
)

func GetProjectAuthConfig(ctx context.Context, configDir *string) (*ProjectToken, string, error) {
	path, err := getConfigPath(ctx, configDir)
	if err != nil {
		return nil, "", errors.Wrap(err, "getting project path for auth config")
	}

	getter := envfs.EnvironmentGetterFromContext(ctx)
	data, err := getter.ReadFile(path)
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

	projectAuthConfig, err := getProjectAuthConfigForProjectRoot(&config, path, getter)
	return projectAuthConfig, path, errors.Wrap(err, "error getting project auth config")
}
