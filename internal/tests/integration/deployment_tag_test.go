package integration

import (
	"context"
	"github.com/chalk-ai/chalk-go"
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
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient:    httpClient,
		DeploymentTag: deploymentTag,
	})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	userIds := []int{1}
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}
	req := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, userIds[0]).
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OnlineQuery(context.Background(), req, nil)

	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Deployment-Tag"), deploymentTag)

	bulkReq := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OnlineQueryBulk(context.Background(), bulkReq)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Deployment-Tag"), deploymentTag)
}
