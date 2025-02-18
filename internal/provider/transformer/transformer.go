package transformer

import (
	"context"
)

type (
	// Transformer is an interface for transforming values.
	Transformer interface {
		// Transform transforms a value.
		Transform(ctx context.Context, value string) (string, error)

		// IsMissing returns true if the error is a missing error.
		IsMissing(err error) bool
	}

	// Transformers is a map of transformers.
	Transformers map[string]Transformer
)
