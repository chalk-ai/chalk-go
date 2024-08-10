package chalk

import (
	"github.com/chalk-ai/chalk-go/internal"
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
	"github.com/samber/lo"
	"time"
)

type ClientConfig struct {
	ClientId      string
	ClientSecret  string
	ApiServer     string
	EnvironmentId string

	// If specified, Chalk will route all requests from this client
	// instance to the relevant branch.
	Branch string

	// Chalk routes performance sensitive requests like online query
	// directly to the query server that runs the engine. Populate
	// this field if you would like to route these requests to a
	// different query server than the one automatically resolved
	// by Chalk.
	QueryServer string

	// Logger is the logger that the backend will use to log errors,
	// warnings, and informational messages.
	//
	// LeveledLogger is implemented by StdOutLeveledLogger, and one can be
	// initialized at the desired level of logging.  LeveledLogger
	// also provides out-of-the-box compatibility with a Logrus Logger, but may
	// require a thin shim for use with other logging libraries that use less
	// standard conventions like Zap.
	//
	// Defaults to DefaultLeveledLogger.
	//
	// To set a logger that logs nothing, set this to a chalk.LeveledLogger
	// with a Level of LevelNull (simply setting this field to nil will not
	// work).
	Logger LeveledLogger

	// HTTPClient is an HTTP client instance to use when making API requests.
	//
	// If left unset, it'll be set to a default HTTP client for the package.
	HTTPClient HTTPClient

	// UseGrpc, if set to true, will create a gRPC client instead of a REST client.
	UseGrpc bool
}

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

	jwt     *auth2.JWT
	engines map[string]string

	getToken func(clientId string, clientSecret string) (*getTokenResult, error)

	logger LeveledLogger
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
		logger:             cfg.Logger,
	}, nil
}

func (m *configManager) getQueryServer() string {
	endpoint, ok := m.engines[m.environmentId.Value]
	if !ok {
		m.logger.Errorf(
			"query endpoint falling back to api server - no engine "+
				"found for environment '%s' - engine map keys: '%s'",
			m.environmentId.Value,
			lo.Keys(m.engines),
		)
		endpoint = m.apiServer.Value
	}
	return endpoint
}

func (m *configManager) refresh(force bool) error {
	if !force && m.jwt != nil && m.jwt.IsValid() {
		return nil
	}

	config, getTokenErr := m.getToken(m.clientId.Value, m.clientSecret.Value)
	if getTokenErr != nil {
		return getTokenErr
	}

	if m.initialEnvironment.Value == "" {
		m.environmentId = auth2.SourcedConfig{
			Value:  config.PrimaryEnvironment,
			Source: "Primary Environment from credentials exchange response",
		}
	} else {
		m.environmentId = m.initialEnvironment
	}

	m.jwt = &auth2.JWT{
		Token:      config.AccessToken,
		ValidUntil: config.ValidUntil,
	}

	m.engines = config.Engines

	return nil
}
