package chalk

import (
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

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
