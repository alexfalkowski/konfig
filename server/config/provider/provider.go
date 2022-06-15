package provider

import (
	"context"
	"strings"

	"github.com/alexfalkowski/konfig/server/config/provider/env"
	"github.com/alexfalkowski/konfig/server/config/provider/ssm"
	"github.com/alexfalkowski/konfig/server/config/provider/vault"
)

const argumentsLen = 2

// Transformer for provider.
type Transformer struct {
	et *env.Transformer
	vt *vault.Transformer
	st *ssm.Transformer
}

// NewTransformer for provider.
func NewTransformer(et *env.Transformer, vt *vault.Transformer, st *ssm.Transformer) *Transformer {
	return &Transformer{et: et, vt: vt, st: st}
}

// Transform for provider.
func (t *Transformer) Transform(ctx context.Context, value string) (any, error) {
	args := strings.Split(value, ":")
	if len(args) != argumentsLen {
		return value, nil
	}

	switch args[0] {
	case "env":
		return t.et.Transform(ctx, args[1])
	case "vault":
		return t.vt.Transform(ctx, args[1])
	case "ssm":
		return t.st.Transform(ctx, args[1])
	default:
		return value, nil
	}
}
