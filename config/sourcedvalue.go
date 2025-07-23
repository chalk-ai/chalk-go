package config

import (
	"context"
	"fmt"
)

type SourcedValueT interface {
	~int | ~int64 | ~float32 | ~float64 | ~string
}

// SourcedConfig represents a configuration value along with information about where it came from
type SourcedConfig[T SourcedValueT] struct {
	Value  T
	Source string
}

// NewFromFile creates a SourcedConfig from a config file
func NewFromFile[T SourcedValueT](path string, value T) SourcedConfig[T] {
	return SourcedConfig[T]{
		Value:  value,
		Source: fmt.Sprintf("config file %s", path),
	}
}

// NewFromArg creates a SourcedConfig from a direct argument
func NewFromArg[T SourcedValueT](value T) SourcedConfig[T] {
	return SourcedConfig[T]{
		Value:  value,
		Source: "NewClient argument",
	}
}

func NewFromEnvVar[T ~string](ctx context.Context, key string) SourcedConfig[T] {
	return SourcedConfig[T]{
		Value:  T(EnvironmentGetterFromContext(ctx).Getenv(key)),
		Source: fmt.Sprintf("environment variable '%s'", key),
	}
}

// GetFirstNonEmpty returns the first non-empty SourcedConfig from the provided list
// using the provided zero-check function
func GetFirstNonEmpty[T SourcedValueT](configs ...SourcedConfig[T]) SourcedConfig[T] {
	var empty T
	for _, config := range configs {
		if config.Value != empty {
			return config
		}
	}
	return SourcedConfig[T]{
		Value:  empty,
		Source: "empty value",
	}
}
