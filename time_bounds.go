package chalk

import (
	"strconv"
	"strings"
	"time"
)

const timeDeltaPrefix = "delta:"

// TimeDelta formats a relative offline-query time bound for use in a gRPC
// request. HTTP callers can use the With...BoundDuration methods instead.
//
// For example, TimeDelta(-7*24*time.Hour) returns a pointer to "delta:-7d".
func TimeDelta(delta time.Duration) *string {
	formatted := timeDeltaPrefix + formatChalkDuration(delta)
	return &formatted
}

// formatChalkDuration matches chalkpy's timedelta_to_duration wire format.
// Chalk duration strings have millisecond precision and use days rather than
// Go's potentially much larger hour counts.
func formatChalkDuration(delta time.Duration) string {
	negative := delta < 0
	var magnitude uint64
	if negative {
		// Avoid overflowing when delta is the minimum possible time.Duration.
		magnitude = uint64(-(delta + 1)) + 1
	} else {
		magnitude = uint64(delta)
	}

	// Python's timedelta formatter rounds to the nearest millisecond.
	totalMilliseconds := (magnitude + uint64(time.Millisecond)/2) / uint64(time.Millisecond)
	if totalMilliseconds == 0 {
		return ""
	}

	const (
		millisecondsPerSecond = uint64(1000)
		millisecondsPerMinute = 60 * millisecondsPerSecond
		millisecondsPerHour   = 60 * millisecondsPerMinute
		millisecondsPerDay    = 24 * millisecondsPerHour
	)

	units := []struct {
		suffix       string
		milliseconds uint64
	}{
		{suffix: "d", milliseconds: millisecondsPerDay},
		{suffix: "h", milliseconds: millisecondsPerHour},
		{suffix: "m", milliseconds: millisecondsPerMinute},
		{suffix: "s", milliseconds: millisecondsPerSecond},
		{suffix: "ms", milliseconds: 1},
	}

	var result strings.Builder
	if negative {
		result.WriteByte('-')
	}
	for _, unit := range units {
		value := totalMilliseconds / unit.milliseconds
		if value == 0 {
			continue
		}
		result.WriteString(strconv.FormatUint(value, 10))
		result.WriteString(unit.suffix)
		totalMilliseconds %= unit.milliseconds
	}
	return result.String()
}
