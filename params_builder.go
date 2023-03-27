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
	Err       error
	Type      BuilderErrorType
	ParamType ParamType
	Feature   any
	Value     any
}

func (e *BuilderError) Error() string {
	switch e.Type {
	case BuilderErrorInvalidFeature:
		err1 := fmt.Errorf("error occured while adding %s feature '%s' with value '%s': %w", e.ParamType, e.Feature, e.Value, e.Err)
		err2 := "Please make sure you are referencing a feature from the root 'Features' struct, for example: Features.MyFeatureClass.NestedFeatureClass.Id"
		err3 := "Please also make sure chalk.InitFeatures() has been called on the root 'Features' struct, and the global variable 'InitFeaturesErr' is nil"
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

func convertPanicToError(panicContents any, paramType ParamType, feature any, value any) *BuilderError {
	if panicContents == nil {
		return nil
	}

	errorStr := "unexpected param builder internal error"
	panicStr, ok := panicContents.(string)
	if ok {
		errorStr = panicStr
	}

	return &BuilderError{
		Err:       fmt.Errorf(errorStr),
		Type:      BuilderErrorUnknown,
		ParamType: paramType,
		Feature:   feature,
		Value:     value,
	}
}
