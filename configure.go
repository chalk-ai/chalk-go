package chalk

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
)

func getConfiguredClient(configOverride *auth.ProjectAuthConfigOverride) *chalkClientImpl {
	if configOverride == nil {
		configOverride = &auth.ProjectAuthConfigOverride{}
	}
	projectAuthConfigFromFile, _, _ := auth.LoadAuthConfig().GetProjectAuthConfigForWD()

	apiServerOverride := getChalkClientArgConfig(configOverride.ApiServer)
	clientIdOverride := getChalkClientArgConfig(configOverride.ClientId)
	clientSecretOverride := getChalkClientArgConfig(configOverride.ClientSecret)
	environmentIdOverride := getChalkClientArgConfig(configOverride.EnvironmentId)

	apiServerEnvVarConfig := getEnvVarConfig(apiServerEnvVarKey)
	clientIdEnvVarConfig := getEnvVarConfig(clientIdEnvVarKey)
	clientSecretEnvVarConfig := getEnvVarConfig(clientSecretEnvVarKey)
	environmentIdEnvVarConfig := getEnvVarConfig(environmentEnvVarKey)

	apiServerFileConfig := getChalkYamlConfig(projectAuthConfigFromFile.ApiServer)
	clientIdFileConfig := getChalkYamlConfig(projectAuthConfigFromFile.ClientId)
	clientSecretFileConfig := getChalkYamlConfig(projectAuthConfigFromFile.ClientSecret)
	environmentIdFileConfig := getChalkYamlConfig(projectAuthConfigFromFile.ActiveEnvironment)

	client := &chalkClientImpl{
		httpClient:    &http.Client{},
		ApiServer:     getFirstNonEmptyConfig(apiServerOverride, apiServerEnvVarConfig, apiServerFileConfig),
		ClientId:      getFirstNonEmptyConfig(clientIdOverride, clientIdEnvVarConfig, clientIdFileConfig),
		clientSecret:  getFirstNonEmptyConfig(clientSecretOverride, clientSecretEnvVarConfig, clientSecretFileConfig),
		EnvironmentId: getFirstNonEmptyConfig(environmentIdOverride, environmentIdEnvVarConfig, environmentIdFileConfig),
	}

	return client
}
