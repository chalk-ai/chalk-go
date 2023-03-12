package internal

import "os"

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
