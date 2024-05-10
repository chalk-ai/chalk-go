package chalk

import (
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestOfflineQueryParamsAllTypes(t *testing.T) {
	// Tests that all types of input, output, and required output parameters can be passed
	// without error.
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	params := OfflineQueryParams{}.
		WithInput(testRootFeatures.AllTypes.String, []any{1}).
		WithOutputs(
			"all_types.string",
			testRootFeatures.AllTypes.Bool,
			testRootFeatures.AllTypes.Float,
			testRootFeatures.AllTypes.String,
			testRootFeatures.AllTypes.Int,
			testRootFeatures.AllTypes.Timestamp,
			testRootFeatures.AllTypes.IntList,
		).
		WithRequiredOutputs(
			testRootFeatures.AllTypes.WindowedInt,
			testRootFeatures.AllTypes.WindowedInt["1m"],
			testRootFeatures.AllTypes.WindowedInt["5m"],
			testRootFeatures.AllTypes.WindowedInt["1h"],
			testRootFeatures.AllTypes.WindowedList,
			testRootFeatures.AllTypes.WindowedList["1m"],
			testRootFeatures.AllTypes.Nested,
			testRootFeatures.AllTypes.Nested.Id,
			testRootFeatures.AllTypes.Dataclass,
			testRootFeatures.AllTypes.Dataclass.Lat,
			testRootFeatures.AllTypes.Dataclass.Lng,
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
