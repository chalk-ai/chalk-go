package integration

import (
	"bytes"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal"
	assert "github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

type Intercepted struct {
	Header http.Header
	Body   []byte
}

type InterceptorHTTPClient struct {
	Intercepted Intercepted
}

func NewCovertHTTPClient() *InterceptorHTTPClient {
	return &InterceptorHTTPClient{}
}

func (c *InterceptorHTTPClient) Do(req *http.Request) (*http.Response, error) {
	bodyBytes, bodyBytesErr := io.ReadAll(req.Body)
	if bodyBytesErr != nil {
		return nil, bodyBytesErr
	}
	req.Body.Close()
	c.Intercepted = Intercepted{
		Header: req.Header,
		Body:   bodyBytes,
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
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
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

// TestOnlineQueryBranchInClient tests that when we
// specify a branch ID in the client, the online query
// request includes the branch ID header.
func TestOnlineQueryBranchInClient(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewCovertHTTPClient()
	branchId := "test-branch-id"
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient: httpClient,
		Branch:     branchId,
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
	_, _ = client.OnlineQuery(req, nil)
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Branch-Id"), branchId)
}

// TestOnlineQueryBulkBranchInClient tests that when we
// specify a branch ID in the client, the bulk online query
// request includes the branch ID header.
func TestOnlineQueryBulkBranchInClient(t *testing.T) {
	// TODO: This can be a non-integration test if we can make
	//       the mock client return a fake JWT when auth is
	//       being performed.
	SkipIfNotIntegrationTester(t)
	httpClient := NewCovertHTTPClient()
	branchId := "test-branch-id"
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient: httpClient,
		Branch:     branchId,
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
	assert.Equal(t, httpClient.Intercepted.Header.Get("X-Chalk-Branch-Id"), branchId)
}

// TestClientBranchSetInFeatherHeader tests that when we
// specify a branch ID in the client, the feather request
// header that we serialize includes the branch ID header.
func TestClientBranchSetInFeatherHeader(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	httpClient := NewCovertHTTPClient()
	expectedBranchId := "test-branch-id"
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient: httpClient,
		Branch:     expectedBranchId,
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
	header, headerErr := internal.GetHeaderFromSerializedOnlineQueryBulkBody(httpClient.Intercepted.Body)
	assert.Nil(t, headerErr)
	actualBranchId, ok := header["branch_id"]
	assert.True(t, ok)
	assert.Equal(t, expectedBranchId, actualBranchId)
}

// TestParamsBranchSetInFeatherHeader tests that when we
// specify a branch ID in the params, the feather request
// header that we serialize includes the branch ID header.
func TestParamsBranchSetInFeatherHeader(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	httpClient := NewCovertHTTPClient()
	expectedBranchId := "test-branch-id"
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
		WithOutputs(testFeatures.User.SocureScore).
		WithBranchId(expectedBranchId)
	_, _ = client.OnlineQueryBulk(req)
	header, headerErr := internal.GetHeaderFromSerializedOnlineQueryBulkBody(httpClient.Intercepted.Body)
	assert.Nil(t, headerErr)
	actualBranchId, ok := header["branch_id"]
	assert.True(t, ok)
	assert.Equal(t, expectedBranchId, actualBranchId)
}
