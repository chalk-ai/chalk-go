package benchmark

import (
	"fmt"
	"github.com/chalk-ai/chalk-go"
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
		myStruct := struct {
			IntFeatures       fixtures.IntFeatures
			FloatFeatures     fixtures.FloatFeatures
			BoolFeatures      fixtures.BoolFeatures
			StringFeatures    fixtures.StringFeatures
			TimestampFeatures fixtures.TimestampFeatures
		}{}
		err := res.UnmarshalInto(&myStruct)
		assert.Equal(b, (*chalk.ClientError)(nil), err)

		assertOnce.Do(func() {
			assert.Equal(b, int64(122.0), *myStruct.IntFeatures.Int1)
			assert.Equal(b, int64(122.0), *myStruct.IntFeatures.Int40)
			assert.Equal(b, float64(1.234), *myStruct.FloatFeatures.Float1)
			assert.Equal(b, float64(1.234), *myStruct.FloatFeatures.Float40)
			assert.Equal(b, "string_val", *myStruct.StringFeatures.String1)
			assert.Equal(b, "string_val", *myStruct.StringFeatures.String40)
			assert.True(b, *myStruct.BoolFeatures.Bool1)
			assert.True(b, *myStruct.BoolFeatures.Bool40)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *myStruct.TimestampFeatures.Timestamp1)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *myStruct.TimestampFeatures.Timestamp40)
		})
	}

	return benchFunc
}

func getBenchmarkUnmarshalMultiNs(t *testing.B) func() {
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
		intFeatures := fixtures.WindowedIntFeatures{}
		floatFeatures := fixtures.WindowedFloatFeatures{}
		boolFeatures := fixtures.WindowedBoolFeatures{}
		stringFeatures := fixtures.WindowedStringFeatures{}
		timestampFeatures := fixtures.WindowedTimestampFeatures{}

		err := res.UnmarshalInto(&intFeatures)
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&floatFeatures)
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&boolFeatures)
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&stringFeatures)
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&timestampFeatures)
		assert.Equal(t, (*chalk.ClientError)(nil), err)

		assertOnce.Do(func() {
			assert.Equal(t, int64(122.0), *intFeatures.Int1["1m"])
			assert.Equal(t, int64(122.0), *intFeatures.Int13["1h"])
			assert.Equal(t, float64(1.234), *floatFeatures.Float1["1m"])
			assert.Equal(t, float64(1.234), *floatFeatures.Float13["1h"])
			assert.Equal(t, "string_val", *stringFeatures.String1["1m"])
			assert.Equal(t, "string_val", *stringFeatures.String13["1h"])
			assert.True(t, *boolFeatures.Bool1["1m"])
			assert.True(t, *boolFeatures.Bool13["1h"])
			assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures.Timestamp1["1m"])
			assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures.Timestamp13["1h"])
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
		err := res.UnmarshalInto(&intFeatures)
		assert.Equal(b, (*chalk.ClientError)(nil), err)

		assertOnce.Do(func() {
			assert.Equal(b, int64(122.0), *intFeatures.Int1)
			assert.Equal(b, int64(122.0), *intFeatures.Int40)
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
	benchmark(t, getBenchmarkUnmarshalMultiNs(t))
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Windowed
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalMultiNsWindowedParallel(t *testing.B) {
	benchmarkParallel(t, getBenchmarkUnmarshalMultiNs(t))
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
