package chalk

import (
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
)

func getConfiguredClient(
	configOverride *ClientConfig,
) *chalkClientImpl {
	if configOverride == nil {
		configOverride = &ClientConfig{}
	}
	projectAuthConfigFromFile, _, _ := auth.LoadAuthConfig().GetProjectAuthConfigForWD()

	apiServerOverride := getChalkClientArgConfig(configOverride.ApiServer)
	clientIdOverride := getChalkClientArgConfig(configOverride.ClientId)
	clientSecretOverride := getChalkClientArgConfig(configOverride.ClientSecret)
	environmentIdOverride := getChalkClientArgConfig(configOverride.EnvironmentId)

	apiServerEnvVarConfig := getEnvVarConfig(internal.ApiServerEnvVarKey)
	clientIdEnvVarConfig := getEnvVarConfig(internal.ClientIdEnvVarKey)
	clientSecretEnvVarConfig := getEnvVarConfig(internal.ClientSecretEnvVarKey)
	environmentIdEnvVarConfig := getEnvVarConfig(internal.EnvironmentEnvVarKey)

	apiServerFileConfig := getChalkYamlConfig(projectAuthConfigFromFile.ApiServer)
	clientIdFileConfig := getChalkYamlConfig(projectAuthConfigFromFile.ClientId)
	clientSecretFileConfig := getChalkYamlConfig(projectAuthConfigFromFile.ClientSecret)
	environmentIdFileConfig := getChalkYamlConfig(projectAuthConfigFromFile.ActiveEnvironment)

	client := &chalkClientImpl{
		httpClient:    configOverride.HTTPClient,
		logger:        configOverride.Logger,
		ApiServer:     getFirstNonEmptyConfig(apiServerOverride, apiServerEnvVarConfig, apiServerFileConfig),
		ClientId:      getFirstNonEmptyConfig(clientIdOverride, clientIdEnvVarConfig, clientIdFileConfig),
		clientSecret:  getFirstNonEmptyConfig(clientSecretOverride, clientSecretEnvVarConfig, clientSecretFileConfig),
		EnvironmentId: getFirstNonEmptyConfig(environmentIdOverride, environmentIdEnvVarConfig, environmentIdFileConfig),
	}

	if client.logger == nil {
		client.logger = &DefaultLeveledLogger
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{}
	}

	return client
}
