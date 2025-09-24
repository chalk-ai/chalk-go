package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/cockroachdb/errors"
)

type Manager struct {
	ApiServer       SourcedConfig[string]
	GRPCQueryServer SourcedConfig[string]
	JSONQueryServer SourcedConfig[string]
	ClientId        SourcedConfig[ClientId]
	ClientSecret    SourcedConfig[ClientSecret]
	EnvironmentId   SourcedConfig[string]
	Scope           SourcedConfig[string]
}

type ManagerInputs struct {
	ApiServer       string
	GRPCQueryServer string
	JSONQueryServer string
	ClientId        ClientId
	ClientSecret    ClientSecret
	EnvironmentId   string
	Scope           string
	ConfigDir       *string
}

func NewManager(ctx context.Context, inputs *ManagerInputs) (*Manager, error) {
	chalkYamlConfigOrNil, configPath, chalkYamlErr := GetProjectAuthConfig(ctx, inputs.ConfigDir)
	chalkYamlConfig := ProjectToken{}
	if chalkYamlConfigOrNil != nil {
		chalkYamlConfig = *chalkYamlConfigOrNil
	}

	addScheme := func(c SourcedConfig[string]) SourcedConfig[string] {
		if strings.HasPrefix(c.Value, "http://") || strings.HasPrefix(c.Value, "https://") || c.Value == "" {
			return c
		}
		return c.WithValue(fmt.Sprintf("https://%s", c.Value))
	}

	manager := &Manager{
		ApiServer: addScheme(GetFirstNonEmpty(
			NewFromArg(inputs.ApiServer),
			NewFromEnvVar[string](ctx, "CHALK_API_SERVER"),
			NewFromEnvVar[string](ctx, "_CHALK_API_SERVER"),
			NewFromFile(configPath, chalkYamlConfig.ApiServer),
			NewFromDefault("https://api.chalk.ai", "default server"),
		)),
		JSONQueryServer: addScheme(GetFirstNonEmpty(
			NewFromArg(inputs.JSONQueryServer),
			NewFromDefault("https://api.chalk.ai", "default server"),
		)),
		GRPCQueryServer: addScheme(GetFirstNonEmpty(
			NewFromArg(inputs.GRPCQueryServer),
			NewFromDefault("https://api.chalk.ai", "default server"),
		)),
		ClientId: GetFirstNonEmpty(
			NewFromArg(inputs.ClientId),
			NewFromEnvVar[ClientId](ctx, "CHALK_CLIENT_ID"),
			NewFromEnvVar[ClientId](ctx, "_CHALK_CLIENT_ID"),
			NewFromFile(configPath, chalkYamlConfig.ClientId),
		),
		ClientSecret: GetFirstNonEmpty(
			NewFromArg(inputs.ClientSecret),
			NewFromEnvVar[ClientSecret](ctx, "CHALK_CLIENT_SECRET"),
			NewFromEnvVar[ClientSecret](ctx, "_CHALK_CLIENT_SECRET"),
			NewFromFile(configPath, chalkYamlConfig.ClientSecret),
		),
		EnvironmentId: GetFirstNonEmpty(
			NewFromArg(inputs.EnvironmentId),
			NewFromEnvVar[string](ctx, "CHALK_ACTIVE_ENVIRONMENT"),
			NewFromEnvVar[string](ctx, "_CHALK_ACTIVE_ENVIRONMENT"),
			NewFromFile(configPath, chalkYamlConfig.ActiveEnvironment),
		),
		Scope: GetFirstNonEmpty(NewFromArg(inputs.Scope)),
	}

	if manager.ClientId.Value == "" || manager.ClientSecret.Value == "" {
		if chalkYamlErr != nil {
			return nil, errors.Wrap(
				chalkYamlErr,
				"could not read chalk.yml and no client ID and client secret were provided",
			)
		}
		return nil, errors.Newf("could not find values for client id and client secret")
	}

	return manager, nil
}
