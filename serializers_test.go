package chalk

import (
	"testing"
	"time"

	"github.com/chalk-ai/chalk-go/internal/colls"
	assert "github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestConvertOnlineQueryParamsToProto(t *testing.T) {
	tags := []string{"tag1", "tag2"}
	requiredResolverTags := []string{"tag3", "tag4"}
	queryName := "test-name"
	queryNameVersion := "1"
	correlationId := "test-correlation-id"
	meta := map[string]string{
		"test-meta-1": "test-meta-value-1",
	}
	now := []time.Time{time.Now()}
	params := OnlineQueryParams{
		IncludeMeta:          true,
		IncludeMetrics:       true,
		StorePlanStages:      true,
		Explain:              true,
		Tags:                 tags,
		RequiredResolverTags: requiredResolverTags,
		QueryName:            queryName,
		QueryNameVersion:     queryNameVersion,
		CorrelationId:        correlationId,
		Meta:                 meta,
		Now:                  now,
	}
	nowProto := colls.Map(now, func(t time.Time) *timestamppb.Timestamp {
		return timestamppb.New(t)
	})
	request, err := convertOnlineQueryParamsToProto(&params)
	assert.NoError(t, err)
	assert.Equal(t, tags, request.GetContext().GetTags())
	assert.Equal(t, requiredResolverTags, request.GetContext().GetRequiredResolverTags())
	assert.Equal(t, queryName, request.GetContext().GetQueryName())
	assert.Equal(t, queryNameVersion, request.GetContext().GetQueryNameVersion())
	assert.Equal(t, correlationId, request.GetContext().GetCorrelationId())
	assert.Equal(t, meta, request.GetResponseOptions().GetMetadata())
	assert.Equal(t, nowProto, request.GetNow())
	assert.True(t, request.GetResponseOptions().GetIncludeMeta())
	assert.NotNil(t, request.GetResponseOptions().GetExplain())
	optionsActual := request.GetContext().GetOptions()
	assert.True(t, optionsActual["include_metrics"].GetBoolValue())
	assert.True(t, optionsActual["store_plan_stages"].GetBoolValue())
}
