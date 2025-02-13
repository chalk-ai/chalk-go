package benchmark

import (
	"sort"
	"time"
)

// PercentileDuration calculates the given percentile from a slice of durations.
func PercentileDuration(durations []time.Duration, percentile int) time.Duration {
	if len(durations) == 0 {
		return 0 // Handle empty slice case
	}

	// Ensure percentile is in the valid range
	if percentile < 0 {
		percentile = 0
	} else if percentile > 100 {
		percentile = 100
	}

	// Sort durations in ascending order
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	// Calculate index
	index := int(float64(len(durations)-1) * float64(percentile) / 100.0)

	return durations[index]
}

// DurationMs converts a time.Duration to a float64 with 2 decimal places in milliseconds.
func DurationMs(d time.Duration) float64 {
	return float64(d.Nanoseconds()) / 1e6
}
