package benchmark

import (
	"context"
	"fmt"
	"github.com/chalk-ai/chalk-go"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func getBenchmarkQueryBulkLoneMultiNsWindowed(b *testing.B) (benchFunc func(), closeFunc func()) {
	type root struct {
		IntFeatures       fixtures.WindowedIntFeatures
		FloatFeatures     fixtures.WindowedFloatFeatures
		BoolFeatures      fixtures.WindowedBoolFeatures
		StringFeatures    fixtures.WindowedStringFeatures
		TimestampFeatures fixtures.WindowedTimestampFeatures
	}

	bulkData := make(map[string]any)
	windows := []int{60, 300, 3600}
	for i := 1; i <= 13; i++ {
		for _, window := range windows {
			bulkData[fmt.Sprintf("windowed_int_features.int_%d__%d__", i, window)] = []int64{122}
			bulkData[fmt.Sprintf("windowed_float_features.float_%d__%d__", i, window)] = []float64{1.234}
			bulkData[fmt.Sprintf("windowed_bool_features.bool_%d__%d__", i, window)] = []bool{true}
			bulkData[fmt.Sprintf("windowed_string_features.string_%d__%d__", i, window)] = []string{"string_val"}
			bulkData[fmt.Sprintf("windowed_timestamp_features.timestamp_%d__%d__", i, window)] = []time.Time{time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC)}
		}
	}

	record, err := internal.ColumnMapToRecord(bulkData)
	assert.NoError(b, err)
	bytes, err := internal.RecordToBytes(record)
	assert.NoError(b, err)
	tf, err := NewTestFixture(&fixtures.MockServerConfig{
		QueryBulkResponse: &commonv1.OnlineQueryBulkResponse{
			ScalarsData: bytes,
		},
	})
	assert.NoError(b, err)
	assertOnce := sync.Once{}
	return func() {
		req := chalk.OnlineQueryParams{}.
			WithInput("user.id", 1).
			WithOutputs("user.id", "user.socure_score")
		res, err := tf.Client.OnlineQueryBulk(context.Background(), req)
		assert.NoError(b, err)
		var results []root
		assert.NoError(b, res.UnmarshalInto(&results))
		assert.Len(b, results, 1)
		firstRes := results[0]
		assertOnce.Do(func() {
			assert.Equal(b, int64(122), *firstRes.IntFeatures.Int1["1m"])
			assert.Equal(b, int64(122), *firstRes.IntFeatures.Int13["1h"])
			assert.Equal(b, float64(1.234), *firstRes.FloatFeatures.Float1["1m"])
			assert.Equal(b, float64(1.234), *firstRes.FloatFeatures.Float13["1h"])
			assert.Equal(b, "string_val", *firstRes.StringFeatures.String1["1m"])
			assert.Equal(b, "string_val", *firstRes.StringFeatures.String13["1h"])
			assert.True(b, *firstRes.BoolFeatures.Bool1["1m"])
			assert.True(b, *firstRes.BoolFeatures.Bool13["1h"])
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *firstRes.TimestampFeatures.Timestamp1["1m"])
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *firstRes.TimestampFeatures.Timestamp13["1h"])
		})
	}, tf.Close
}

/*
 * Query: Bulk (single row)
 * Namespaces: Multi
 * Feature Type: Windowed
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkQueryBulkLoneMultiNsWindowed(b *testing.B) {
	benchFunc, closeFunc := getBenchmarkQueryBulkLoneMultiNsWindowed(b)
	benchmarkWithClose(b, benchFunc, closeFunc)
}

/*
 * Query: Bulk (single row)
 * Namespaces: Multi
 * Feature Type: Windowed
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkQueryBulkLoneMultiNsWindowedParallel(b *testing.B) {
	benchFunc, closeFunc := getBenchmarkQueryBulkLoneMultiNsWindowed(b)
	benchmarkParallelWithClose(b, benchFunc, closeFunc)
}
