package chalk

import (
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestOfflineQueryParamsAllTypes(t *testing.T) {
	// Tests that all types of input, output, and required output parameters can be passed
	// without error.
	initErr := InitFeatures(&ptf)
	assert.Nil(t, initErr)
	params := OfflineQueryParams{}.
		WithInput(ptf.AllTypes.String, []any{1}).
		WithOutputs(
			"all_types.string",
			ptf.AllTypes.Bool,
			ptf.AllTypes.Float,
			ptf.AllTypes.String,
			ptf.AllTypes.Int,
			ptf.AllTypes.Timestamp,
			ptf.AllTypes.IntList,
		).
		WithRequiredOutputs(
			ptf.AllTypes.WindowedInt,
			ptf.AllTypes.WindowedInt["1m"],
			ptf.AllTypes.WindowedInt["5m"],
			ptf.AllTypes.WindowedInt["1h"],
			ptf.AllTypes.WindowedList,
			ptf.AllTypes.WindowedList["1m"],
			ptf.AllTypes.Nested,
			ptf.AllTypes.Nested.Id,
			ptf.AllTypes.Dataclass,
			ptf.AllTypes.Dataclass.Lat,
			ptf.AllTypes.Dataclass.Lng,
		)
	assert.Empty(t, params.underlying.builderErrors)
}

func TestOfflineQueryInputParamInteger(t *testing.T) {
	// Tests passing an integer as input feature reference. Should fail.
	var invalidFeatureReference int
	params := OfflineQueryParams{}.WithInput(invalidFeatureReference, []any{1})
	assert.NotEmpty(t, params.underlying.builderErrors)
}

func TestOfflineQueryOutputParamInteger(t *testing.T) {
	// Tests passing an integer as output feature reference. Should fail.
	var invalidFeatureReference int
	params := OfflineQueryParams{}.WithOutputs(invalidFeatureReference)
	assert.NotEmpty(t, params.underlying.builderErrors)
}

func TestOfflineQueryRequiredOutputsParamInteger(t *testing.T) {
	// Tests passing an integer as staleness feature reference. Should fail.
	var invalidFeatureReference int
	underlying := OfflineQueryParams{}.withRequiredOutputs(invalidFeatureReference, time.Second*5)
	assert.NotEmpty(t, underlying.builderErrors)
}
