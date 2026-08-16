package envfs

import (
	"io/fs"
	"path/filepath"
)

// DirectoryOverrideEnvironmentGetter wraps another EnvironmentGetter but overrides the working directory
type DirectoryOverrideEnvironmentGetter struct {
	EnvironmentGetter
	base        EnvironmentGetter
	overrideDir string
}

// NewDirectoryOverrideEnvironmentGetter creates a new getter with an overridden working directory
func NewDirectoryOverrideEnvironmentGetter(base EnvironmentGetter, dir string) *DirectoryOverrideEnvironmentGetter {
	return &DirectoryOverrideEnvironmentGetter{
		base:        base,
		overrideDir: dir,
	}
}

func (d *DirectoryOverrideEnvironmentGetter) Getenv(key string) string {
	return d.base.Getenv(key)
}

func (d *DirectoryOverrideEnvironmentGetter) Getwd() (string, error) {
	if d.overrideDir != "" {
		// Convert to absolute path if needed
		absPath, err := filepath.Abs(d.overrideDir)
		if err != nil {
			return "", err
		}
		return absPath, nil
	}
	return d.base.Getwd()
}

func (d *DirectoryOverrideEnvironmentGetter) Abs(path string) (string, error) {
	return d.base.Abs(path)
}

func (d *DirectoryOverrideEnvironmentGetter) Stat(name string) (fs.FileInfo, error) {
	return d.base.Stat(name)
}

func (d *DirectoryOverrideEnvironmentGetter) ReadFile(name string) ([]byte, error) {
	return d.base.ReadFile(name)
}

func (d *DirectoryOverrideEnvironmentGetter) Open(name string) (File, error) {
	return d.base.Open(name)
}

func (d *DirectoryOverrideEnvironmentGetter) CreateTemp(dir string, pattern string) (WritableFile, error) {
	return d.base.CreateTemp(dir, pattern)
}

func (d *DirectoryOverrideEnvironmentGetter) Remove(name string) error {
	return d.base.Remove(name)
}
