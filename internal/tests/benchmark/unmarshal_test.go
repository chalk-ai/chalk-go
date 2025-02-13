package benchmark

import (
	"context"
	"fmt"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal/tests/benchmark/fixtures"
	assert "github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
	"sync"
	"testing"
	"time"
)

func benchmarkRateLimited(b *testing.B, benchmarkFunc func(), qps int) {
	b.Helper()
	limiter := rate.NewLimiter(rate.Limit(qps), 1)
	var durations []time.Duration
	var mu sync.Mutex

	start := time.Now()
	duration := 1 * time.Second
	var wg sync.WaitGroup

	for time.Since(start) < duration {
		limiter.Wait(context.Background())

		wg.Add(1)
		go func() {
			defer wg.Done()
			opStart := time.Now()
			benchmarkFunc()
			opDuration := time.Since(opStart)
			mu.Lock()
			durations = append(durations, opDuration)
			mu.Unlock()
		}()
	}

	wg.Wait()

	p95 := PercentileDuration(durations, 95)
	b.ReportMetric(float64(p95.Nanoseconds()), "ns/op")
	b.ReportMetric(DurationMs(p95), "ms/op(p95)")
	b.ReportMetric(DurationMs(PercentileDuration(durations, 50)), "ms/op(median)")
	b.ReportMetric(DurationMs(PercentileDuration(durations, 100)), "ms/op(max)")
}

func BenchmarkUnmarshalMultiNsPrimitives(b *testing.B) {
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
		intFeatures := fixtures.IntFeatures{}
		floatFeatures := fixtures.FloatFeatures{}
		boolFeatures := fixtures.BoolFeatures{}
		stringFeatures := fixtures.StringFeatures{}
		timestampFeatures := fixtures.TimestampFeatures{}

		err := res.UnmarshalInto(&intFeatures)
		assert.Equal(b, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&floatFeatures)
		assert.Equal(b, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&boolFeatures)
		assert.Equal(b, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&stringFeatures)
		assert.Equal(b, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&timestampFeatures)
		assert.Equal(b, (*chalk.ClientError)(nil), err)

		assertOnce.Do(func() {
			assert.Equal(b, int64(122.0), *intFeatures.Int1)
			assert.Equal(b, int64(122.0), *intFeatures.Int40)
			assert.Equal(b, float64(1.234), *floatFeatures.Float1)
			assert.Equal(b, float64(1.234), *floatFeatures.Float40)
			assert.Equal(b, "string_val", *stringFeatures.String1)
			assert.Equal(b, "string_val", *stringFeatures.String40)
			assert.True(b, *boolFeatures.Bool1)
			assert.True(b, *boolFeatures.Bool40)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures.Timestamp1)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures.Timestamp40)
		})
	}

	benchmarkRateLimited(b, benchFunc, 5000)
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Windowed
 * Protocol: REST
 */
func BenchmarkUnmarshalMultiNsWindowed(t *testing.B) {
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

	benchmarkRateLimited(t, benchmarkFunc, 5000)
}
