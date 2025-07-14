package errors

import "github.com/cockroachdb/errors"

var ErrEmptyHasManyValues = errors.New("values of has-many results is empty")
