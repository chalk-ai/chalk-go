package chalk

import (
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

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
