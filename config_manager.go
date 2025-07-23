package chalk

import (
	"context"
	"fmt"
	"github.com/chalk-ai/chalk-go/config"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/auth"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/cockroachdb/errors"
	"strings"
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

// ConfigManager manages configuration for the Chalk client
type ConfigManager struct {
	apiServer          *config.SourcedConfig[string]
	clientId           *config.SourcedConfig[string]
	clientSecret       *config.SourcedConfig[string]
	environmentId      *config.SourcedConfig[string]
	initialEnvironment *config.SourcedConfig[string]

	jwt      *auth.JWT
	engines  map[string]string
	getToken func(ctx context.Context, clientId string, clientSecret string) (*getTokenResult, error)

	logger LeveledLogger
}

// NewConfigManager creates a new configuration manager with the specified parameters
func NewConfigManager(
	apiServer *config.SourcedConfig[string],
	clientId *config.SourcedConfig[string],
	clientSecret *config.SourcedConfig[string],
	environmentId *config.SourcedConfig[string],
	configDir *string,
	logger LeveledLogger,
) (*ConfigManager, error) {
	chalkYamlConfigOrNil, chalkYamlErr := auth.GetProjectAuthConfig(configDir)
	chalkYamlConfig := auth.ProjectToken{}
	if chalkYamlConfigOrNil != nil {
		chalkYamlConfig = *chalkYamlConfigOrNil
	}

	envIdConfig := config.GetFirstNonEmpty(
		environmentId,
		auth.GetEnvVarConfig(internal.EnvironmentEnvVarKey),
		auth.GetChalkYamlConfig(chalkYamlConfig.ActiveEnvironment, configDir),
	)

	if logger == nil {
		logger = DefaultLeveledLogger
	}
	manager := &ConfigManager{
		apiServer: config.GetFirstNonEmpty(
			apiServer,
			auth.GetEnvVarConfig(internal.ApiServerEnvVarKey),
			auth.GetChalkYamlConfig(chalkYamlConfig.ApiServer, configDir),
		),
		clientId: config.GetFirstNonEmpty(
			clientId,
			auth.GetEnvVarConfig(internal.ClientIdEnvVarKey),
			auth.GetChalkYamlConfig(chalkYamlConfig.ClientId, configDir),
		),
		clientSecret: config.GetFirstNonEmpty(
			clientSecret,
			auth.GetEnvVarConfig(internal.ClientSecretEnvVarKey),
			auth.GetChalkYamlConfig(chalkYamlConfig.ClientSecret, configDir),
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

func (m *ConfigManager) getQueryServer(queryServerOverride *string) string {
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

	if strings.HasPrefix(endpoint, "https://") || strings.HasPrefix(endpoint, "http://") {
		return endpoint
	}
	return fmt.Sprintf("https://%s", endpoint)
}

func (m *ConfigManager) refresh(ctx context.Context, force bool) error {
	if !force && m.jwt != nil && m.jwt.IsValid() {
		return nil
	}

	c, err := m.getToken(ctx, m.clientId.Value, m.clientSecret.Value)
	if err != nil {
		return errors.Wrap(err, "refreshing token")
	}

	if m.initialEnvironment.Value == "" {
		m.environmentId = &config.SourcedConfig[string]{
			Value:  c.PrimaryEnvironment,
			Source: "Primary Environment from credentials exchange response",
		}
	} else {
		m.environmentId = m.initialEnvironment
	}
	m.jwt = &auth.JWT{
		Token:      c.AccessToken,
		ValidUntil: c.ValidUntil,
	}
	m.engines = c.Engines
	return nil
}
