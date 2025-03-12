package chalk

import (
	"context"
	"github.com/chalk-ai/chalk-go/v2/internal"
	"github.com/chalk-ai/chalk-go/v2/internal/auth"
	"github.com/chalk-ai/chalk-go/v2/internal/colls"
	"github.com/cockroachdb/errors"
	"time"
)

// getTokenResult is agnostic to whether the token
// was obtained via gRPC or REST.
type getTokenResult struct {
	AccessToken        string
	PrimaryEnvironment string
	ValidUntil         time.Time
	Engines            map[string]string
}

type configManager struct {
	apiServer          auth.SourcedConfig
	clientId           auth.SourcedConfig
	clientSecret       auth.SourcedConfig
	environmentId      auth.SourcedConfig
	initialEnvironment auth.SourcedConfig

	jwt      *auth.JWT
	engines  map[string]string
	getToken func(ctx context.Context, clientId string, clientSecret string) (*getTokenResult, error)

	logger LeveledLogger
}

func newConfigManager(
	apiServer string,
	clientId string,
	clientSecret string,
	environmentId string,
	logger LeveledLogger,
) (*configManager, error) {
	chalkYamlConfigOrNil, chalkYamlErr := auth.GetProjectAuthConfig()
	chalkYamlConfig := auth.ProjectToken{}
	if chalkYamlConfigOrNil != nil {
		chalkYamlConfig = *chalkYamlConfigOrNil
	}

	envIdConfig := auth.GetFirstNonEmptyConfig(
		auth.GetChalkClientArgConfig(environmentId),
		auth.GetEnvVarConfig(internal.EnvironmentEnvVarKey),
		auth.GetChalkYamlConfig(chalkYamlConfig.ActiveEnvironment),
	)

	if logger == nil {
		logger = DefaultLeveledLogger
	}
	manager := &configManager{
		apiServer: auth.GetFirstNonEmptyConfig(
			auth.GetChalkClientArgConfig(apiServer),
			auth.GetEnvVarConfig(internal.ApiServerEnvVarKey),
			auth.GetChalkYamlConfig(chalkYamlConfig.ApiServer),
		),
		clientId: auth.GetFirstNonEmptyConfig(
			auth.GetChalkClientArgConfig(clientId),
			auth.GetEnvVarConfig(internal.ClientIdEnvVarKey),
			auth.GetChalkYamlConfig(chalkYamlConfig.ClientId),
		),
		clientSecret: auth.GetFirstNonEmptyConfig(
			auth.GetChalkClientArgConfig(clientSecret),
			auth.GetEnvVarConfig(internal.ClientSecretEnvVarKey),
			auth.GetChalkYamlConfig(chalkYamlConfig.ClientSecret),
		),
		environmentId:      envIdConfig,
		initialEnvironment: envIdConfig,
		logger:             logger,
	}

	if chalkYamlErr != nil && manager.clientId.Value == "" && manager.clientSecret.Value == "" {
		return nil, errors.Wrap(
			chalkYamlErr,
			"could not read chalk.yml and no client ID and client secret were provided",
		)
	}

	return manager, nil
}

func (m *configManager) getQueryServer(queryServerOverride *string) string {
	if queryServerOverride != nil {
		return *queryServerOverride
	}

	endpoint, ok := m.engines[m.environmentId.Value]
	if !ok {
		if m.engines != nil {
			m.logger.Errorf(
				"query endpoint falling back to api server - no engine found for environment '%s' - engine map keys: '%s'",
				m.environmentId.Value,
				colls.Keys(m.engines),
			)
		} else {
			m.logger.Errorf(
				"query endpoint falling back to api server - no engine found for environment '%s'",
				m.environmentId.Value,
			)
		}
		endpoint = m.apiServer.Value
	}
	return endpoint
}

func (m *configManager) refresh(ctx context.Context, force bool) error {
	if !force && m.jwt != nil && m.jwt.IsValid() {
		return nil
	}

	config, err := m.getToken(ctx, m.clientId.Value, m.clientSecret.Value)
	if err != nil {
		return errors.Wrap(err, "refreshing token")
	}

	if m.initialEnvironment.Value == "" {
		m.environmentId = auth.SourcedConfig{
			Value:  config.PrimaryEnvironment,
			Source: "Primary Environment from credentials exchange response",
		}
	} else {
		m.environmentId = m.initialEnvironment
	}
	m.jwt = &auth.JWT{
		Token:      config.AccessToken,
		ValidUntil: config.ValidUntil,
	}
	m.engines = config.Engines
	return nil
}
