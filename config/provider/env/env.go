package env

import (
	"os"
)

// Transformer for env.
type Transformer struct {
	value string
}

// NewTransformer for env.
func NewTransformer(value string) *Transformer {
	return &Transformer{value: value}
}

// Transform for env.
func (e *Transformer) Transform() string {
	return os.Getenv(e.value)
}
