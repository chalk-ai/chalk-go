package benchmark

import (
	"fmt"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal/tests/benchmark/fixtures"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func BenchmarkMultiNsDeserPrimitives(t *testing.B) {
	data := []chalk.FeatureResult{}
	for i := 1; i <= 40; i++ {
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("int_features.int%d", i),
			Value: float64(122.0),
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("float_features.float%d", i),
			Value: float64(1.234),
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("bool_features.bool%d", i),
			Value: true,
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("string_features.string%d", i),
			Value: "string_val",
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("timestamp_features.timestamp%d", i),
			Value: "2024-05-09T22:29:00Z",
		})
	}
	res := chalk.OnlineQueryResult{
		Data: data,
	}
	for i := 0; i < t.N; i++ {
		err := res.UnmarshalInto(&fixtures.IntFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&fixtures.FloatFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&fixtures.BoolFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&fixtures.StringFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&fixtures.TimestampFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
	}
}

func BenchmarkMultiNsDeserWindowed(t *testing.B) {
	newData := []chalk.FeatureResult{}
	windows := []int{60, 300, 3600}
	for i := 1; i <= 13; i++ {
		for _, window := range windows {
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_int_features.int%d__%d__", i, window),
				Value: float64(122.0),
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_float_features.float%d__%d__", i, window),
				Value: float64(1.234),
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_bool_features.bool%d__%d__", i, window),
				Value: true,
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_string_features.string%d__%d__", i, window),
				Value: "string_val",
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_timestamp_features.timestamp%d__%d__", i, window),
				Value: "2024-05-09T22:29:00Z",
			})
		}
	}
	res := chalk.OnlineQueryResult{
		Data: newData,
	}
	for i := 0; i < t.N; i++ {
		err := res.UnmarshalInto(&fixtures.WindowedIntFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&fixtures.WindowedFloatFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&fixtures.WindowedBoolFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&fixtures.WindowedStringFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
		err = res.UnmarshalInto(&fixtures.WindowedTimestampFeatures{})
		assert.Equal(t, (*chalk.ClientError)(nil), err)
	}
}
