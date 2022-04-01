package provider

import (
	"strings"

	"github.com/alexfalkowski/konfig/config/provider/env"
	"github.com/alexfalkowski/konfig/config/provider/vault"
)

const argumentsLen = 2

// Transformer for provider.
type Transformer struct {
	et *env.Transformer
	vt *vault.Transformer
}

// NewTransformer for provider.
func NewTransformer(et *env.Transformer, vt *vault.Transformer) *Transformer {
	return &Transformer{et: et, vt: vt}
}

// Transform for provider.
func (t *Transformer) Transform(value string) (string, error) {
	args := strings.Split(value, ":")
	if len(args) != argumentsLen {
		return value, nil
	}

	switch args[0] {
	case "env":
		return t.et.Transform(args[1])
	case "vault":
		return t.vt.Transform(args[1])
	default:
		return value, nil
	}
}
