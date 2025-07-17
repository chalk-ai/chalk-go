package benchmark

import (
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
		fixtures.WindowedIntFeatures
		fixtures.WindowedFloatFeatures
		fixtures.WindowedBoolFeatures
		fixtures.WindowedStringFeatures
		fixtures.WindowedTimestampFeatures
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

	bytes, err := internal.InputsToArrowBytes(bulkData, fixtures.TestAllocator)
	assert.NoError(b, err)
	tf, err := NewTestFixture(
		b.Context(),
		&fixtures.MockServerConfig{
			QueryBulkResponse: &commonv1.OnlineQueryBulkResponse{
				ScalarsData: bytes,
			},
		})
	assert.NoError(b, err)
	assertOnce := sync.Once{}
	return func() {
		req := chalk.OnlineQueryParams{}.
			WithInput("user.id", []int{1}).
			WithOutputs("user.id", "user.socure_score")
		res, err := tf.Client.OnlineQueryBulk(b.Context(), req)
		assert.NoError(b, err)
		var results []root
		assert.NoError(b, res.UnmarshalInto(&results))
		assert.Len(b, results, 1)
		firstRes := results[0]
		assertOnce.Do(func() {
			assert.Equal(b, int64(122), *firstRes.WindowedIntFeatures.Int1["1m"])
			assert.Equal(b, int64(122), *firstRes.WindowedIntFeatures.Int13["1h"])
			assert.Equal(b, float64(1.234), *firstRes.WindowedFloatFeatures.Float1["1m"])
			assert.Equal(b, float64(1.234), *firstRes.WindowedFloatFeatures.Float13["1h"])
			assert.Equal(b, "string_val", *firstRes.WindowedStringFeatures.String1["1m"])
			assert.Equal(b, "string_val", *firstRes.WindowedStringFeatures.String13["1h"])
			assert.True(b, *firstRes.WindowedBoolFeatures.Bool1["1m"])
			assert.True(b, *firstRes.WindowedBoolFeatures.Bool13["1h"])
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *firstRes.WindowedTimestampFeatures.Timestamp1["1m"])
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *firstRes.WindowedTimestampFeatures.Timestamp13["1h"])
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
