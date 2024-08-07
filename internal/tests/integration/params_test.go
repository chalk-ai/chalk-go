package integration

import (
	"encoding/json"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/samber/lo"
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestParamsSetInFeatherHeader tests that params are threaded
// through to the feather request header.
func TestParamsSetInFeatherHeader(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	err := chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}

	expectedBranchId := "test-branch-id"
	expectedTags := []string{"tags-1", "tags-2"}
	requiredResolverTags := []string{"required1tag", "required-2tag"}
	now := []time.Time{time.Now(), time.Now()}
	staleness := map[any]time.Duration{
		testFeatures.User.SocureScore: time.Minute,
	}
	storePlanStages := true
	correlationId := "correlating-id"
	queryName := "fraud_fighter"
	queryNameVersion := "fraud_fighter_first"
	meta := map[string]string{
		"abTestId": "bee",
		"bbTestId": "cee",
	}

	httpClient := NewInterceptorHTTPClient()
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient: httpClient,
	})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	userIds := []int{1}

	req := chalk.OnlineQueryParams{
		Tags:                 expectedTags,
		RequiredResolverTags: requiredResolverTags,
		Now:                  now,
		StorePlanStages:      storePlanStages,
		CorrelationId:        correlationId,
		QueryName:            queryName,
		QueryNameVersion:     queryNameVersion,
		Meta:                 meta,
	}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.SocureScore).
		WithBranchId(expectedBranchId)
	for k, v := range staleness {
		req = req.WithStaleness(k, v)
	}

	_, _ = client.OnlineQueryBulk(req)
	headerMap, headerErr := internal.GetHeaderFromSerializedOnlineQueryBulkBody(httpClient.Intercepted.Body)
	assert.Nil(t, headerErr)

	headerJson, err := json.Marshal(headerMap)
	assert.NoError(t, err)
	var header internal.FeatherRequestHeader
	assert.NoError(t, json.Unmarshal(headerJson, &header))

	stalenessConverted := lo.MapEntries(staleness, func(key any, val time.Duration) (string, string) {
		feature, err := chalk.UnwrapFeature(key)
		assert.NoError(t, err)
		return feature.Fqn, internal.FormatBucketDuration(int(val.Seconds()))
	})
	nowConverted := lo.Map(now, func(val time.Time, _ int) string {
		return val.Format(time.RFC3339)
	})

	assert.Equal(t, expectedBranchId, *header.BranchId)
	assert.Equal(t, expectedTags, header.Context.Tags)
	assert.Equal(t, requiredResolverTags, header.Context.RequiredResolverTags)
	assert.Equal(t, nowConverted, header.Now)
	assert.Equal(t, stalenessConverted, header.Staleness)
	assert.Equal(t, storePlanStages, header.StorePlanStages)
	assert.NotNil(t, header.CorrelationId)
	assert.Equal(t, correlationId, *header.CorrelationId)
	assert.NotNil(t, header.QueryName)
	assert.Equal(t, queryName, *header.QueryName)
	assert.NotNil(t, header.QueryNameVersion)
	assert.Equal(t, queryNameVersion, *header.QueryNameVersion)
	assert.Equal(t, meta, header.Meta)
}

// TestParamsSetInOnlineQuery tests that we set all params
// correctly in online query.
func TestParamsSetInOnlineQuery(t *testing.T) {
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
	requiredResolverTags := []string{"required1tag", "required-2tag"}
	now := []time.Time{time.Now()}
	staleness := map[any]time.Duration{
		testFeatures.User.SocureScore: time.Minute,
	}
	storePlanStages := true
	correlationId := "correlating-id"
	queryName := "fraud_fighter"
	queryNameVersion := "fraud_fighter_first"
	meta := map[string]string{
		"abTestId": "bee",
		"bbTestId": "cee",
	}

	stalenessConverted := lo.MapEntries(staleness, func(key any, val time.Duration) (string, string) {
		feature, err := chalk.UnwrapFeature(key)
		assert.NoError(t, err)
		return feature.Fqn, internal.FormatBucketDuration(int(val.Seconds()))
	})
	nowConverted := lo.Map(now, func(val time.Time, _ int) string {
		return val.Format(time.RFC3339)
	})

	req := chalk.OnlineQueryParams{
		Tags:                 expectedTags,
		RequiredResolverTags: requiredResolverTags,
		StorePlanStages:      storePlanStages,
		Meta:                 meta,
		QueryName:            queryName,
		QueryNameVersion:     queryNameVersion,
		CorrelationId:        correlationId,
		Now:                  now,
	}.
		WithInput(testFeatures.User.Id, "1").
		WithOutputs(testFeatures.User.SocureScore)

	for k, v := range staleness {
		req = req.WithStaleness(k, v)
	}

	_, _ = client.OnlineQuery(req, nil)
	var request internal.OnlineQueryRequestSerialized
	assert.NoError(t, json.Unmarshal(httpClient.Intercepted.Body, &request))
	assert.Equal(t, expectedTags, request.Context.Tags)

	assert.Equal(t, requiredResolverTags, request.Context.RequiredResolverTags)
	assert.NotNil(t, request.Now)
	assert.Equal(t, nowConverted[0], *request.Now)
	assert.Equal(t, stalenessConverted, request.Staleness)
	assert.Equal(t, storePlanStages, request.StorePlanStages)
	assert.NotNil(t, request.CorrelationId)
	assert.Equal(t, correlationId, *request.CorrelationId)
	assert.NotNil(t, request.QueryName)
	assert.Equal(t, queryName, *request.QueryName)
	assert.NotNil(t, request.QueryNameVersion)
	assert.Equal(t, queryNameVersion, *request.QueryNameVersion)
	assert.Equal(t, meta, request.Meta)
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
