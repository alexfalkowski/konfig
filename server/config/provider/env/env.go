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
func (e *Transformer) Transform(ctx context.Context, value string) (string, error) {
	return os.Getenv(value), nil
}
