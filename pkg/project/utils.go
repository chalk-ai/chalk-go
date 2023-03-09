package project

import "os"

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
