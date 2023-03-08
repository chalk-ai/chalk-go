package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
)

func New(configOverride auth.ProjectAuthConfigOverride) (*Client, error) {
	client := getConfiguredClient(configOverride)

	jwt, getJWTErr := client.getJwt()
	if getJWTErr != nil {
		return nil, getJWTErr
	}
	client.jwt = jwt

	return client, nil
}

func getConfiguredClient(configOverride auth.ProjectAuthConfigOverride) *Client {
	fileConfigOveridden := false
	var client *Client
	if configOverride.ClientId != "" && configOverride.ClientSecret != "" && configOverride.ApiServer != "" && configOverride.EnvironmentId != "" {
		// Skip loading from file
		fileConfigOveridden = true
		client = &Client{
			BaseUrl:       configOverride.ApiServer,
			httpClient:    &http.Client{},
			EnvironmentId: configOverride.EnvironmentId,
			ClientId:      configOverride.ClientId,
			ClientSecret:  configOverride.ClientSecret,
		}
	} else {
		projectAuthConfigFromFile, _, _ := auth.LoadAuthConfig().GetProjectAuthConfigByDirectory()
		fileConfigOveridden = configOverride.ClientSecret != "" && configOverride.ClientId != ""
		client = &Client{
			BaseUrl:       projectAuthConfigFromFile.ApiServer,
			httpClient:    &http.Client{},
			EnvironmentId: projectAuthConfigFromFile.ActiveEnvironment,
			ClientId:      projectAuthConfigFromFile.ClientId,
			ClientSecret:  projectAuthConfigFromFile.ClientSecret,
		}

		if configOverride.ApiServer != "" {
			client.BaseUrl = configOverride.ApiServer
		}
		if configOverride.EnvironmentId != "" {
			client.EnvironmentId = configOverride.EnvironmentId
		}
		if fileConfigOveridden {
			client.ClientId = configOverride.ClientId
			client.ClientSecret = configOverride.ClientSecret
		}
	}

	return client
}
