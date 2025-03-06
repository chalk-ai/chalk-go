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

func getBenchmarkQueryBulkLoneMultiNsPrimitives(b *testing.B) (benchFunc func(), closeFunc func()) {
	type root struct {
		StringFeatures    fixtures.StringFeatures
		IntFeatures       fixtures.IntFeatures
		FloatFeatures     fixtures.FloatFeatures
		BoolFeatures      fixtures.BoolFeatures
		TimestampFeatures fixtures.TimestampFeatures
	}

	bulkData := make(map[string]any)
	for j := 1; j <= 40; j++ {
		bulkData[fmt.Sprintf("string_features.string_%d", j)] = []string{fmt.Sprintf("string_val_%d", j)}
		bulkData[fmt.Sprintf("int_features.int_%d", j)] = []int64{int64(j)}
		bulkData[fmt.Sprintf("float_features.float_%d", j)] = []float64{float64(j)}
		bulkData[fmt.Sprintf("bool_features.bool_%d", j)] = []bool{j%2 == 0}
		bulkData[fmt.Sprintf("timestamp_features.timestamp_%d", j)] = []time.Time{time.Date(2021, 1, 1, 0, j, 0, 0, time.UTC)}
	}

	record, err := internal.ColumnMapToRecord(bulkData)
	assert.NoError(b, err)
	bytes, err := internal.RecordToBytes(record)
	tf, err := fixtures.NewTestFixture(&fixtures.MockServerConfig{
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
		res, err := tf.Client.OnlineQuery(context.Background(), req, nil)
		assert.NoError(b, err)
		var results root
		assert.NoError(b, res.UnmarshalInto(&results))
		assertOnce.Do(func() {
			assert.Equal(b, "string_val_1", *results.StringFeatures.String1)
			assert.Equal(b, "string_val_40", *results.StringFeatures.String40)
			assert.Equal(b, int64(1), *results.IntFeatures.Int1)
			assert.Equal(b, int64(40), *results.IntFeatures.Int40)
			assert.Equal(b, float64(1), *results.FloatFeatures.Float1)
			assert.Equal(b, float64(40), *results.FloatFeatures.Float40)
			assert.Equal(b, false, *results.BoolFeatures.Bool1)
			assert.Equal(b, true, *results.BoolFeatures.Bool40)
			assert.Equal(b, time.Date(2021, 1, 1, 0, 1, 0, 0, time.UTC), *results.TimestampFeatures.Timestamp1)
			assert.Equal(b, time.Date(2021, 1, 1, 0, 40, 0, 0, time.UTC), *results.TimestampFeatures.Timestamp40)
		})
	}, tf.Close
}

/*
 * Query: Bulk (single row)
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkQueryBulkLoneMultiNsPrimitives(b *testing.B) {
	benchFunc, closeFunc := getBenchmarkQueryBulkLoneMultiNsPrimitives(b)
	benchmarkWithClose(b, benchFunc, closeFunc)
}

/*
 * Query: Bulk (single row)
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkQueryBulkLoneMultiNsPrimitivesParallel(b *testing.B) {
	benchFunc, closeFunc := getBenchmarkQueryBulkLoneMultiNsPrimitives(b)
	benchmarkParallelWithClose(b, benchFunc, closeFunc)
}
