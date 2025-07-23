package config

import (
	"context"
	"github.com/cockroachdb/errors"
)

type Manager struct {
	ApiServer     SourcedConfig[string]
	ClientId      SourcedConfig[ClientId]
	ClientSecret  SourcedConfig[ClientSecret]
	EnvironmentId SourcedConfig[string]
	Scope         SourcedConfig[string]
}

// NewManager creates a new configuration manager with the specified parameters
func NewManager(
	ctx context.Context,
	apiServer SourcedConfig[string],
	clientId SourcedConfig[ClientId],
	clientSecret SourcedConfig[ClientSecret],
	environmentId SourcedConfig[string],
	configDir *string,
) (*Manager, error) {
	chalkYamlConfigOrNil, configPath, chalkYamlErr := GetProjectAuthConfig(ctx, configDir)
	chalkYamlConfig := ProjectToken{}
	if chalkYamlConfigOrNil != nil {
		chalkYamlConfig = *chalkYamlConfigOrNil
	}

	manager := &Manager{
		ApiServer: GetFirstNonEmpty(
			apiServer,
			NewFromEnvVar[string](ctx, ApiServerEnvVarKey),
			NewFromEnvVar[string](ctx, "_CHALK_API_SERVER"),
			NewFromFile(configPath, chalkYamlConfig.ApiServer),
		),
		ClientId: GetFirstNonEmpty(
			clientId,
			NewFromEnvVar[ClientId](ctx, ClientIdEnvVarKey),
			NewFromEnvVar[ClientId](ctx, "_CHALK_CLIENT_ID"),
			NewFromFile(configPath, chalkYamlConfig.ClientId),
		),
		ClientSecret: GetFirstNonEmpty(
			clientSecret,
			NewFromEnvVar[ClientSecret](ctx, ClientSecretEnvVarKey),
			NewFromEnvVar[ClientSecret](ctx, "_CHALK_CLIENT_SECRET"),
			NewFromFile(configPath, chalkYamlConfig.ClientSecret),
		),
		EnvironmentId: GetFirstNonEmpty(
			environmentId,
			NewFromEnvVar[string](ctx, EnvironmentEnvVarKey),
			NewFromEnvVar[string](ctx, "_CHALK_ACTIVE_ENVIRONMENT"),
			NewFromFile(configPath, chalkYamlConfig.ActiveEnvironment),
		),
	}

	if chalkYamlErr != nil && manager.ClientId.Value == "" && manager.ClientSecret.Value == "" {
		return nil, errors.Wrap(
			chalkYamlErr,
			"could not read chalk.yml and no client ID and client secret were provided",
		)
	}

	return manager, nil
}
