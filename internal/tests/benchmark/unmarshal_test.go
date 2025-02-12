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

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 */
func BenchmarkUnmarshalMultiNsPrimitivesNew(t *testing.B) {
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

	t.ResetTimer()
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			intFeatures := fixtures.IntFeatures{}
			floatFeatures := fixtures.FloatFeatures{}
			boolFeatures := fixtures.BoolFeatures{}
			stringFeatures := fixtures.StringFeatures{}
			timestampFeatures := fixtures.TimestampFeatures{}
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
			assert.Equal(t, int64(122.0), *intFeatures.Int1)
			assert.Equal(t, int64(122.0), *intFeatures.Int40)
			assert.Equal(t, float64(1.234), *floatFeatures.Float1)
			assert.Equal(t, float64(1.234), *floatFeatures.Float40)
			assert.Equal(t, "string_val", *stringFeatures.String1)
			assert.Equal(t, "string_val", *stringFeatures.String40)
			assert.True(t, *boolFeatures.Bool1)
			assert.True(t, *boolFeatures.Bool40)
			assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures.Timestamp1)
			assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures.Timestamp40)
		}
	})
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 */
func BenchmarkUnmarshalMultiNsPrimitives(t *testing.B) {
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

	for i := 0; i < t.N; i++ {
		var wg sync.WaitGroup
		numConcurrentReqs := 200
		for j := 0; j < numConcurrentReqs; j++ {
			wg.Add(1)
			go func() {
				intFeatures := fixtures.IntFeatures{}
				floatFeatures := fixtures.FloatFeatures{}
				boolFeatures := fixtures.BoolFeatures{}
				stringFeatures := fixtures.StringFeatures{}
				timestampFeatures := fixtures.TimestampFeatures{}
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
				wg.Done()

				assert.Equal(t, int64(122.0), *intFeatures.Int1)
				assert.Equal(t, int64(122.0), *intFeatures.Int40)
				assert.Equal(t, float64(1.234), *floatFeatures.Float1)
				assert.Equal(t, float64(1.234), *floatFeatures.Float40)
				assert.Equal(t, "string_val", *stringFeatures.String1)
				assert.Equal(t, "string_val", *stringFeatures.String40)
				assert.True(t, *boolFeatures.Bool1)
				assert.True(t, *boolFeatures.Bool40)
				assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures.Timestamp1)
				assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *timestampFeatures.Timestamp40)
			}()
		}
		wg.Wait()
	}

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

	for i := 0; i < t.N; i++ {
		var wg sync.WaitGroup
		numConcurrentReqs := 200
		for j := 0; j < numConcurrentReqs; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
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
			}()
		}
		wg.Wait()
	}
}
