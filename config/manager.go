package config

import (
	"context"

	"github.com/cockroachdb/errors"
)

type Manager struct {
	ApiServer       *SourcedConfig[string]
	GRPCQueryServer *SourcedConfig[string]
	JSONQueryServer *SourcedConfig[string]
	ClientId        *SourcedConfig[ClientId]
	ClientSecret    *SourcedConfig[ClientSecret]
	EnvironmentId   *SourcedConfig[string]
	Scope           *SourcedConfig[string]
}

type ManagerInputs struct {
	ApiServer       *SourcedConfig[string]
	GRPCQueryServer *SourcedConfig[string]
	JSONQueryServer *SourcedConfig[string]
	ClientId        *SourcedConfig[ClientId]
	ClientSecret    *SourcedConfig[ClientSecret]
	EnvironmentId   *SourcedConfig[string]
	Scope           *SourcedConfig[string]
	ConfigDir       *string
}

func NewManager(ctx context.Context, inputs *ManagerInputs) (*Manager, error) {
	chalkYamlConfigOrNil, configPath, chalkYamlErr := GetProjectAuthConfig(ctx, inputs.ConfigDir)
	chalkYamlConfig := ProjectToken{}
	if chalkYamlConfigOrNil != nil {
		chalkYamlConfig = *chalkYamlConfigOrNil
	}

	manager := &Manager{
		ApiServer: GetFirstNonEmpty(
			inputs.ApiServer,
			NewFromEnvVar[string](ctx, ApiServerEnvVarKey),
			NewFromEnvVar[string](ctx, "_CHALK_API_SERVER"),
			NewFromFile(configPath, chalkYamlConfig.ApiServer),
			NewFromDefault("https://api.chalk.ai", "default server"),
		),
		JSONQueryServer: GetFirstNonEmpty(
			inputs.JSONQueryServer,
			NewFromDefault("https://api.chalk.ai", "default server"),
		),
		GRPCQueryServer: GetFirstNonEmpty(
			inputs.GRPCQueryServer,
			NewFromDefault("https://api.chalk.ai", "default server"),
		),
		ClientId: GetFirstNonEmpty(
			inputs.ClientId,
			NewFromEnvVar[ClientId](ctx, ClientIdEnvVarKey),
			NewFromEnvVar[ClientId](ctx, "_CHALK_CLIENT_ID"),
			NewFromFile(configPath, chalkYamlConfig.ClientId),
		),
		ClientSecret: GetFirstNonEmpty(
			inputs.ClientSecret,
			NewFromEnvVar[ClientSecret](ctx, ClientSecretEnvVarKey),
			NewFromEnvVar[ClientSecret](ctx, "_CHALK_CLIENT_SECRET"),
			NewFromFile(configPath, chalkYamlConfig.ClientSecret),
		),
		EnvironmentId: GetFirstNonEmpty(
			inputs.EnvironmentId,
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
