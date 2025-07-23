package config

import (
	"fmt"
	"strings"
)

type SourcedValueT interface {
	~int | ~int64 | ~float32 | ~float64 | ~string
}

// SourcedConfig represents a configuration value along with information about where it came from
type SourcedConfig[T SourcedValueT] struct {
	Value  T
	Source string
}

// NewFromFlag creates a SourcedConfig from a command line flag value.
// If the flag doesn't start with "--", it will be prefixed with "--".
func NewFromFlag[T SourcedValueT](flag string, value T) *SourcedConfig[T] {
	if !strings.HasPrefix(flag, "--") {
		flag = "--" + flag
	}
	return &SourcedConfig[T]{
		Value:  value,
		Source: fmt.Sprintf("command line flag '%s'", flag),
	}
}

// NewFromEnvVar creates a SourcedConfig from an environment variable
func NewFromEnvVar[T SourcedValueT](key string, value T) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  value,
		Source: fmt.Sprintf("environment variable '%s'", key),
	}
}

// NewFromFile creates a SourcedConfig from a config file
func NewFromFile[T SourcedValueT](path string, value T) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  value,
		Source: fmt.Sprintf("config file %s", path),
	}
}

// NewFromArg creates a SourcedConfig from a direct argument
func NewFromArg[T SourcedValueT](value T) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  value,
		Source: "NewClient argument",
	}
}

// GetFirstNonEmpty returns the first non-empty SourcedConfig from the provided list
// using the provided zero-check function
func GetFirstNonEmpty[T SourcedValueT](configs ...*SourcedConfig[T]) *SourcedConfig[T] {
	var empty T
	for _, config := range configs {
		if config.Value != empty {
			return config
		}
	}
	var zero T
	return &SourcedConfig[T]{
		Value:  zero,
		Source: "value in '~/.chalk.yml' does not exist or is empty",
	}
}
