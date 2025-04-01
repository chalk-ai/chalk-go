package tests

import (
	"context"
	"errors"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

type paramsTestFeatures struct {
	MyFeature *int
}

var paramsTestRoot = struct {
	MyFeaturesClass *paramsTestFeatures
}{}

func init() {
	err := chalk.InitFeatures(&paramsTestRoot)
	if err != nil {
		panic(err)
	}
}

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

// TestHeadersSetInOnlineQueryBulk tests that when we specify certain params,
// they propagate to the request.
func TestHeadersSetInOnlineQueryBulk(t *testing.T) {
	client, interceptor, err := newClientWithInterceptor()
	assert.NoError(t, err)
	resourceGroup := "bogus-resource-group"
	req := chalk.OnlineQueryParams{
		ResourceGroup: resourceGroup,
	}.
		WithInput(paramsTestRoot.MyFeaturesClass.MyFeature, []int{1}).
		WithOutputs(paramsTestRoot.MyFeaturesClass.MyFeature)
	_, _ = client.OnlineQueryBulk(context.Background(), req)
	assert.Equal(t, resourceGroup, interceptor.Intercepted.Header.Get(chalk.HeaderKeyResourceGroup))
	assert.Equal(t, "true", interceptor.Intercepted.Header.Get("X-Chalk-Features-Versioned"))
}

// TestHeadersSetOnlineQuery tests that when we specify
// features using codegen-ed structs, we set the "versioned" header.
func TestHeadersSetOnlineQuery(t *testing.T) {
	client, httpClient, err := newClientWithInterceptor()
	assert.NoError(t, err)

	resourceGroup := "bogus-resource-group"
	req := chalk.OnlineQueryParams{
		ResourceGroup: resourceGroup,
	}.
		WithInput(paramsTestRoot.MyFeaturesClass.MyFeature, 1).
		WithOutputs(paramsTestRoot.MyFeaturesClass.MyFeature)
	_, _ = client.OnlineQuery(context.Background(), req, nil)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Features-Versioned"), "true")
	assert.Equal(t, resourceGroup, httpClient.Intercepted.Header.Get(chalk.HeaderKeyResourceGroup))
}

func TestHeadersSetOfflineQuery(t *testing.T) {
	client, interceptor, err := newClientWithInterceptor()
	assert.NoError(t, err)
	req := chalk.OfflineQueryParams{}.
		WithInput(paramsTestRoot.MyFeaturesClass.MyFeature, []any{1}).
		WithOutputs(paramsTestRoot.MyFeaturesClass.MyFeature)
	_, _ = client.OfflineQuery(context.Background(), req)
	assert.Equal(t, "true", interceptor.Intercepted.Header.Get("X-Chalk-Features-Versioned"))
}

// TestQueryServerOverride tests that when we specify
// a query server override that query server is actually used.
func TestQueryServerOverride(t *testing.T) {
	queryServer := "https://my-bogus-server.ai"
	client, httpClient, err := newClientWithInterceptor(interceptorClientOverrides{
		QueryServer: queryServer,
	})
	assert.NoError(t, err)
	req := chalk.OnlineQueryParams{}.
		WithInput("bogus.feature", 1).
		WithOutputs("bogus.output")
	_, _ = client.OnlineQuery(context.Background(), req, nil)
	parsed, err := url.Parse(queryServer)
	assert.Nil(t, err)
	assert.Equal(t, httpClient.Intercepted.URL.Host, parsed.Host)
}

// TestOnlineQueryAndQueryBulkDeploymentTagInRequest tests that when we
// specify a deployment Tag in the client build, the request
// includes the deployment Tag header.
func TestOnlineQueryAndQueryBulkDeploymentTagInRequest(t *testing.T) {
	deploymentTag := "test-deployment-tag"
	client, httpClient, err := newClientWithInterceptor(interceptorClientOverrides{
		DeploymentTag: deploymentTag,
	})
	assert.NoError(t, err)
	ids := []int{1}
	req := chalk.OnlineQueryParams{}.
		WithInput("bogus.feature", ids[0]).
		WithOutputs("bogus.output")
	_, _ = client.OnlineQuery(context.Background(), req, nil)

	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Deployment-Tag"), deploymentTag)

	bulkReq := chalk.OnlineQueryParams{}.
		WithInput("bogus.feature", ids).
		WithOutputs("bogus.output")
	_, _ = client.OnlineQueryBulk(context.Background(), bulkReq)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Deployment-Tag"), deploymentTag)
}
