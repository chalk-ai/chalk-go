package internal

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func StringOrNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

// ParseBucketDuration parses a bucket duration string
// and returns the duration in seconds.
// The input string must be of the form "Nunit" where
// N is a positive integer and unit is one of "s", "m", "h", "d", or "w".
func ParseBucketDuration(durationStr string) (int, error) {
	unitMap := map[string]int{
		"s": 1,
		"m": 60,
		"h": 3600,
		"d": 86400,
		"w": 604800,
	}

	// Parse the input string
	re := regexp.MustCompile(`^(\d+)([smhdw])$`)
	matches := re.FindStringSubmatch(durationStr)
	if matches == nil {
		return 0, fmt.Errorf("invalid bucket duration string: %s", durationStr)
	}

	// Convert the duration to seconds
	duration, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	unit := matches[2]
	seconds := duration * unitMap[unit]
	return seconds, nil
}
