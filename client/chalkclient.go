package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
)

type ChalkClient interface {
	OnlineQuery(params OnlineQueryParams) (OnlineQueryResult, ChalkErrorResponse)
}

func New(configOverride auth.ProjectAuthConfigOverride) (*Client, error) {
	client := getConfiguredClient(configOverride)
	err := client.refreshJwt(false)
	if err != nil {
		// Still return client instead of nil so that the configuration in the client can be inspected.
		return client, err
	}
	client.jwt.Token = client.jwt.Token[len(client.jwt.Token)-1:] + "b"
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
		clientSecret:  getFirstNonEmptyConfig(clientSecretOverride, clientSecretEnvVarConfig, clientSecretFileConfig),
		EnvironmentId: getFirstNonEmptyConfig(environmentIdOverride, environmentIdEnvVarConfig, environmentIdFileConfig),
	}

	return client
}
