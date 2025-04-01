package tests

import (
	"context"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
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
