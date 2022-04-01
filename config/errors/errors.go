package errors

import (
	"errors"
)

var (
	// ErrUnmarshalError in config.
	ErrUnmarshalError = errors.New("unmarshal issue")

	// ErrMarshalError in config.
	ErrMarshalError = errors.New("marshal issue")
)
