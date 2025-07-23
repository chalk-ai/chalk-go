package config

import (
	"context"
	"os"
	"path/filepath"
)

// EnvironmentGetter is an interface for getting environment variables
// and the current working directory. This can help with testing by making
// otherwise impure functions pure.
type EnvironmentGetter interface {
	Getenv(key string) string
	Getwd() (string, error)
	Abs(path string) (string, error)
}

type DefaultEnvironmentGetter struct {
	EnvironmentGetter
}

func (d DefaultEnvironmentGetter) Getenv(key string) string {
	return os.Getenv(key)
}

func (d DefaultEnvironmentGetter) Getwd() (string, error) {
	return os.Getwd()
}

func (d DefaultEnvironmentGetter) Abs(path string) (string, error) {
	return defaultAbs(d, path)
}

func defaultAbs(getter EnvironmentGetter, path string) (string, error) {
	wd, err := getter.Getwd()
	if err != nil {
		return "", err
	}
	if filepath.IsAbs(path) {
		return path, nil
	}
	return filepath.Join(wd, path), nil
}

var defaultEnvironmentGetter = &DefaultEnvironmentGetter{}

type MapEnvironmentGetter struct {
	EnvironmentGetter
	Env map[string]string
	Wd  string
}

func (m MapEnvironmentGetter) Getenv(key string) string {
	return m.Env[key]
}

func (m MapEnvironmentGetter) Getwd() (string, error) {
	return m.Wd, nil
}

func (m MapEnvironmentGetter) Abs(path string) (string, error) {
	return defaultAbs(m, path)
}

func NewMockEnvironmentGetter(
	env map[string]string,
	wd string,
) *MapEnvironmentGetter {
	return &MapEnvironmentGetter{Env: env, Wd: wd}
}

// DirectoryOverrideEnvironmentGetter wraps another EnvironmentGetter but overrides the working directory
type DirectoryOverrideEnvironmentGetter struct {
	EnvironmentGetter
	base        EnvironmentGetter
	overrideDir string
}

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
	return defaultAbs(d, path)
}

type customEnvironmentContextKey struct{}

func ContextWithEnvironmentGetter(ctx context.Context, getter EnvironmentGetter) context.Context {
	return context.WithValue(ctx, customEnvironmentContextKey{}, getter)
}

func EnvironmentGetterFromContext(ctx context.Context) EnvironmentGetter {
	if getter, ok := ctx.Value(customEnvironmentContextKey{}).(EnvironmentGetter); ok {
		return getter
	}

	return defaultEnvironmentGetter
}
