package chalk

import (
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestOfflineQueryParamsAllTypes(t *testing.T) {
	t.Parallel()
	// Tests that all types of input, output, and required output parameters can be passed
	// without error.
	fixtureRoot, initErr := GetRootFeatures()
	assert.NoError(t, initErr)
	params := OfflineQueryParams{}.
		WithInput(fixtureRoot.AllTypes.String, []any{1}).
		WithOutputs(
			"all_types.string",
			fixtureRoot.AllTypes.Bool,
			fixtureRoot.AllTypes.Float,
			fixtureRoot.AllTypes.String,
			fixtureRoot.AllTypes.Int,
			fixtureRoot.AllTypes.Timestamp,
			fixtureRoot.AllTypes.IntList,
		).
		WithRequiredOutputs(
			fixtureRoot.AllTypes.WindowedInt,
			fixtureRoot.AllTypes.WindowedInt["1m"],
			fixtureRoot.AllTypes.WindowedInt["5m"],
			fixtureRoot.AllTypes.WindowedInt["1h"],
			fixtureRoot.AllTypes.WindowedList,
			fixtureRoot.AllTypes.WindowedList["1m"],
			fixtureRoot.AllTypes.Nested,
			fixtureRoot.AllTypes.Nested.Id,
			fixtureRoot.AllTypes.Dataclass,
			fixtureRoot.AllTypes.Dataclass.Lat,
			fixtureRoot.AllTypes.Dataclass.Lng,
		)
	_, err := params.underlying.resolve()
	assert.NoError(t, err)
}

func TestOfflineQueryInputParamInteger(t *testing.T) {
	t.Parallel()
	// Tests passing an integer as input feature reference. Should fail.
	var invalidFeatureReference int
	params := OfflineQueryParams{}.WithInput(invalidFeatureReference, []any{1})
	_, err := params.underlying.resolve()
	assert.Error(t, err)
}

func TestOfflineQueryOutputParamInteger(t *testing.T) {
	t.Parallel()
	// Tests passing an integer as output feature reference. Should fail.
	var invalidFeatureReference int
	params := OfflineQueryParams{}.WithOutputs(invalidFeatureReference)
	_, err := params.underlying.resolve()
	assert.Error(t, err)
}

func TestOfflineQueryRequiredOutputsParamInteger(t *testing.T) {
	t.Parallel()
	// Tests passing an integer as staleness feature reference. Should fail.
	var invalidFeatureReference int
	underlying := OfflineQueryParams{}.withRequiredOutputs(invalidFeatureReference, time.Second*5)
	_, err := underlying.resolve()
	assert.Error(t, err)
}
