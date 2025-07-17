package tests

import (
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

// TestQueryServerOverride tests that when we specify
// a query server override that query server is actually used.
func TestQueryServerOverride(t *testing.T) {
	t.Parallel()
	queryServer := "https://my-bogus-server.ai"
	client, httpClient, err := newClientWithInterceptor(
		t.Context(),
		interceptorClientOverrides{
			QueryServer: queryServer,
		})
	assert.NoError(t, err)
	req := chalk.OnlineQueryParams{}.
		WithInput("bogus.feature", 1).
		WithOutputs("bogus.output")
	_, _ = client.OnlineQuery(t.Context(), req, nil)
	parsed, err := url.Parse(queryServer)
	assert.Nil(t, err)
	assert.Equal(t, httpClient.Intercepted.URL.Host, parsed.Host)
}
