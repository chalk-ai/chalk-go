package utils

import (
	"errors"
	"fmt"
)

func GenerateGetEnumFunction[K comparable](valueToEnum map[string]K, enumName string) func(string) (*K, error) {
	return func(value string) (*K, error) {
		enum, found := valueToEnum[value]
		if !found {
			return nil, errors.New(fmt.Sprintf("Cannot find enum value '%s' among all %s", value, enumName))
		}

		return &enum, nil
	}
}
