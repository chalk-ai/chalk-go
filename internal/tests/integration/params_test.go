package integration

import (
	"context"
	"encoding/json"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
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
	correlationId := "correlating-id"
	queryName := "fraud_fighter"
	queryNameVersion := "fraud_fighter_first"
	meta := map[string]string{
		"abTestId": "bee",
		"bbTestId": "cee",
	}

	plannerOptions := map[string]any{
		"CHALK_SOME_PLANNER_OPTION": "some_value",
	}

	httpClient := NewInterceptorHTTPClient()
	client, err := chalk.NewClient(&chalk.ClientConfig{
		HTTPClient: httpClient,
	})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	userIds := []int{1}
	queryContext, err := chalk.NewQueryContext(map[string]any{"key": "value"})
	assert.NoError(t, err)

	req := chalk.OnlineQueryParams{
		Tags:                 expectedTags,
		RequiredResolverTags: requiredResolverTags,
		Now:                  now,
		StorePlanStages:      true,
		CorrelationId:        correlationId,
		QueryName:            queryName,
		QueryNameVersion:     queryNameVersion,
		Meta:                 meta,
		Explain:              true,
		IncludeMeta:          true,
		QueryContext:         queryContext,
		PlannerOptions:       plannerOptions,
	}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.SocureScore).
		WithBranchId(expectedBranchId)
	for k, v := range staleness {
		req = req.WithStaleness(k, v)
	}

	_, _ = client.OnlineQueryBulk(context.Background(), req)
	headerMap, headerErr := internal.GetHeaderFromSerializedOnlineQueryBulkBody(httpClient.Intercepted.Body)
	assert.Nil(t, headerErr)

	headerJson, err := json.Marshal(headerMap)
	assert.NoError(t, err)
	var header internal.FeatherRequestHeader
	assert.NoError(t, json.Unmarshal(headerJson, &header))

	stalenessConverted := make(map[string]string)
	for k, v := range staleness {
		feature, err := chalk.UnwrapFeature(k)
		assert.NoError(t, err)
		stalenessConverted[feature.Fqn] = internal.FormatBucketDuration(int(v.Seconds()))
	}

	nowConverted := colls.Map(now, func(val time.Time) string {
		return val.Format(internal.NowTimeFormat)
	})

	assert.Equal(t, expectedBranchId, *header.BranchId)
	assert.Equal(t, expectedTags, header.Context.Tags)
	assert.Equal(t, requiredResolverTags, header.Context.RequiredResolverTags)
	assert.Equal(t, nowConverted, header.Now)
	assert.Equal(t, stalenessConverted, header.Staleness)
	assert.True(t, header.StorePlanStages)
	assert.NotNil(t, header.CorrelationId)
	assert.Equal(t, correlationId, *header.CorrelationId)
	assert.NotNil(t, header.QueryName)
	assert.Equal(t, queryName, *header.QueryName)
	assert.NotNil(t, header.QueryNameVersion)
	assert.Equal(t, queryNameVersion, *header.QueryNameVersion)
	assert.Equal(t, meta, header.Meta)
	assert.True(t, header.Explain)
	assert.True(t, header.IncludeMeta)
	assert.NotNil(t, header.QueryContext)
	assert.Equal(t, &map[string]any{"key": "value"}, header.QueryContext)
	assert.Equal(t, plannerOptions, header.PlannerOptions)
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
	correlationId := "correlating-id"
	queryName := "fraud_fighter"
	queryNameVersion := "fraud_fighter_first"
	meta := map[string]string{
		"abTestId": "bee",
		"bbTestId": "cee",
	}
	queryContext, err := chalk.NewQueryContext(map[string]any{"key": "value"})
	assert.NoError(t, err)

	stalenessConverted := make(map[string]string)
	staleness = map[any]time.Duration{}

	nowConverted := colls.Map(now, func(val time.Time) string {
		return val.Format(internal.NowTimeFormat)
	})

	plannerOption := map[string]any{"CHALK_SOME_PLANNER_OPTION": "1"}

	req := chalk.OnlineQueryParams{
		Tags:                 expectedTags,
		RequiredResolverTags: requiredResolverTags,
		StorePlanStages:      true,
		Meta:                 meta,
		QueryName:            queryName,
		QueryNameVersion:     queryNameVersion,
		CorrelationId:        correlationId,
		Now:                  now,
		Explain:              true,
		IncludeMeta:          true,
		IncludeMetrics:       true,
		QueryContext:         queryContext,
		EncodingOptions: &chalk.FeatureEncodingOptions{
			EncodeStructsAsObjects: true,
		},
		PlannerOptions: plannerOption,
	}.
		WithInput(testFeatures.User.Id, "1").
		WithOutputs(testFeatures.User.SocureScore)

	for k, v := range staleness {
		req = req.WithStaleness(k, v)
	}

	_, _ = client.OnlineQuery(context.Background(), req, nil)
	var request internal.OnlineQueryRequestSerialized
	assert.NoError(t, json.Unmarshal(httpClient.Intercepted.Body, &request))
	assert.Equal(t, expectedTags, request.Context.Tags)

	assert.Equal(t, requiredResolverTags, request.Context.RequiredResolverTags)
	assert.NotNil(t, request.Now)
	assert.Equal(t, nowConverted[0], *request.Now)
	assert.Equal(t, stalenessConverted, request.Staleness)
	assert.True(t, request.StorePlanStages)
	assert.NotNil(t, request.CorrelationId)
	assert.Equal(t, correlationId, *request.CorrelationId)
	assert.NotNil(t, request.QueryName)
	assert.Equal(t, queryName, *request.QueryName)
	assert.NotNil(t, request.QueryNameVersion)
	assert.Equal(t, queryNameVersion, *request.QueryNameVersion)
	assert.Equal(t, meta, request.Meta)
	assert.True(t, request.Explain)
	assert.True(t, request.IncludeMeta)
	assert.True(t, request.IncludeMetrics)
	assert.True(t, request.EncodingOptions.EncodeStructsAsObjects)
	assert.NotNil(t, request.QueryContext)
	assert.Equal(t, &map[string]any{"key": "value"}, request.QueryContext)
	assert.Equal(t, plannerOption, request.PlannerOptions)
}

// TestParamsSetInOfflineQuery tests that we set params in
// online query.
func TestParamsSetInOfflineQuery(t *testing.T) {
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
	queryContext, err := chalk.NewQueryContext(map[string]any{"key": "value"})
	assert.NoError(t, err)
	expectedTags := []string{"tags-1", "tags-2"}
	req := chalk.OfflineQueryParams{
		Tags:         expectedTags,
		QueryContext: queryContext,
	}.
		WithInput(testFeatures.User.Id, []any{int64(1)}).
		WithOutputs(testFeatures.User.SocureScore)
	_, _ = client.OfflineQuery(context.Background(), req)
	var request internal.OfflineQueryRequestSerialized
	assert.NoError(t, json.Unmarshal(httpClient.Intercepted.Body, &request))
	assert.Equal(t, expectedTags, request.Tags)
	assert.NotNil(t, request.QueryContext)
	assert.Equal(t, request.QueryContext, &map[string]any{"key": "value"})
}
