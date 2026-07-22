package chalk

import (
	"encoding/json"
	"testing"
	"time"

	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	assert "github.com/stretchr/testify/require"
)

func TestTimeDeltaMatchesChalkpyFormatting(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		delta    time.Duration
		expected string
	}{
		{name: "zero", delta: 0, expected: "delta:"},
		{name: "rounds sub-milliseconds down", delta: 499 * time.Microsecond, expected: "delta:"},
		{name: "rounds sub-milliseconds up", delta: 500 * time.Microsecond, expected: "delta:1ms"},
		{name: "compound", delta: 2*24*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second + 6*time.Millisecond, expected: "delta:2d3h4m5s6ms"},
		{name: "negative", delta: -(7*24*time.Hour + 90*time.Minute), expected: "delta:-7d1h30m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, *TimeDelta(tt.delta))
		})
	}
}

func TestTimeDeltaCanPopulateGRPCOfflineQueryBounds(t *testing.T) {
	t.Parallel()

	request := &commonv1.OfflineQueryRequest{
		ObservedAtLowerBound: TimeDelta(30 * 24 * time.Hour),
		ObservedAtUpperBound: TimeDelta(24 * time.Hour),
		InsertedAtLowerBound: TimeDelta(12 * time.Hour),
		InsertedAtUpperBound: TimeDelta(time.Hour),
	}

	assert.Equal(t, "delta:30d", request.GetObservedAtLowerBound())
	assert.Equal(t, "delta:1d", request.GetObservedAtUpperBound())
	assert.Equal(t, "delta:12h", request.GetInsertedAtLowerBound())
	assert.Equal(t, "delta:1h", request.GetInsertedAtUpperBound())
}

func TestHTTPOfflineQuerySerializesDurationBounds(t *testing.T) {
	t.Parallel()

	observedLower := 30 * 24 * time.Hour
	observedUpper := 24 * time.Hour
	insertedLower := 12 * time.Hour
	insertedUpper := time.Hour
	params := &OfflineQueryParams{
		ObservedAtLowerBoundDuration: &observedLower,
		ObservedAtUpperBoundDuration: &observedUpper,
		InsertedAtLowerBoundDuration: &insertedLower,
		InsertedAtUpperBoundDuration: &insertedUpper,
	}
	resolved := &offlineQueryParamsResolved{
		inputs:          make(map[string][]TsFeatureValue),
		outputs:         []string{"test.feature"},
		requiredOutputs: []string{},
	}

	serialized, err := serializeOfflineQueryParams(params, resolved)
	assert.NoError(t, err)

	var payload struct {
		ObservedAtLowerBound *string `json:"observed_at_lower_bound"`
		ObservedAtUpperBound *string `json:"observed_at_upper_bound"`
		InsertedAtLowerBound *string `json:"inserted_at_lower_bound"`
		InsertedAtUpperBound *string `json:"inserted_at_upper_bound"`
	}
	assert.NoError(t, json.Unmarshal(serialized, &payload))
	assert.Equal(t, "delta:30d", *payload.ObservedAtLowerBound)
	assert.Equal(t, "delta:1d", *payload.ObservedAtUpperBound)
	assert.Equal(t, "delta:12h", *payload.InsertedAtLowerBound)
	assert.Equal(t, "delta:1h", *payload.InsertedAtUpperBound)
}

func TestHTTPOfflineQueryRejectsConflictingBounds(t *testing.T) {
	t.Parallel()

	timestamp := time.Now()
	delta := -time.Hour
	params := &OfflineQueryParams{
		ObservedAtLowerBound:         &timestamp,
		ObservedAtLowerBoundDuration: &delta,
	}
	resolved := &offlineQueryParamsResolved{
		inputs:          make(map[string][]TsFeatureValue),
		outputs:         []string{"test.feature"},
		requiredOutputs: []string{},
	}

	_, err := serializeOfflineQueryParams(params, resolved)
	assert.ErrorContains(t, err, "absolute and relative bounds cannot both be set")
}
