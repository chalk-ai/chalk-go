package integration

import (
	"context"
	"github.com/chalk-ai/chalk-go/v2"
	assert "github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

// TestHeadersSetOnlineQueryBulk tests that when we specify
// features using codegen-ed structs, we set the "versioned" header.
func TestHeadersSetOnlineQueryBulk(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewInterceptorHTTPClient()
	client, err := chalk.NewClient(
		context.Background(),
		&chalk.ClientConfig{
			HTTPClient: httpClient,
		},
	)
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	userIds := []int{1}
	resourceGroup := "bogus-resource-group"
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}
	req := chalk.OnlineQueryParams{
		ResourceGroup: resourceGroup,
	}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OnlineQueryBulk(context.Background(), req)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Features-Versioned"), "true")
	assert.Equal(t, resourceGroup, httpClient.Intercepted.Header.Get(chalk.HeaderKeyResourceGroup))
}

// TestHeadersSetOnlineQuery tests that when we specify
// features using codegen-ed structs, we set the "versioned" header.
func TestHeadersSetOnlineQuery(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewInterceptorHTTPClient()
	client, err := chalk.NewClient(
		context.Background(),
		&chalk.ClientConfig{
			HTTPClient: httpClient,
		})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}

	resourceGroup := "bogus-resource-group"

	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}
	req := chalk.OnlineQueryParams{
		ResourceGroup: resourceGroup,
	}.
		WithInput(testFeatures.User.Id, 1).
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OnlineQuery(context.Background(), req, nil)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Features-Versioned"), "true")
	assert.Equal(t, resourceGroup, httpClient.Intercepted.Header.Get(chalk.HeaderKeyResourceGroup))
}

// TestVersionHeaderSetOfflineQuery tests that when we specify
// features using codegen-ed structs, we set the "versioned" header.
func TestVersionHeaderSetOfflineQuery(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewInterceptorHTTPClient()
	client, err := chalk.NewClient(
		context.Background(),
		&chalk.ClientConfig{
			HTTPClient: httpClient,
		},
	)
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}
	userIds := []any{1}
	req := chalk.OfflineQueryParams{}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OfflineQuery(context.Background(), req)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Features-Versioned"), "true")
}

// TestQueryServerOverride tests that when we specify
// a query server override that query server is actually used.
func TestQueryServerOverride(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewInterceptorHTTPClient()
	queryServer := "https://my-bogus-server.ai"
	client, err := chalk.NewClient(
		context.Background(),
		&chalk.ClientConfig{
			HTTPClient:  httpClient,
			QueryServer: queryServer,
		},
	)
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}
	req := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, 1).
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OnlineQuery(context.Background(), req, nil)
	parsed, err := url.Parse(queryServer)
	assert.Nil(t, err)
	assert.Equal(t, httpClient.Intercepted.URL.Host, parsed.Host)
}
