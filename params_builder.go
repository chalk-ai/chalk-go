package chalk

import (
	"fmt"
	"strings"
)

type BuilderErrorType int

const (
	BuilderErrorUnknown BuilderErrorType = iota
	BuilderErrorInvalidFeature
)

type ParamType string

const (
	ParamTypeUnknown    ParamType = "unknown param"
	ParamOutput         ParamType = "output"
	ParamInput          ParamType = "input"
	ParamStaleness      ParamType = "staleness"
	ParamRequiredOutput ParamType = "required output"
)

type BuilderError struct {
	Err        error
	Type       BuilderErrorType
	ParamType  ParamType
	FeatureArg any
}

func (e *BuilderError) Error() string {
	switch e.Type {
	case BuilderErrorInvalidFeature:
		return fmt.Errorf("error occurred while adding %s feature '%s': %w", e.ParamType, e.FeatureArg, e.Err).Error()
	}

	return fmt.Errorf("error occurred while adding %s: %w", e.ParamType, e.Err).Error()
}

type BuilderErrors []*BuilderError

func (e BuilderErrors) Error() string {
	errStrings := make([]string, len(e))
	for idx, err := range e {
		errStrings[idx] = err.Error()
	}
	return strings.Join(errStrings, "\n")
}
