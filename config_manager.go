package chalk

import (
	"github.com/chalk-ai/chalk-go/internal"
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
	"github.com/chalk-ai/chalk-go/internal/colls"
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
	apiServer          auth2.SourcedConfig
	clientId           auth2.SourcedConfig
	clientSecret       auth2.SourcedConfig
	environmentId      auth2.SourcedConfig
	initialEnvironment auth2.SourcedConfig

	jwt      *auth2.JWT
	engines  map[string]string
	getToken func(clientId string, clientSecret string) (*getTokenResult, error)

	logger LeveledLogger
}

func newConfigManager(
	apiServer string,
	clientId string,
	clientSecret string,
	environmentId string,
	logger LeveledLogger,
) (*configManager, error) {
	chalkYamlConfig, chalkYamlErr := auth2.GetProjectAuthConfig()
	if logger == nil {
		logger = DefaultLeveledLogger
	}

	envIdConfig := auth2.GetFirstNonEmptyConfig(
		auth2.GetChalkClientArgConfig(environmentId),
		auth2.GetEnvVarConfig(internal.EnvironmentEnvVarKey),
		auth2.GetChalkYamlConfig(chalkYamlConfig.ActiveEnvironment),
	)

	manager := &configManager{
		apiServer: auth2.GetFirstNonEmptyConfig(
			auth2.GetChalkClientArgConfig(apiServer),
			auth2.GetEnvVarConfig(internal.ApiServerEnvVarKey),
			auth2.GetChalkYamlConfig(chalkYamlConfig.ApiServer),
		),
		clientId: auth2.GetFirstNonEmptyConfig(
			auth2.GetChalkClientArgConfig(clientId),
			auth2.GetEnvVarConfig(internal.ClientIdEnvVarKey),
			auth2.GetChalkYamlConfig(chalkYamlConfig.ClientId),
		),
		clientSecret: auth2.GetFirstNonEmptyConfig(
			auth2.GetChalkClientArgConfig(clientSecret),
			auth2.GetEnvVarConfig(internal.ClientSecretEnvVarKey),
			auth2.GetChalkYamlConfig(chalkYamlConfig.ClientSecret),
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
				"query endpoint falling back to api server - no engine "+
					"found for environment '%s' - engine map keys: '%s'",
				m.environmentId.Value,
				colls.Keys(m.engines),
			)
		} else {
			m.logger.Errorf(
				"query endpoint falling back to api server - no engine "+
					"found for environment '%s'",
				m.environmentId.Value,
			)
		}
		endpoint = m.apiServer.Value
	}
	return endpoint
}

func (r *configManager) refresh(force bool) error {
	if !force && r.jwt != nil && r.jwt.IsValid() {
		return nil
	}

	config, getTokenErr := r.getToken(r.clientId.Value, r.clientSecret.Value)
	if getTokenErr != nil {
		return getTokenErr
	}

	if r.initialEnvironment.Value == "" {
		r.environmentId = auth2.SourcedConfig{
			Value:  config.PrimaryEnvironment,
			Source: "Primary Environment from credentials exchange response",
		}
	} else {
		r.environmentId = r.initialEnvironment
	}

	r.jwt = &auth2.JWT{
		Token:      config.AccessToken,
		ValidUntil: config.ValidUntil,
	}

	r.engines = config.Engines

	return nil
}
