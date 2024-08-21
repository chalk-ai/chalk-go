package chalk

import (
	"github.com/chalk-ai/chalk-go/internal"
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
	"github.com/chalk-ai/chalk-go/internal/colls"
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

	apiServerOverride := auth2.GetChalkClientArgConfig(apiServer)
	clientIdOverride := auth2.GetChalkClientArgConfig(clientId)
	clientSecretOverride := auth2.GetChalkClientArgConfig(clientSecret)
	environmentIdOverride := auth2.GetChalkClientArgConfig(environmentId)

	apiServerEnvVarConfig := auth2.GetEnvVarConfig(internal.ApiServerEnvVarKey)
	clientIdEnvVarConfig := auth2.GetEnvVarConfig(internal.ClientIdEnvVarKey)
	clientSecretEnvVarConfig := auth2.GetEnvVarConfig(internal.ClientSecretEnvVarKey)
	environmentIdEnvVarConfig := auth2.GetEnvVarConfig(internal.EnvironmentEnvVarKey)

	apiServerFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ApiServer)
	clientIdFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ClientId)
	clientSecretFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ClientSecret)
	environmentIdFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ActiveEnvironment)

	apiServerConfig := auth2.GetFirstNonEmptyConfig(apiServerOverride, apiServerEnvVarConfig, apiServerFileConfig)
	clientIdConfig := auth2.GetFirstNonEmptyConfig(clientIdOverride, clientIdEnvVarConfig, clientIdFileConfig)
	clientSecretConfig := auth2.GetFirstNonEmptyConfig(clientSecretOverride, clientSecretEnvVarConfig, clientSecretFileConfig)
	environmentIdConfig := auth2.GetFirstNonEmptyConfig(environmentIdOverride, environmentIdEnvVarConfig, environmentIdFileConfig)

	if chalkYamlErr != nil && clientIdConfig.Value == "" && clientSecretConfig.Value == "" {
		return nil, chalkYamlErr
	}

	return &configManager{
		apiServer:          apiServerConfig,
		clientId:           clientIdConfig,
		clientSecret:       clientSecretConfig,
		environmentId:      environmentIdConfig,
		initialEnvironment: environmentIdConfig,
		logger:             logger,
	}, nil
}

func (m *configManager) getQueryServer(queryServerOverride *string) string {
	if queryServerOverride != nil {
		return *queryServerOverride
	}

	endpoint, ok := m.engines[m.environmentId.Value]
	if !ok {
		m.logger.Errorf(
			"query endpoint falling back to api server - no engine "+
				"found for environment '%s' - engine map keys: '%s'",
			m.environmentId.Value,
			colls.Keys(m.engines),
		)
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
