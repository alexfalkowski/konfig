package env

import (
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
func (e *Transformer) Transform(value string) (string, error) {
	return os.Getenv(value), nil
}
