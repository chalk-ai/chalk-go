package config

import (
	"context"
	"fmt"
)

type SourcedValueT interface {
	~int | ~int64 | ~float32 | ~float64 | ~string
}

type SourceKind string

var (
	DefaultSourceKind = SourceKind("default")
	EnvSourceKind     = SourceKind("env")
	FileSourceKind    = SourceKind("file")
	ArgSourceKind     = SourceKind("arg")
	TokenSourceKind   = SourceKind("service")
	EmptySourceKind   = SourceKind("empty")
)

// SourcedConfig represents a configuration value along with information about where it came from
type SourcedConfig[T SourcedValueT] struct {
	Value  T
	Source string
	Kind   SourceKind
}

func (c *SourcedConfig[T]) WithValue(value T) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  value,
		Source: c.Source,
		Kind:   c.Kind,
	}
}

func (c *SourcedConfig[T]) WithSource(source string) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  c.Value,
		Source: source,
		Kind:   c.Kind,
	}
}

func (c *SourcedConfig[T]) WithSourceF(source string, args ...any) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  c.Value,
		Source: fmt.Sprintf(source, args...),
		Kind:   c.Kind,
	}
}

func NewFromToken[T SourcedValueT](value T, description string) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  value,
		Source: description,
		Kind:   TokenSourceKind,
	}
}

// NewFromFile creates a SourcedConfig from a config file
func NewFromFile[T SourcedValueT](path string, value T) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  value,
		Source: fmt.Sprintf("config file %s", path),
		Kind:   FileSourceKind,
	}
}

func NewFromDefault[T SourcedValueT](value T, desc string) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  value,
		Source: desc,
		Kind:   DefaultSourceKind,
	}
}

// NewFromArg creates a SourcedConfig from a direct argument
func NewFromArg[T SourcedValueT](value T) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  value,
		Source: "explicit argument",
		Kind:   ArgSourceKind,
	}
}

func NewFromEnvVar[T ~string](ctx context.Context, key string) *SourcedConfig[T] {
	return &SourcedConfig[T]{
		Value:  T(EnvironmentGetterFromContext(ctx).Getenv(key)),
		Source: fmt.Sprintf("environment variable '%s'", key),
		Kind:   EnvSourceKind,
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
	return nil
}
