package chalk

import (
	"github.com/chalk-ai/chalk-go/internal"
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
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
	queryServer        *string

	jwt     *auth2.JWT
	engines map[string]string

	getToken func() (*getTokenResult, error)
}

func getConfigManager(cfg ClientConfig) (*configManager, error) {
	chalkYamlConfig, chalkYamlErr := auth2.GetProjectAuthConfig()

	apiServerOverride := auth2.GetChalkClientArgConfig(cfg.ApiServer)
	clientIdOverride := auth2.GetChalkClientArgConfig(cfg.ClientId)
	clientSecretOverride := auth2.GetChalkClientArgConfig(cfg.ClientSecret)
	environmentIdOverride := auth2.GetChalkClientArgConfig(cfg.EnvironmentId)

	apiServerEnvVarConfig := auth2.GetEnvVarConfig(internal.ApiServerEnvVarKey)
	clientIdEnvVarConfig := auth2.GetEnvVarConfig(internal.ClientIdEnvVarKey)
	clientSecretEnvVarConfig := auth2.GetEnvVarConfig(internal.ClientSecretEnvVarKey)
	environmentIdEnvVarConfig := auth2.GetEnvVarConfig(internal.EnvironmentEnvVarKey)

	apiServerFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ApiServer)
	clientIdFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ClientId)
	clientSecretFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ClientSecret)
	environmentIdFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ActiveEnvironment)

	apiServer := auth2.GetFirstNonEmptyConfig(apiServerOverride, apiServerEnvVarConfig, apiServerFileConfig)
	clientId := auth2.GetFirstNonEmptyConfig(clientIdOverride, clientIdEnvVarConfig, clientIdFileConfig)
	clientSecret := auth2.GetFirstNonEmptyConfig(clientSecretOverride, clientSecretEnvVarConfig, clientSecretFileConfig)
	environmentId := auth2.GetFirstNonEmptyConfig(environmentIdOverride, environmentIdEnvVarConfig, environmentIdFileConfig)

	if chalkYamlErr != nil && clientId.Value == "" && clientSecret.Value == "" {
		return nil, chalkYamlErr
	}

	return &configManager{
		apiServer:          apiServer,
		clientId:           clientId,
		clientSecret:       clientSecret,
		environmentId:      environmentId,
		initialEnvironment: environmentId,
		queryServer:        internal.Ptr(cfg.QueryServer),
	}, nil
}

func (r *configManager) refresh(force bool) error {
	if !force && r.jwt != nil && r.jwt.IsValid() {
		return nil
	}

	config, getTokenErr := r.getToken()
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
