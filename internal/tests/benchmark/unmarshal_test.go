package benchmark

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/tests/benchmark/fixtures"
	assert "github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func benchmark(b *testing.B, benchmarkFunc func()) {
	b.Helper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkFunc()
	}
	b.StopTimer()

	avg := b.Elapsed() / time.Duration(b.N)
	b.ReportMetric(0, "ns/op")                                  // Effective hides the default ns/op metric
	b.ReportMetric((float64(avg.Nanoseconds()) / 1e6), "ms/op") // The same metric but in ms
}

func benchmarkParallel(b *testing.B, benchmarkFunc func()) {
	b.Helper()
	numParallel := 200
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < numParallel; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				benchmarkFunc()
			}()
		}
		wg.Wait()
	}
	b.StopTimer()
	avg := b.Elapsed() / time.Duration(b.N)
	b.ReportMetric(0, "ns/op")                                  // Effective hides the default ns/op metric
	b.ReportMetric((float64(avg.Nanoseconds()) / 1e6), "ms/op") // The same metric but in ms
}

func getBenchmarkBulkMultiNsPrimitives(b *testing.B) func() {
	bulkData := make(map[string]any)
	for i := 0; i < 100; i++ {
		for j := 1; j <= 40; j++ {
			fqn := fmt.Sprintf("int_features.int_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []float64{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]float64), float64(122.0))

			fqn = fmt.Sprintf("float_features.float_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []float64{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]float64), float64(1.234))

			fqn = fmt.Sprintf("bool_features.bool_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []bool{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]bool), true)

			fqn = fmt.Sprintf("string_features.string_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []string{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]string), fmt.Sprintf("string_val_%d", i))

			fqn = fmt.Sprintf("timestamp_features.timestamp_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []string{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]string), "2024-05-09T22:29:00Z")
		}
	}

	record, err := internal.ColumnMapToRecord(bulkData)
	assert.NoError(b, err)

	table := array.NewTableFromRecords(record.Schema(), []arrow.Record{record})

	res := chalk.OnlineQueryBulkResult{
		ScalarsTable: table,
	}

	assertOnce := sync.Once{}
	benchFunc := func() {
		intFeatures := []fixtures.IntFeatures{}
		floatFeatures := []fixtures.FloatFeatures{}
		boolFeatures := []fixtures.BoolFeatures{}
		stringFeatures := []fixtures.StringFeatures{}
		timestampFeatures := []fixtures.TimestampFeatures{}

		assert.NoError(b, res.UnmarshalInto(&intFeatures)))
		assert.NoError(b, res.UnmarshalInto(&floatFeatures))
		assert.NoError(b, res.UnmarshalInto(&boolFeatures))
		assert.NoError(b, res.UnmarshalInto(&stringFeatures))
		assert.NoError(b, res.UnmarshalInto(&timestampFeatures))

		assertOnce.Do(func() {
			for i := 0; i < 100; i++ {
				assert.Equal(b, int64(122.0), *intFeatures[i].Int1)
				assert.Equal(b, int64(122.0), *intFeatures[i].Int40)
				assert.Equal(b, float64(1.234), *floatFeatures[i].Float1)
				assert.Equal(b, float64(1.234), *floatFeatures[i].Float40)
				assert.Equal(b, fmt.Sprintf("string_val_%d", i), *stringFeatures[i].String1)
				assert.Equal(b, fmt.Sprintf("string_val_%d", i), *stringFeatures[i].String40)
				assert.True(b, *boolFeatures[i].Bool1)
				assert.True(b, *boolFeatures[i].Bool40)
				assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures[i].Timestamp1)
				assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures[i].Timestamp40)
			}
		})
	}

	return benchFunc
}

func getBenchmarkMultiNsPrimitives(b *testing.B) func() {
	data := []chalk.FeatureResult{}
	for i := 1; i <= 40; i++ {
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("int_features.int_%d", i),
			Value: float64(122.0),
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("float_features.float_%d", i),
			Value: float64(1.234),
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("bool_features.bool_%d", i),
			Value: true,
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("string_features.string_%d", i),
			Value: "string_val",
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("timestamp_features.timestamp_%d", i),
			Value: "2024-05-09T22:29:00Z",
		})
	}
	res := chalk.OnlineQueryResult{
		Data: data,
	}
	assertOnce := sync.Once{}
	benchFunc := func() {
		rootStruct := struct {
			IntFeatures       fixtures.IntFeatures
			FloatFeatures     fixtures.FloatFeatures
			BoolFeatures      fixtures.BoolFeatures
			StringFeatures    fixtures.StringFeatures
			TimestampFeatures fixtures.TimestampFeatures
		}{}
		assert.NoError(b, res.UnmarshalInto(&rootStruct))

		assertOnce.Do(func() {
			assert.Equal(b, int64(122.0), *rootStruct.IntFeatures.Int1)
			assert.Equal(b, int64(122.0), *rootStruct.IntFeatures.Int40)
			assert.Equal(b, float64(1.234), *rootStruct.FloatFeatures.Float1)
			assert.Equal(b, float64(1.234), *rootStruct.FloatFeatures.Float40)
			assert.Equal(b, "string_val", *rootStruct.StringFeatures.String1)
			assert.Equal(b, "string_val", *rootStruct.StringFeatures.String40)
			assert.True(b, *rootStruct.BoolFeatures.Bool1)
			assert.True(b, *rootStruct.BoolFeatures.Bool40)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct.TimestampFeatures.Timestamp1)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct.TimestampFeatures.Timestamp40)
		})
	}

	return benchFunc
}

