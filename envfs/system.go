package envfs

import (
	"io/fs"
	"os"
)

// SystemEnvironmentGetter provides direct access to the system's environment
// and filesystem operations.
type SystemEnvironmentGetter struct {
	EnvironmentGetter
}

func (s SystemEnvironmentGetter) Getenv(key string) string {
	return os.Getenv(key)
}

func (s SystemEnvironmentGetter) Getwd() (string, error) {
	return os.Getwd()
}

func (s SystemEnvironmentGetter) Abs(path string) (string, error) {
	return defaultAbs(s, path)
}

func (s SystemEnvironmentGetter) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (s SystemEnvironmentGetter) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// defaultSystemEnvironmentGetter is the default instance used when no custom
// environment getter is provided via context.
var defaultSystemEnvironmentGetter = &SystemEnvironmentGetter{}