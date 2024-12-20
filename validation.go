package chalk

import (
	"fmt"
	"reflect"
)

func isArray(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Array
}

func isScalar(v any) bool {
	return !isArray(v)
}

func (p OnlineQueryParamsComplete) validateLists() BuilderErrors {
	lengthList := -1
	for k, v := range p.underlying.inputs {
		if !isArray(v) {
			continue
		}
		v = v.([]any)
		if lengthList == -1 {
			lengthList = len(v.([]any))
		} else {
			if lengthList != len(v.([]any)) {
				return BuilderErrors{
					&BuilderError{
						Err:       fmt.Errorf("all lists must be the same length - found length %d for feature '%s' but expected length %d", len(v.([]any)), k, lengthList),
						Type:      InvalidRequest,
						ParamType: ParamInput,
						Feature:   nil,
						Value:     nil,
					},
				}
			}
		}
	}
	return BuilderErrors{}
}

func (p OnlineQueryParamsComplete) validateAllListsOrAllScalars() BuilderErrors {
	allScalars := true
	allLists := true
	for _, v := range p.underlying.inputs {
		allScalars = allScalars && isScalar(v)
		allLists = allLists && isArray(v)
	}

	if allScalars || allLists {
		return nil
	}

	return BuilderErrors{
		&BuilderError{
			Err:       fmt.Errorf("inputs must be all scalars or all lists"),
			Type:      InvalidRequest,
			ParamType: ParamInput,
			Feature:   nil,
			Value:     nil,
		},
	}
}

func (p OnlineQueryParamsComplete) validatePostBuild() BuilderErrors {
	errors := BuilderErrors{}
	errors = append(errors, p.validateLists()...)
	errors = append(errors, p.validateAllListsOrAllScalars()...)
	return errors
}

func isValidQueryContextValue(v any) bool {
	switch v.(type) {
	case string, float64, bool:
		return true
	default:
		return false
	}
}
