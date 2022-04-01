package provider

import (
	"strings"

	"github.com/alexfalkowski/konfig/config/provider/env"
)

const argumentsLen = 2

// Transformer for provider.
type Transformer interface {
	Transform() string
}

// NewTransformer for provider.
func NewTransformer(value string) Transformer {
	args := strings.Split(value, ":")
	if len(args) != argumentsLen {
		return nil
	}

	switch args[0] {
	case "env":
		return env.NewTransformer(args[1])
	default:
		return nil
	}
}
