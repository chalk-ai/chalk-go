package internal

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var NameTag = "name"
var WindowsTag = "windows"

var IsoFormat = "2006-01-02T15:04:05.000000-07:00"

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

// FormatBucketDuration takes an integer number of seconds and
// returns a string representation that satisfies these conditions:
//   - the largest possible unit of time (e.g. "10m" instead of "600s")
//   - a single unit of time (e.g. "601s" instead of "10m1s")
func FormatBucketDuration(duration int) string {
	units := []string{"s", "m", "h", "d"}
	divisors := []int{60, 60, 24, 7}

	for i, unit := range units {
		div := divisors[i]
		if duration%div != 0 {
			return fmt.Sprintf("%d%s", duration, unit)
		}
		duration = duration / div
	}

	return fmt.Sprintf("%dw", duration)
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

func ExpandTilde(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return strings.Replace(path, "~", homeDir, 1), nil
}

func getFeatureNameFromFqn(fqn string) string {
	lastPart := strings.Split(fqn, ".")
	return lastPart[len(lastPart)-1]
}

// ChalkpySnakeCase aims to be in parity with
// our Python implementation of snake_case
func ChalkpySnakeCase(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isASCIIUpper(c) {
			if i > 0 && s[i-1] != '.' {
				b = append(b, '_')
			}
			c += 'a' - 'A'
		} else if isASCIIDigit(c) && i > 0 && isASCIILower(s[i-1]) {
			b = append(b, '_')
		}
		b = append(b, c)
	}
	return string(b)
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func ResolveFeatureName(field reflect.StructField) (string, error) {
	if tag := field.Tag.Get(NameTag); tag != "" {
		return tag, nil
	}
	versioned := field.Tag.Get("versioned")
	fieldName := ChalkpySnakeCase(field.Name)
	if versioned == "true" {
		parts := strings.Split(fieldName, "_")
		nameErr := fmt.Errorf(
			"versioned feature must have a version suffix `VN` at the"+
				" end of the attribute name, but found '%s' instead",
			fieldName,
		)
		if len(parts) == 1 {
			return "", nameErr
		}
		lastPart := parts[len(parts)-1]
		if !strings.HasPrefix(lastPart, "v") {
			return "", nameErr
		}
		version := lastPart[1:]
		prefix := strings.Join(parts[:len(parts)-1], "_")
		if version == "1" {
			fieldName = prefix
		} else {
			fieldName = prefix + "@" + version
		}
	} else if strings.HasPrefix(versioned, "default(") && strings.HasSuffix(versioned, ")") {
		version := versioned[len("default(") : len(versioned)-len(")")]
		_, convertErr := strconv.Atoi(version)
		if convertErr != nil {
			return "", fmt.Errorf(
				"expected struct tag `versioned:\"default(N)\"` "+
					"where N is an integer, but found %s instead",
				versioned,
			)
		}
		if version != "1" {
			fieldName = fieldName + "@" + version
		}
	} else if versioned != "" {
		return "", fmt.Errorf(
			"expected struct tag `versioned:\"true\"` or `versioned:\"default(N)\"` "+
				"where N is an integer, but found '%s' instead",
			versioned,
		)
	}
	return fieldName, nil
}

func GetWindowBucketsFromStructTag(field reflect.StructField) ([]string, error) {
	tag := field.Tag.Get(WindowsTag)
	tags := strings.Split(tag, ",")
	if tag == "" || len(tags) == 0 {
		return nil, errors.Newf("Window bucket struct tag missing or empty, e.g. `%s:\"1m,5m,...\"`", WindowsTag)
	}
	return tags, nil
}

func GetWindowBucketsSecondsFromStructTag(field reflect.StructField) ([]int, error) {
	buckets, err := GetWindowBucketsFromStructTag(field)
	if err != nil {
		return nil, err
	}

	seconds := make([]int, len(buckets))
	for i, bucket := range buckets {
		val, err := ParseBucketDuration(bucket)
		if err != nil {
			return nil, err
		}
		seconds[i] = val
	}
	return seconds, nil
}

func Ptr[T any](value T) *T {
	return &value
}

func allValid(l int) []bool {
	valid := make([]bool, l)
	for i := range valid {
		valid[i] = true
	}
	return valid
}

func GetBucketFromFqn(fqn string) (string, error) {
	sections := strings.Split(fqn, ".")
	lastSection := sections[len(sections)-1]
	lastSectionSplit := strings.Split(lastSection, "__")
	formatErr := fmt.Errorf(
		"error unmarshalling value for windowed bucket feature %s: "+
			"expected windowed bucket feature to have fqn of the format "+
			"`{fqn}__{bucket seconds}__` ",
		fqn,
	)
	if len(lastSectionSplit) < 2 {
		return "", formatErr
	}
	secondsStr := lastSectionSplit[1]
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return "", formatErr
	}
	return FormatBucketDuration(seconds), nil
}
