package errors

import (
	"errors"
)

// ErrNotFound in source.
var ErrNotFound = errors.New("not found")

// IsNotFoundError in source.
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}
