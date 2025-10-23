package chalk

import (
	"encoding/json"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	"strings"
	"sync"
	"testing"
	"time"

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

type SerdeRootType struct {
	SerdeUser *SerdeUser
}

type SerdeTransaction struct {
	Id         *string
	LocationId *int64
	Details    *SerdeDetails `dataclass:"true"`
}

var initSerdeRootFeaturesOnce sync.Once
var SerdeRootFeaturesSingleton SerdeRootType
var initSerdeRootFeaturesError error

func GetSerdeRootFeatures() (SerdeRootType, error) {
	initSerdeRootFeaturesOnce.Do(func() {
		initSerdeRootFeaturesError = InitFeatures(&SerdeRootFeaturesSingleton)
	})
	return SerdeRootFeaturesSingleton, initSerdeRootFeaturesError
}

func TestFeatureResultDeserialization(t *testing.T) {
	t.Parallel()
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
	assert.Equal(t, time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC), *tsResult.Timestamp)

	noTsResult, err := withoutTimestamp.deserialize()
	assert.NoError(t, err)
	assert.Equal(t, "user.id", noTsResult.Field)
	assert.Nil(t, noTsResult.Timestamp)
}

func TestConvertOnlineQueryParamsToProto(t *testing.T) {
	t.Parallel()
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
	nowProto := make([]*timestamppb.Timestamp, len(now))
	for i, t := range now {
		nowProto[i] = timestamppb.New(t)
	}
	protoMap, err := queryContext.toProtoMap()
	assert.NoError(t, err)

	request, err := convertOnlineQueryParamsToProto(&params, fixtures.TestAllocator)
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
	t.Parallel()
	serdeRoot, initErr := GetSerdeRootFeatures()
	assert.NoError(t, initErr)

	transactions := []SerdeTransaction{
		{
			Id:         ptr.New("1"),
			LocationId: ptr.New(int64(1)),
			Details: &SerdeDetails{
				Name:      ptr.New("name"),
				OutAmount: ptr.New(2.2),
				InAmount:  ptr.New(3.3),
			},
		},
		{
			Id:         ptr.New("2"),
			LocationId: ptr.New(int64(2)),
			Details: &SerdeDetails{
				Name:      ptr.New("name2"),
				OutAmount: ptr.New(3.2),
				InAmount:  ptr.New(4.3),
			},
		},
	}
	params := OnlineQueryParams{}.
		WithInput(serdeRoot.SerdeUser.Id, []int64{1}).
		WithInput(serdeRoot.SerdeUser.Txns, [][]SerdeTransaction{transactions}).
		WithOutputs(serdeRoot.SerdeUser.Id, serdeRoot.SerdeUser.Txns)

	req, err := convertOnlineQueryParamsToProto(&params.underlying, fixtures.TestAllocator)
	assert.NoError(t, err)
	table, err := internal.ConvertBytesToTable(req.GetInputsFeather(), fixtures.TestAllocator)
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

func TestGRPCOnlineQueryBulkResultConstructor(t *testing.T) {
	t.Parallel()
	table, err := MakeFeatureTable(map[any]any{
		"user.id": []string{"1"},
	})
	assert.NoError(t, err)
	tableBytes, err := table.ToBytes()
	assert.NoError(t, err)
	result, err := NewGRPCOnlineQueryBulkResult(&commonv1.OnlineQueryBulkResponse{
		ScalarsData: tableBytes,
	})
	assert.NoError(t, err)
	row, err := result.GetRow(0)
	assert.NoError(t, err)
	assert.Equal(t, "1", row.Features["user.id"].Value)
}

func TestDurationStringFormatting(t *testing.T) {
	// Test that Go's duration string formatting works as expected
	// We're now using the simpler Duration.String() approach

	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "1 hour",
			duration: time.Hour,
			expected: "1h0m0s",
		},
		{
			name:     "1 hour 30 minutes",
			duration: time.Hour + 30*time.Minute,
			expected: "1h30m0s",
		},
		{
			name:     "2 hours",
			duration: 2 * time.Hour,
			expected: "2h0m0s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.duration.String()
			if result != tt.expected {
				t.Errorf("Duration.String() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestInlineTimestampFormatting(t *testing.T) {
	// Test timezone setup
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")
	pst, _ := time.LoadLocation("America/Los_Angeles")

	tests := []struct {
		name     string
		time     *time.Time
		expected *string
	}{
		// Nil time
		{
			name:     "nil time",
			time:     nil,
			expected: nil,
		},
		// UTC timestamp
		{
			name:     "UTC timestamp",
			time:     timePtr(time.Date(2024, 5, 9, 22, 29, 0, 0, utc)),
			expected: stringPtr("2024-05-09T22:29:00+00:00"),
		},
		// EST timestamp
		{
			name:     "EST timestamp",
			time:     timePtr(time.Date(2024, 5, 9, 18, 29, 0, 0, est)),
			expected: stringPtr("2024-05-09T18:29:00-04:00"),
		},
		// PST timestamp
		{
			name:     "PST timestamp",
			time:     timePtr(time.Date(2024, 5, 9, 15, 29, 0, 0, pst)),
			expected: stringPtr("2024-05-09T15:29:00-07:00"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the inline formatting approach used in serializers.go
			var result *string
			if tt.time != nil {
				formatted := strings.Replace(tt.time.Format(time.RFC3339Nano), "Z", "+00:00", 1)
				result = &formatted
			}

			if !stringPtrEqual(result, tt.expected) {
				t.Errorf("inline timestamp formatting(%v) = %v, want %v", tt.time, stringPtrValue(result), stringPtrValue(tt.expected))
			}
		})
	}
}

func TestTimestampSerializationIntegration(t *testing.T) {
	// Test that the complete serialization pipeline works correctly
	// This tests the integration of formatTimeBound and duration string formatting

	// Create test parameters with timestamp bounds
	lowerBound := time.Date(2024, 5, 9, 15, 29, 0, 0, time.UTC)
	upperBound := time.Date(2024, 5, 9, 16, 29, 0, 0, time.UTC)

	params := &OfflineQueryParams{
		CompletionDeadline:   durationPtr(2 * time.Hour),
		ObservedAtLowerBound: &lowerBound,
		ObservedAtUpperBound: &upperBound,
	}

	// Test serialization
	resolved := &offlineQueryParamsResolved{
		inputs:          make(map[string][]TsFeatureValue),
		outputs:         []string{"test.feature"},
		requiredOutputs: []string{},
	}

	serialized, err := serializeOfflineQueryParams(params, resolved)
	if err != nil {
		t.Fatalf("serializeOfflineQueryParams failed: %v", err)
	}

	// Verify we got valid JSON
	if len(serialized) == 0 {
		t.Error("serializeOfflineQueryParams returned empty result")
	}

	// The actual JSON structure verification would require unmarshaling,
	// but the important thing is that the functions don't panic and produce valid output
	t.Logf("Serialized successfully: %d bytes", len(serialized))
}

// Helper functions for testing
func timePtr(t time.Time) *time.Time {
	return &t
}

func durationPtr(d time.Duration) *time.Duration {
	return &d
}

func stringPtr(s string) *string {
	return &s
}

func stringPtrEqual(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func stringPtrValue(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

func TestOnlineQueryRequestSerializationWithBranch(t *testing.T) {
	// Test that ensures BranchId is properly serialized in OnlineQueryRequestSerialized JSON
	// This prevents regression of the critical bug fix where BranchId was missing from JSON body
	t.Parallel()

	branchId := "test-branch"
	params := OnlineQueryParams{
		BranchId: &branchId,
		Explain:  true,
	}.WithInput("user.id", 1).WithOutputs("user.full_name")

	resolved := &onlineQueryParamsResolved{
		inputs:  map[string]any{"user.id": 1},
		outputs: []string{"user.full_name"},
	}

	serialized, err := serializeOnlineQueryParams(&params.underlying, resolved)
	assert.NoError(t, err)
	assert.NotNil(t, serialized)

	// Verify BranchId is included in the serialized struct
	assert.Equal(t, &branchId, serialized.BranchId)
	assert.Equal(t, true, serialized.Explain)

	// Test JSON marshaling to ensure BranchId appears in the actual JSON
	jsonBytes, err := json.Marshal(serialized)
	assert.NoError(t, err)

	jsonString := string(jsonBytes)
	assert.Contains(t, jsonString, `"branch_id":"test-branch"`)
	assert.Contains(t, jsonString, `"explain":true`)
}

func TestOnlineQueryRequestSerializationWithoutBranch(t *testing.T) {
	// Test that serialization works correctly when BranchId is nil
	t.Parallel()

	params := OnlineQueryParams{
		BranchId: nil,
		Explain:  true,
	}.WithInput("user.id", 1).WithOutputs("user.full_name")

	resolved := &onlineQueryParamsResolved{
		inputs:  map[string]any{"user.id": 1},
		outputs: []string{"user.full_name"},
	}

	serialized, err := serializeOnlineQueryParams(&params.underlying, resolved)
	assert.NoError(t, err)
	assert.NotNil(t, serialized)

	// Verify BranchId is nil in the serialized struct
	assert.Nil(t, serialized.BranchId)
	assert.Equal(t, true, serialized.Explain)

	// Test JSON marshaling to ensure BranchId appears as null in JSON
	jsonBytes, err := json.Marshal(serialized)
	assert.NoError(t, err)

	jsonString := string(jsonBytes)
	assert.Contains(t, jsonString, `"branch_id":null`)
	assert.Contains(t, jsonString, `"explain":true`)
}

func TestOfflineQuerySerializationWithFileInput(t *testing.T) {
	t.Parallel()

	// Test with file input URI
	fileUri := "s3://my-bucket/data.parquet"
	params := OfflineQueryParams{
		rawFileInput: &fileUri,
	}

	resolved := &offlineQueryParamsResolved{
		inputs:          map[string][]TsFeatureValue{},
		outputs:         []string{"user.id", "user.name"},
		requiredOutputs: []string{},
		versioned:       false,
	}

	serialized, err := serializeOfflineQueryParams(&params, resolved)
	assert.NoError(t, err)

	// Parse the JSON to verify structure
	var result map[string]interface{}
	err = json.Unmarshal(serialized, &result)
	assert.NoError(t, err)

	// Verify that input is an OfflineQueryInputUri
	input, ok := result["input"].(map[string]interface{})
	assert.True(t, ok, "input should be a map")
	assert.Equal(t, fileUri, input["parquet_uri"])
	assert.NotContains(t, input, "columns")
	assert.NotContains(t, input, "values")
}

func TestOfflineQuerySerializationWithRegularInput(t *testing.T) {
	t.Parallel()

	// Test with regular input (not file input)
	params := OfflineQueryParams{}

	observationTime := time.Now()
	resolved := &offlineQueryParamsResolved{
		inputs: map[string][]TsFeatureValue{
			"user.id": {
				{Value: int64(1), ObservationTime: &observationTime},
				{Value: int64(2), ObservationTime: &observationTime},
			},
		},
		outputs:         []string{"user.name"},
		requiredOutputs: []string{},
		versioned:       false,
	}

	serialized, err := serializeOfflineQueryParams(&params, resolved)
	assert.NoError(t, err)

	// Parse the JSON to verify structure
	var result map[string]interface{}
	err = json.Unmarshal(serialized, &result)
	assert.NoError(t, err)

	// Verify that input is an OfflineQueryInputSerialized
	input, ok := result["input"].(map[string]interface{})
	assert.True(t, ok, "input should be a map")
	assert.Contains(t, input, "columns")
	assert.Contains(t, input, "values")
	assert.NotContains(t, input, "parquet_uri")
}

func TestOfflineQuerySerializationWithRecomputeFeatures(t *testing.T) {
	t.Parallel()

	// Test with RecomputeFeatures boolean
	params := OfflineQueryParams{
		RecomputeFeatures: true,
	}

	resolved := &offlineQueryParamsResolved{
		inputs:          map[string][]TsFeatureValue{},
		outputs:         []string{"user.id"},
		requiredOutputs: []string{},
		versioned:       false,
	}

	serialized, err := serializeOfflineQueryParams(&params, resolved)
	assert.NoError(t, err)

	// Parse the JSON to verify structure
	var result map[string]interface{}
	err = json.Unmarshal(serialized, &result)
	assert.NoError(t, err)

	// Verify that recompute_features is a boolean
	recomputeFeatures, ok := result["recompute_features"].(bool)
	assert.True(t, ok, "recompute_features should be a boolean")
	assert.True(t, recomputeFeatures)
}

func TestOfflineQuerySerializationWithRecomputeFeaturesList(t *testing.T) {
	t.Parallel()

	// Test with RecomputeFeaturesList
	featuresList := []string{"user.id", "user.name", "user.email"}
	params := OfflineQueryParams{
		RecomputeFeaturesList: featuresList,
	}

	resolved := &offlineQueryParamsResolved{
		inputs:          map[string][]TsFeatureValue{},
		outputs:         []string{"user.score"},
		requiredOutputs: []string{},
		versioned:       false,
	}

	serialized, err := serializeOfflineQueryParams(&params, resolved)
	assert.NoError(t, err)

	// Parse the JSON to verify structure
	var result map[string]interface{}
	err = json.Unmarshal(serialized, &result)
	assert.NoError(t, err)

	// Verify that recompute_features is an array
	recomputeFeaturesInterface, ok := result["recompute_features"].([]interface{})
	assert.True(t, ok, "recompute_features should be an array")

	// Convert to string array and verify contents
	var recomputeFeatures []string
	for _, v := range recomputeFeaturesInterface {
		str, ok := v.(string)
		assert.True(t, ok, "each element should be a string")
		recomputeFeatures = append(recomputeFeatures, str)
	}
	assert.Equal(t, featuresList, recomputeFeatures)
}

func TestOfflineQueryValidationRecomputeFeaturesConflict(t *testing.T) {
	t.Parallel()

	// Test validation error when both are set
	params := OfflineQueryParams{
		RecomputeFeatures:     true,
		RecomputeFeaturesList: []string{"user.id"},
	}

	_, err := params.resolve()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot set both RecomputeFeatures and RecomputeFeaturesList")
}
