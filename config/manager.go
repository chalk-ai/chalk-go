package config

import (
	"context"
)

type Manager struct {
	apiServer       SourcedConfig[string]
	grpcQueryServer SourcedConfig[string]
	jsonQueryServer SourcedConfig[string]
	ClientId        SourcedConfig[ClientId]
	ClientSecret    SourcedConfig[ClientSecret]
	EnvironmentId   SourcedConfig[string]
	Scope           SourcedConfig[string]

	// projectConfigErr records why chalk.yml could not be read, if it could not.
	// Resolution proceeds regardless: credentials may come from a flag or the
	// environment, or a pre-issued JWT may make them unnecessary.
	projectConfigErr error
}

// ProjectConfigErr is advisory. A non-nil result does not mean the resolved
// config is unusable -- credentials may have come from a flag or the
// environment, or a pre-issued JWT may make them unnecessary. It is meaningful
// only once authentication has found nothing to authenticate with, as the
// likely explanation for why.
func (m *Manager) ProjectConfigErr() error {
	return m.projectConfigErr
}

type ManagerInputs struct {
	APIServer       string
	GRPCQueryServer string
	JSONQueryServer string
	ClientId        ClientId
	ClientSecret    ClientSecret
	EnvironmentId   string
	Scope           string
	ConfigDir       *string
}

func (m *Manager) GetAPIServer() SourcedConfig[string] {
	return m.apiServer
}
func (m *Manager) SetAPIServer(server SourcedConfig[string]) {
	m.apiServer = AddScheme(server)
}

func (m *Manager) GetGRPCQueryServer() SourcedConfig[string] {
	return m.grpcQueryServer
}
func (m *Manager) SetGRPCQueryServer(server SourcedConfig[string]) {
	m.grpcQueryServer = AddScheme(server)
}

func (m *Manager) GetJSONQueryServer() SourcedConfig[string] {
	return m.jsonQueryServer
}
func (m *Manager) SetJSONQueryServer(server SourcedConfig[string]) {
	m.jsonQueryServer = AddScheme(server)
}

func NewManager(ctx context.Context, inputs *ManagerInputs) (*Manager, error) {
	// A failed read is not fatal: credentials may come from a flag or the
	// environment, or a pre-issued JWT may make them unnecessary entirely. The
	// error is retained on the Manager so auth can report it if authentication
	// then turns out to have nothing to use.
	chalkYamlConfigOrNil, configPath, chalkYamlErr := GetProjectAuthConfig(ctx, inputs.ConfigDir)
	chalkYamlConfig := ProjectToken{}
	if chalkYamlConfigOrNil != nil {
		chalkYamlConfig = *chalkYamlConfigOrNil
	}

	manager := &Manager{
		apiServer: AddScheme(GetFirstNonEmpty(
			NewFromArg(inputs.APIServer),
			NewFromEnvVar[string](ctx, "CHALK_API_SERVER"),
			NewFromEnvVar[string](ctx, "_CHALK_API_SERVER"),
			NewFromFile(configPath, chalkYamlConfig.ApiServer),
			NewFromDefault("https://api.chalk.ai", "default server"),
		)),
		jsonQueryServer: AddScheme(GetFirstNonEmpty(
			NewFromArg(inputs.JSONQueryServer),
			NewFromDefault("https://api.chalk.ai", "default server"),
		)),
		grpcQueryServer: AddScheme(GetFirstNonEmpty(
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
		Scope:            GetFirstNonEmpty(NewFromArg(inputs.Scope)),
		projectConfigErr: chalkYamlErr,
	}

	return manager, nil
}
