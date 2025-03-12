package chalk

import (
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestOfflineQueryParamsAllTypes(t *testing.T) {
	// Tests that all types of input, output, and required output parameters can be passed
	// without error.
	initErr := InitFeatures(&fixtures.Root)
	assert.Nil(t, initErr)
	params := OfflineQueryParams{}.
		WithInput(fixtures.Root.AllTypes.String, []any{1}).
		WithOutputs(
			"all_types.string",
			fixtures.Root.AllTypes.Bool,
			fixtures.Root.AllTypes.Float,
			fixtures.Root.AllTypes.String,
			fixtures.Root.AllTypes.Int,
			fixtures.Root.AllTypes.Timestamp,
			fixtures.Root.AllTypes.IntList,
		).
		WithRequiredOutputs(
			fixtures.Root.AllTypes.WindowedInt,
			fixtures.Root.AllTypes.WindowedInt["1m"],
			fixtures.Root.AllTypes.WindowedInt["5m"],
			fixtures.Root.AllTypes.WindowedInt["1h"],
			fixtures.Root.AllTypes.WindowedList,
			fixtures.Root.AllTypes.WindowedList["1m"],
			fixtures.Root.AllTypes.Nested,
			fixtures.Root.AllTypes.Nested.Id,
			fixtures.Root.AllTypes.Dataclass,
			fixtures.Root.AllTypes.Dataclass.Lat,
			fixtures.Root.AllTypes.Dataclass.Lng,
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
