package env

import (
	"context"
	"os"
)

// Transformer for env.
type Transformer struct {
}

// NewTransformer for env.
func NewTransformer() *Transformer {
	return &Transformer{}
}

// Transform for env.
func (e *Transformer) Transform(_ context.Context, value string) (any, error) {
	return os.Getenv(value), nil
}
