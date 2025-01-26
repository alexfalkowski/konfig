package errors

import (
	"errors"
)

var (
	// ErrNotFound in source.
	ErrNotFound = errors.New("not found")

	// ErrInvalidFolder in source.
	ErrInvalidFolder = errors.New("invalid folder")
)

// IsNotFound in source.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}
