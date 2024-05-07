package integration

import (
	"bytes"
	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

type Intercepted struct {
	Header http.Header
}

type InterceptorHTTPClient struct {
	Intercepted Intercepted
}

func NewCovertHTTPClient() *InterceptorHTTPClient {
	return &InterceptorHTTPClient{}
}

func (c *InterceptorHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.Intercepted = Intercepted{
		Header: req.Header,
	}
	body := io.NopCloser(bytes.NewBufferString(`{"data": {"something": "exciting"}}`))
	return &http.Response{StatusCode: 200, Body: body}, nil
}

func (c *InterceptorHTTPClient) Get(url string) (*http.Response, error) {
	actualClient := &http.Client{}
	return actualClient.Get(url)
}

// TestOnlineQueryAndQueryBulkBranchInRequest tests that when we
// specify a branch ID in online query params, the request
// includes the branch ID header.
func TestOnlineQueryAndQueryBulkBranchInRequest(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	httpClient := NewCovertHTTPClient()
	branchId := "test-branch-id"
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
		WithInput(testFeatures.User.Id, userIds[0]).
		WithOutputs(testFeatures.User.SocureScore).
		WithBranchId(branchId)
	_, _ = client.OnlineQuery(req, nil)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Branch-Id"), branchId)

	bulkBranchId := "bulk-branch-id"
	bulkReq := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.SocureScore).
		WithBranchId(bulkBranchId)
	_, _ = client.OnlineQueryBulk(bulkReq)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Branch-Id"), bulkBranchId)
}
