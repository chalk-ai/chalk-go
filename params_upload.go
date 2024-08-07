package chalk

import (
	"github.com/cockroachdb/errors"
	"reflect"
)

func getConvertedInputsMap(inputs map[any]any) (map[string]any, error) {
	res := make(map[string]any)

	allLength := -1
	for k, v := range inputs {
		var fqn string
		if _, ok := k.(string); ok {
			fqn = k.(string)
		} else {
			feature, err := UnwrapFeature(k)
			if err != nil {
				return nil, errors.Newf(
					"Invalid inputs key '%v' with type '%T'. Expected `string` or `Feature`",
					k, k,
				)
			}
			fqn = feature.Fqn
		}
		res[fqn] = v

		currLength := -1
		if reflect.TypeOf(v).Kind() == reflect.Slice || reflect.TypeOf(v).Kind() == reflect.Array {
			currLength = reflect.ValueOf(v).Len()
		} else {
			return nil, errors.Newf("Values for feature '%s' must be a slice or array", fqn)
		}

		if allLength == -1 {
			allLength = currLength
		}
		if allLength != currLength {
			return nil, errors.Newf(
				"All input slices or arrays must be the same length - "+
					"found length %d for feature '%s' but expected length %d",
				currLength, fqn, allLength,
			)
		}
		if currLength == 0 {
			return nil, errors.Newf(
				"All input slices or arrays must be non-empty - found length %d for feature '%s'",
				currLength, fqn,
			)
		}
	}
	return res, nil
}
