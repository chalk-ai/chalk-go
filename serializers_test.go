package chalk

import (
	"github.com/samber/lo"
	assert "github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestConvertOnlineQueryParamsToProto(t *testing.T) {
	environmentId := "test-env"
	branchId := "test-branch"
	tags := []string{"tag1", "tag2"}
	requiredResolverTags := []string{"tag3", "tag4"}
	queryName := "test-name"
	queryNameVersion := "1"
	deploymentId := "test-deployment-id"
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
		EnvironmentId:        environmentId,
		BranchId:             &branchId,
		Tags:                 tags,
		RequiredResolverTags: requiredResolverTags,
		QueryName:            queryName,
		QueryNameVersion:     queryNameVersion,
		PreviewDeploymentId:  deploymentId,
		CorrelationId:        correlationId,
		Meta:                 meta,
		Now:                  now,
	}
	nowProto := lo.Map(now, func(t time.Time, _ int) *timestamppb.Timestamp {
		return timestamppb.New(t)
	})
	request, err := convertOnlineQueryParamsToProto(&params)
	assert.NoError(t, err)
	assert.Equal(t, environmentId, request.GetContext().GetEnvironment())
	assert.Equal(t, tags, request.GetContext().GetTags())
	assert.Equal(t, requiredResolverTags, request.GetContext().GetRequiredResolverTags())
	assert.Equal(t, queryName, request.GetContext().GetQueryName())
	assert.Equal(t, queryNameVersion, request.GetContext().GetQueryNameVersion())
	assert.Equal(t, correlationId, request.GetContext().GetCorrelationId())
	assert.Equal(t, meta, request.GetResponseOptions().GetMetadata())
	assert.Equal(t, nowProto, request.GetNow())
	assert.Equal(t, deploymentId, request.GetContext().GetDeploymentId())
	assert.Equal(t, branchId, request.GetContext().GetBranchId())
	assert.True(t, request.GetResponseOptions().GetIncludeMeta())
	assert.NotNil(t, request.GetResponseOptions().GetExplain())
	optionsActual := request.GetContext().GetOptions()
	assert.True(t, optionsActual["include_metrics"].GetBoolValue())
	assert.True(t, optionsActual["store_plan_stages"].GetBoolValue())
}
