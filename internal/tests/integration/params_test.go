package integration

import (
	"encoding/json"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/samber/lo"
	assert "github.com/stretchr/testify/require"
	"testing"
)

// TestParamsSetInFeatherHeader tests that params are threaded
// through to the feather request header. Params tested:
// - Branch ID
// - Query tags
func TestParamsSetInFeatherHeader(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	httpClient := NewInterceptorHTTPClient()
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
	expectedTags := []string{"tags-1", "tags-2"}
	req := chalk.OnlineQueryParams{Tags: expectedTags}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.SocureScore).
		WithBranchId(expectedBranchId)
	_, _ = client.OnlineQueryBulk(req)
	header, headerErr := internal.GetHeaderFromSerializedOnlineQueryBulkBody(httpClient.Intercepted.Body)
	assert.Nil(t, headerErr)
	actualBranchId, ok := header["branch_id"]
	assert.True(t, ok)
	assert.Equal(t, expectedBranchId, actualBranchId)

	context, ok := header["context"].(map[string]any)
	assert.True(t, ok)
	tagsAny, ok := context["tags"].([]any)
	tagsString := lo.Map(tagsAny, func(tag any, _ int) string {
		return tag.(string)
	})
	assert.True(t, ok)
	assert.Equal(t, expectedTags, tagsString)
}

// TestTagsSetInOnlineQuery tests that we set tags in
// online query.
func TestTagsSetInOnlineQuery(t *testing.T) {
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
	expectedTags := []string{"tags-1", "tags-2"}
	req := chalk.OnlineQueryParams{Tags: expectedTags}.
		WithInput(testFeatures.User.Id, "1").
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OnlineQuery(req, nil)
	var request internal.OnlineQueryRequestSerialized
	assert.NoError(t, json.Unmarshal(httpClient.Intercepted.Body, &request))
	assert.Equal(t, expectedTags, request.Context.Tags)
}

// TestTagsSetInOfflineQuery tests that we set tags in
// online query.
func TestTagsSetInOfflineQuery(t *testing.T) {
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
	expectedTags := []string{"tags-1", "tags-2"}
	req := chalk.OfflineQueryParams{Tags: expectedTags}.
		WithInput(testFeatures.User.Id, []any{int64(1)}).
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OfflineQuery(req)
	var request internal.OfflineQueryRequestSerialized
	assert.NoError(t, json.Unmarshal(httpClient.Intercepted.Body, &request))
	assert.Equal(t, expectedTags, request.Tags)
}
