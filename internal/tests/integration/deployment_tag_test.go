package integration

import (
	"context"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

// TestOnlineQueryAndQueryBulkDeploymentTagInRequest tests that when we
// specify a deployment Tag in the client build, the request
// includes the deployment Tag header.
func TestOnlineQueryAndQueryBulkDeploymentTagInRequest(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewInterceptorHTTPClient()
	deploymentTag := "test-deployment-tag"
	client, err := chalk.NewClient(
		context.Background(),
		&chalk.ClientConfig{
			HTTPClient:    httpClient,
			DeploymentTag: deploymentTag,
		},
	)
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	ids := []int{1}
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}
	req := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.AllTypes.Id, ids[0]).
		WithOutputs(testFeatures.AllTypes.StrFeat)
	_, _ = client.OnlineQuery(context.Background(), req, nil)

	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Deployment-Tag"), deploymentTag)

	bulkReq := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.AllTypes.Id, ids).
		WithOutputs(testFeatures.AllTypes.StrFeat)
	_, _ = client.OnlineQueryBulk(context.Background(), bulkReq)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Deployment-Tag"), deploymentTag)
}
