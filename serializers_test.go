package chalk

import (
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"testing"
	"time"

	"github.com/chalk-ai/chalk-go/internal/colls"
	assert "github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SerdeDetails struct {
	Name      *string  `dataclass_field:"true"`
	OutAmount *float64 `dataclass_field:"true"`
	InAmount  *float64 `dataclass_field:"true"`
}

type SerdeUser struct {
	Id   *int64
	Txns *[]SerdeTransaction `has_many:"id,location_id"`
}

var SerdeRoot struct {
	SerdeUser *SerdeUser
}

type SerdeTransaction struct {
	Id         *string
	LocationId *int64
	Details    *SerdeDetails `dataclass:"true"`
}

func TestFeatureResultDeserialization(t *testing.T) {
	withTimestamp := featureResultSerialized{
		Field:     "user.id",
		Value:     "1",
		Pkey:      "1",
		Timestamp: "2021-09-01T00:00:00Z",
		Meta:      nil,
		Error:     nil,
	}
	withoutTimestamp := featureResultSerialized{
		Field:     "user.id",
		Value:     "1",
		Pkey:      "1",
		Timestamp: "",
		Meta:      nil,
		Error:     nil,
	}
	tsResult, err := withTimestamp.deserialize()
	assert.NoError(t, err)
	assert.Equal(t, "user.id", tsResult.Field)
	assert.Equal(t, time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC), tsResult.Timestamp)

	noTsResult, err := withoutTimestamp.deserialize()
	assert.NoError(t, err)
	assert.Equal(t, "user.id", noTsResult.Field)
	assert.Equal(t, time.Time{}, noTsResult.Timestamp)
}

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
	queryContext, err := NewQueryContext(map[string]any{"key": "value"})
	assert.NoError(t, err)

	params := OnlineQueryParams{
		IncludeMeta:          true,
		StorePlanStages:      true,
		Explain:              true,
		Tags:                 tags,
		RequiredResolverTags: requiredResolverTags,
		QueryName:            queryName,
		QueryNameVersion:     queryNameVersion,
		CorrelationId:        correlationId,
		QueryContext:         queryContext,
		Meta:                 meta,
		Now:                  now,
	}
	nowProto := colls.Map(now, func(t time.Time) *timestamppb.Timestamp {
		return timestamppb.New(t)
	})
	protoMap, err := queryContext.toProtoMap()
	assert.NoError(t, err)

	request, err := convertOnlineQueryParamsToProto(&params)
	assert.NoError(t, err)
	assert.Equal(t, tags, request.GetContext().GetTags())
	assert.Equal(t, requiredResolverTags, request.GetContext().GetRequiredResolverTags())
	assert.Equal(t, queryName, request.GetContext().GetQueryName())
	assert.Equal(t, queryNameVersion, request.GetContext().GetQueryNameVersion())
	assert.Equal(t, correlationId, request.GetContext().GetCorrelationId())
	assert.Equal(t, meta, request.GetResponseOptions().GetMetadata())
	assert.Equal(t, request.Context.QueryContext, protoMap)
	assert.Equal(t, nowProto, request.GetNow())
	assert.True(t, request.GetResponseOptions().GetIncludeMeta())
	assert.NotNil(t, request.GetResponseOptions().GetExplain())
	optionsActual := request.GetContext().GetOptions()
	assert.True(t, optionsActual["store_plan_stages"].GetBoolValue())
}

/* Test that we can serialize and deserialize a dataclass nested in a features class without loss*/
func TestSerializingDataclassNestedInFeaturesClass(t *testing.T) {
	t.Skip(
		"Need to extend this to test omitting fields for bulk queries. " +
			"For reference, we are testing that for singular queries in " +
			"TestOnlineQueryParamsOmitNilFields",
	)
	assert.NoError(t, InitFeatures(&SerdeRoot))

	transactions := []SerdeTransaction{
		{
			Id:         ptr.Ptr("1"),
			LocationId: ptr.Ptr(int64(1)),
			Details: &SerdeDetails{
				Name:      ptr.Ptr("name"),
				OutAmount: ptr.Ptr(2.2),
				InAmount:  ptr.Ptr(3.3),
			},
		},
		{
			Id:         ptr.Ptr("2"),
			LocationId: ptr.Ptr(int64(2)),
			Details: &SerdeDetails{
				Name:      ptr.Ptr("name2"),
				OutAmount: ptr.Ptr(3.2),
				InAmount:  ptr.Ptr(4.3),
			},
		},
	}
	params := OnlineQueryParams{}.
		WithInput(SerdeRoot.SerdeUser.Id, []int64{1}).
		WithInput(SerdeRoot.SerdeUser.Txns, [][]SerdeTransaction{transactions}).
		WithOutputs(SerdeRoot.SerdeUser.Id, SerdeRoot.SerdeUser.Txns)

	req, err := convertOnlineQueryParamsToProto(&params.underlying)
	assert.NoError(t, err)
	table, err := internal.ConvertBytesToTable(req.GetInputsFeather())
	assert.NoError(t, err)
	assert.Equal(t, int64(1), table.NumRows())
	assert.Equal(t, int64(2), table.NumCols())

	var actualUser []SerdeUser
	bulkRes := OnlineQueryBulkResult{ScalarsTable: table}
	if err := bulkRes.UnmarshalInto(&actualUser); err != nil {
		assert.FailNow(t, "Failed to unmarshal bulk result into user", err)
	}
	assert.Equal(t, 1, len(actualUser))
	assert.Equal(t, transactions, *actualUser[0].Txns)

}
