package errors

import (
	"errors"
)

// ErrNotFound in source.
var ErrNotFound = errors.New("not found")

// IsNotFound in source.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}
