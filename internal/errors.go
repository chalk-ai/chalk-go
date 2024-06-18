package internal

import (
	"fmt"
	"reflect"
)

func KindMismatchError(expected reflect.Kind, actual reflect.Kind) error {
	return fmt.Errorf("expected value of kind '%s', got '%s'", expected, actual)
}
