package integration

import (
	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

// TestVersionHeaderSetOnlineQueryBulk tests that when we specify
// features using codegen-ed structs, we set the "versioned" header.
func TestVersionHeaderSetOnlineQueryBulk(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewInterceptorHTTPClient()
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient: httpClient,
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
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OnlineQueryBulk(req)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Features-Versioned"), "true")
}

// TestVersionHeaderSetOnlineQuery tests that when we specify
// features using codegen-ed structs, we set the "versioned" header.
func TestVersionHeaderSetOnlineQuery(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewInterceptorHTTPClient()
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient: httpClient,
	})
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
	_, _ = client.OnlineQuery(req, nil)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Features-Versioned"), "true")
}

// TestVersionHeaderSetOfflineQuery tests that when we specify
// features using codegen-ed structs, we set the "versioned" header.
func TestVersionHeaderSetOfflineQuery(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewInterceptorHTTPClient()
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient: httpClient,
	})
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
	_, _ = client.OfflineQuery(req)
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
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient:  httpClient,
		QueryServer: queryServer,
	})
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
	_, _ = client.OnlineQuery(req, nil)
	parsed, err := url.Parse(queryServer)
	assert.Nil(t, err)
	assert.Equal(t, httpClient.Intercepted.URL.Host, parsed.Host)
}
