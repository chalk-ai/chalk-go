package tests

import (
	"context"
	"errors"
	"github.com/chalk-ai/chalk-go"
)

type interceptorClientOverrides struct {
	QueryServer   string
	DeploymentTag string
}

func newClientWithInterceptor(overrides ...interceptorClientOverrides) (chalk.Client, *InterceptorHTTPClient, error) {
	var queryServer = ""
	var deploymentTag = ""
	if len(overrides) > 1 {
		return nil, nil, errors.New("too many overrides")
	} else if len(overrides) == 1 {
		queryServer = overrides[0].QueryServer
		deploymentTag = overrides[0].DeploymentTag
	}

	httpClient := NewInterceptorHTTPClient()
	client, err := chalk.NewClient(
		context.Background(),
		&chalk.ClientConfig{
			HTTPClient:    httpClient,
			ApiServer:     "https://bogus.com",
			ClientId:      "bogus-client-id",
			ClientSecret:  "ts-bogus-client-secret",
			QueryServer:   queryServer,
			DeploymentTag: deploymentTag,
		},
	)
	if err != nil {
		return nil, nil, err
	}
	return client, httpClient, nil
}
