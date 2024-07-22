package chalk

import (
	"fmt"
	"reflect"
)

func (p *UploadFeaturesParams) getConvertedInputsMap() (map[string]any, error) {
	inputs := p.Inputs
	res := make(map[string]any)

	allLength := -1
	for k, v := range inputs {
		var fqn string
		if _, ok := k.(string); ok {
			fqn = k.(string)
		} else {
			feature, err := UnwrapFeature(k)
			if err != nil {
				msg := fmt.Sprintf("Invalid inputs key '%v' with type '%T'. Expected `string` or `Feature`", k, k)
				return nil, &ErrorResponse{ClientError: &ClientError{Message: msg}}
			}
			fqn = feature.Fqn
		}
		res[fqn] = v

		currLength := -1
		if reflect.TypeOf(v).Kind() == reflect.Slice || reflect.TypeOf(v).Kind() == reflect.Array {
			currLength = reflect.ValueOf(v).Len()
		} else {
			return nil, &ErrorResponse{
				ClientError: &ClientError{
					Message: fmt.Sprintf("Values for feature '%s' must be a slice or array", fqn),
				},
			}
		}

		if allLength == -1 {
			allLength = currLength
		}
		if allLength != currLength {
			err := &ClientError{
				Message: fmt.Sprintf("All input slices or arrays must be the same length - found length %d for feature '%s' but expected length %d", currLength, fqn, allLength),
			}
			return nil, &ErrorResponse{ClientError: err}
		}
		if currLength == 0 {
			err := &ClientError{
				Message: fmt.Sprintf("All input slices or arrays must be non-empty - found length %d for feature '%s'", currLength, fqn),
			}
			return nil, &ErrorResponse{ClientError: err}
		}
	}
	return res, nil
}
