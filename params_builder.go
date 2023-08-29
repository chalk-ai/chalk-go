package chalk

import (
	"fmt"
	"reflect"
	"strings"
)

type BuilderErrorType int

const (
	BuilderErrorUnknown BuilderErrorType = iota
	InvalidFeatureType
	UnwrapFeatureError
	InvalidRequest
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
	Err       error
	Type      BuilderErrorType
	ParamType ParamType
	Feature   any
	Value     any
}

func (e *BuilderError) Error() string {
	switch e.Type {
	case UnwrapFeatureError:
		err1 := fmt.Errorf("error occured while adding %s feature '%s' with value '%s': %w", e.ParamType, e.Feature, e.Value, e.Err)
		err2 := "Please make sure you are referencing a feature by string or by passing a feature rooted in the 'Features' struct"
		err3 := "If using 'Features', please make sure to use InitFeatures and ensure InitFeaturesErr is nil"
		return strings.Join([]string{err1.Error(), err2, err3}, "\n")
	}

	return fmt.Errorf("error occurred while adding %s: %w", e.ParamType, e.Err).Error()
}

type BuilderErrors []*BuilderError

func (e BuilderErrors) Error() string {
	errStrings := make([]string, 0)
	for _, err := range e {
		if err == nil {
			continue
		}
		errStrings = append(errStrings, err.Error())
	}
	return strings.Join(errStrings, "\n")
}

func validateFeature(feature any, paramType ParamType) *BuilderError {
	value := reflect.ValueOf(feature)
	kind := value.Kind()
	if kind == reflect.String || kind == reflect.Ptr || kind == reflect.Map || kind == reflect.Slice || kind == reflect.Array {
		return nil
	}
	return &BuilderError{
		Err:       fmt.Errorf("expected string or pointer, but found invalid feature type: %s. If using 'Features', please make sure to use InitFeatures and ensure InitFeaturesErr is nil", kind),
		Type:      InvalidFeatureType,
		ParamType: paramType,
	}
}

func validateFeatures(features []any, paramType ParamType) *BuilderError {
	for _, feature := range features {
		if err := validateFeature(feature, paramType); err != nil {
			return err
		}
	}
	return nil
}
