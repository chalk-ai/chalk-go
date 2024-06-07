package internal

import (
	"fmt"
	"strconv"
	"strings"
)

// Taken from https://github.com/DopplerHQ/cli/blob/bc700c85f773476a3bf646cd38e784b1d17e573e/pkg/version/version.go

// Version semver
type Version struct {
	Major int16
	Minor int16
	Patch int16
}

// Unwrap get the original error
func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// CompareVersions returns -1 if first is greater, 1 if second is greater, and 0 otherwise
func CompareVersions(a Version, b Version) int {
	// major version
	if a.Major > b.Major {
		return -1
	}
	if b.Major > a.Major {
		return 1
	}

	// minor version
	if a.Minor > b.Minor {
		return -1
	}
	if b.Minor > a.Minor {
		return 1
	}

	// patch version
	if a.Patch > b.Patch {
		return -1
	}
	if b.Patch > a.Patch {
		return 1
	}

	return 0
}

// ParseVersion from a string
func ParseVersion(s string) (Version, error) {
	if strings.HasPrefix(s, "v") {
		s = s[1:]
	}
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return Version{}, fmt.Errorf("Invalid version %s", s)
	}

	var v Version
	var major int64
	var minor int64
	var patch int64
	var err error
	if major, err = strconv.ParseInt(parts[0], 10, 16); err != nil {
		return Version{}, fmt.Errorf("Invalid version %s", s)
	}
	if minor, err = strconv.ParseInt(parts[1], 10, 16); err != nil {
		return Version{}, fmt.Errorf("Invalid version %s", s)
	}
	if patch, err = strconv.ParseInt(parts[2], 10, 16); err != nil {
		return Version{}, fmt.Errorf("Invalid version %s", s)
	}

	v.Major = int16(major)
	v.Minor = int16(minor)
	v.Patch = int16(patch)
	return v, nil
}
