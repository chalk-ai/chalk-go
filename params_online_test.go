package chalk

import (
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestOnlineQueryParamsAllTypes(t *testing.T) {
	// Tests that all types of input, output, and staleness parameters can be passed
	// without error.
	initErr := InitFeatures(&ptf)
	assert.Nil(t, initErr)
	params := OnlineQueryParams{}.
		WithInput(ptf.AllTypes.String, 1).
		WithOutputs(
			"all_types.string",
			ptf.AllTypes.Float,
			ptf.AllTypes.Int,
			ptf.AllTypes.Timestamp,
			ptf.AllTypes.IntList,
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
		).
		WithStaleness(
			ptf.AllTypes.Bool, time.Second*5,
		)
	assert.Empty(t, params.underlying.builderErrors)
}

func TestOnlineQueryInputParamInteger(t *testing.T) {
	// Tests passing an integer as input feature reference. Should fail.
	var invalidFeatureReference int
	params := OnlineQueryParams{}.WithInput(invalidFeatureReference, "1")
	assert.NotEmpty(t, params.underlying.builderErrors)
}

func TestOnlineQueryOutputParamInteger(t *testing.T) {
	// Tests passing an integer as output feature reference. Should fail.
	var invalidFeatureReference int
	params := OnlineQueryParams{}.WithOutputs(invalidFeatureReference)
	assert.NotEmpty(t, params.underlying.builderErrors)
}

func TestOnlineQueryStalenessParamInteger(t *testing.T) {
	// Tests passing an integer as staleness feature reference. Should fail.
	var invalidFeatureReference int
	underlying := OnlineQueryParams{}.WithStaleness(invalidFeatureReference, time.Second*5)
	assert.NotEmpty(t, underlying.builderErrors)
}
