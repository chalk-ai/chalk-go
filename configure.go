package chalk

import (
	"github.com/chalk-ai/chalk-go/internal"
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
	"net/http"
)

func getConfiguredClient(configOverride *auth2.ProjectAuthConfigOverride) (*ChalkClientImpl, error) {
	if configOverride == nil {
		configOverride = &auth2.ProjectAuthConfigOverride{}
	}
	// TODO: Check error here
	projectAuthConfigFromFile, _, _ := auth2.LoadAuthConfig().GetProjectAuthConfigForWD()

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

	client := &ChalkClientImpl{
		httpClient:    &http.Client{},
		ApiServer:     getFirstNonEmptyConfig(apiServerOverride, apiServerEnvVarConfig, apiServerFileConfig),
		ClientId:      getFirstNonEmptyConfig(clientIdOverride, clientIdEnvVarConfig, clientIdFileConfig),
		clientSecret:  getFirstNonEmptyConfig(clientSecretOverride, clientSecretEnvVarConfig, clientSecretFileConfig),
		EnvironmentId: getFirstNonEmptyConfig(environmentIdOverride, environmentIdEnvVarConfig, environmentIdFileConfig),
		logger:        &DefaultLeveledLogger,
	}

	err := client.refreshJwt(false)
	if err != nil {
		return nil, err
	}
	return client, nil
}
