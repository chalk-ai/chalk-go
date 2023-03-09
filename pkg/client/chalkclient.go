package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
)

func ChalkClient(configOverride auth.ProjectAuthConfigOverride) (*Client, error) {
	client := getConfiguredClient(configOverride)

	jwt, getJWTErr := client.getJwt()
	if getJWTErr != nil {
		return client, getJWTErr
	}
	client.jwt = jwt

	return client, nil
}

func getConfiguredClient(configOverride auth.ProjectAuthConfigOverride) *Client {
	var client *Client

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

	client = &Client{
		httpClient:    &http.Client{},
		ApiServer:     getFirstNonEmptyConfig(apiServerOverride, apiServerEnvVarConfig, apiServerFileConfig),
		ClientId:      getFirstNonEmptyConfig(clientIdOverride, clientIdEnvVarConfig, clientIdFileConfig),
		ClientSecret:  getFirstNonEmptyConfig(clientSecretOverride, clientSecretEnvVarConfig, clientSecretFileConfig),
		EnvironmentId: getFirstNonEmptyConfig(environmentIdOverride, environmentIdEnvVarConfig, environmentIdFileConfig),
	}

	return client
}