func getBenchmarkUnmarshalMultiNsWindowed(t *testing.B) func() {
	newData := []chalk.FeatureResult{}
	windows := []int{60, 300, 3600}
	for i := 1; i <= 13; i++ {
		for _, window := range windows {
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_int_features.int_%d__%d__", i, window),
				Value: float64(122.0),
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_float_features.float_%d__%d__", i, window),
				Value: float64(1.234),
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_bool_features.bool_%d__%d__", i, window),
				Value: true,
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_string_features.string_%d__%d__", i, window),
				Value: "string_val",
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_timestamp_features.timestamp_%d__%d__", i, window),
				Value: "2024-05-09T22:29:00Z",
			})
		}
	}
	res := chalk.OnlineQueryResult{
		Data: newData,
	}

	assertOnce := sync.Once{}

	benchmarkFunc := func() {
		rootStruct := struct {
			IntFeatures       fixtures.WindowedIntFeatures
			FloatFeatures     fixtures.WindowedFloatFeatures
			BoolFeatures      fixtures.WindowedBoolFeatures
			StringFeatures    fixtures.WindowedStringFeatures
			TimestampFeatures fixtures.WindowedTimestampFeatures
		}{}
		assert.NoError(t, res.UnmarshalInto(&rootStruct))

		assertOnce.Do(func() {
			assert.Equal(t, int64(122.0), *rootStruct.IntFeatures.Int1["1m"])
			assert.Equal(t, int64(122.0), *rootStruct.IntFeatures.Int13["1h"])
			assert.Equal(t, float64(1.234), *rootStruct.FloatFeatures.Float1["1m"])
			assert.Equal(t, float64(1.234), *rootStruct.FloatFeatures.Float13["1h"])
			assert.Equal(t, "string_val", *rootStruct.StringFeatures.String1["1m"])
			assert.Equal(t, "string_val", *rootStruct.StringFeatures.String13["1h"])
			assert.True(t, *rootStruct.BoolFeatures.Bool1["1m"])
			assert.True(t, *rootStruct.BoolFeatures.Bool13["1h"])
			assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct.TimestampFeatures.Timestamp1["1m"])
			assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct.TimestampFeatures.Timestamp13["1h"])
		})
	}

	return benchmarkFunc
}

func getBenchmarkSingleNs(b *testing.B) func() {
	data := []chalk.FeatureResult{}
	for i := 1; i <= 40; i++ {
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("int_features.int_%d", i),
			Value: float64(122.0),
		})
	}
	res := chalk.OnlineQueryResult{
		Data: data,
	}
	assertOnce := sync.Once{}
	benchFunc := func() {
		intFeatures := fixtures.IntFeatures{}
		assert.NoError(b, res.UnmarshalInto(&intFeatures))

		assertOnce.Do(func() {
			assert.Equal(b, int64(122.0), *intFeatures.Int1)
			assert.Equal(b, int64(122.0), *intFeatures.Int40)
		})
	}

	return benchFunc
}

func getBenchmarkBulkSingleNs(b *testing.B) func() {
	bulkData := make(map[string]any)
	for i := 0; i < 100; i++ {
		for j := 1; j <= 40; j++ {
			fqn := fmt.Sprintf("string_features.string_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []string{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]string), fmt.Sprintf("string_val_%d_%d", i, j))
		}
	}

	record, err := internal.ColumnMapToRecord(bulkData)
	assert.NoError(b, err)

	table := array.NewTableFromRecords(record.Schema(), []arrow.Record{record})

	res := chalk.OnlineQueryBulkResult{
		ScalarsTable: table,
	}
	assertOnce := sync.Once{}
	benchFunc := func() {
		stringFeatures := []fixtures.StringFeatures{}
		assert.NoError(t, res.UnmarshalInto(&stringFeatures))

		assertOnce.Do(func() {
			for i := 0; i < 100; i++ {
				assert.Equal(b, fmt.Sprintf("string_val_%d_%d", i, 1), *stringFeatures[i].String1)
				assert.Equal(b, fmt.Sprintf("string_val_%d_%d", i, 40), *stringFeatures[i].String40)
			}
		})
	}

	return benchFunc
}

/*
 * Query: Single
 * Namespaces: Single
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalSingleNsPrimitivesSingle(b *testing.B) {
	benchmark(b, getBenchmarkSingleNs(b))
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Windowed
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalMultiNsWindowedSingle(t *testing.B) {
	benchmark(t, getBenchmarkUnmarshalMultiNsWindowed(t))
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Windowed
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalMultiNsWindowedParallel(t *testing.B) {
	benchmarkParallel(t, getBenchmarkUnmarshalMultiNsWindowed(t))
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalMultiNsPrimitivesSingle(b *testing.B) {
	benchmark(b, getBenchmarkMultiNsPrimitives(b))
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalMultiNsPrimitivesParallel(b *testing.B) {
	benchmarkParallel(b, getBenchmarkMultiNsPrimitives(b))
}

/*
 * Query: Bulk
 * Namespaces: Single
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalBulkSingleNsPrimitivesSingle(b *testing.B) {
	benchmark(b, getBenchmarkBulkSingleNs(b))
}

/*
 * Query: Bulk
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalBulkMultiNsPrimitivesSingle(b *testing.B) {
	benchmark(b, getBenchmarkBulkMultiNsPrimitives(b))
}

/*
 * Query: Bulk
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalBulkMultiNsPrimitivesParallel(b *testing.B) {
	benchmarkParallel(b, getBenchmarkBulkMultiNsPrimitives(b))
}
