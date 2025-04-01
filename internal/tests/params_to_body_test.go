package tests

import (
	"context"
	"encoding/json"
	chalk "github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestParamsSetInFeatherHeader tests that params are threaded
// through to the feather request header.
func TestParamsSetInFeatherHeader(t *testing.T) {
	expectedBranchId := "test-branch-id"
	expectedTags := []string{"tags-1", "tags-2"}
	requiredResolverTags := []string{"required1tag", "required-2tag"}
	now := []time.Time{time.Now(), time.Now()}
	staleness := map[any]time.Duration{
		"bogus.socure_score": time.Minute,
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

	client, httpClient, err := newClientWithInterceptor()
	assert.NoError(t, err)

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
		WithInput("bogus.feature", userIds).
		WithOutputs("bogus.output").
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
		stalenessConverted[k.(string)] = internal.FormatBucketDuration(int(v.Seconds()))
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

// TestParamsSetInOnlineQueryBody tests that we set all params
// correctly in online query request body.
func TestParamsSetInOnlineQueryBody(t *testing.T) {
	client, httpClient, err := newClientWithInterceptor()
	assert.NoError(t, err)

	expectedTags := []string{"tags-1", "tags-2"}
	requiredResolverTags := []string{"required1tag", "required-2tag"}
	now := []time.Time{time.Now()}
	staleness := map[any]time.Duration{}
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
		QueryContext:         queryContext,
		PlannerOptions:       plannerOption,
	}.
		WithInput("bogus.feature", 1).
		WithOutputs("bogus.output")

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
	assert.True(t, request.EncodingOptions.EncodeStructsAsObjects)
	assert.NotNil(t, request.QueryContext)
	assert.Equal(t, &map[string]any{"key": "value"}, request.QueryContext)
	assert.Equal(t, plannerOption, request.PlannerOptions)
}

// TestParamsSetInOfflineQuery tests that offline query
// params gets propagated through to the request body.
func TestParamsSetInOfflineQueryBody(t *testing.T) {
	client, httpClient, err := newClientWithInterceptor()
	assert.NoError(t, err)

	queryContext, err := chalk.NewQueryContext(map[string]any{"key": "value"})
	assert.NoError(t, err)
	expectedTags := []string{"tags-1", "tags-2"}
	req := chalk.OfflineQueryParams{
		Tags:         expectedTags,
		QueryContext: queryContext,
	}.
		WithInput("bogus.feature", []any{int64(1)}).
		WithOutputs("bogus.output")
	_, _ = client.OfflineQuery(context.Background(), req)
	var request internal.OfflineQueryRequestSerialized
	assert.NoError(t, json.Unmarshal(httpClient.Intercepted.Body, &request))
	assert.Equal(t, expectedTags, request.Tags)
	assert.NotNil(t, request.QueryContext)
	assert.Equal(t, request.QueryContext, &map[string]any{"key": "value"})
}

// TestClientBranchSetInFeatherHeader tests that when we
// specify a branch ID in the client, the feather request
// header that we serialize includes the branch ID header.
func TestClientBranchSetInFeatherHeader(t *testing.T) {
	expectedBranchId := "test-branch-id"
	client, httpClient, err := newClientWithInterceptor(interceptorClientOverrides{
		Branch: expectedBranchId,
	})
	assert.NoError(t, err)
	
	req := chalk.OnlineQueryParams{}.
		WithInput("bogus.feature", []int{1}).
		WithOutputs("bogus.output")
	_, _ = client.OnlineQueryBulk(context.Background(), req)
	
	headerMap, headerErr := internal.GetHeaderFromSerializedOnlineQueryBulkBody(httpClient.Intercepted.Body)
	assert.Nil(t, headerErr)
	
	branchId, ok := headerMap["branch_id"]
	assert.True(t, ok)
	assert.Equal(t, expectedBranchId, branchId)
}
