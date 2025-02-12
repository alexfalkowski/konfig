package provider

import (
	"context"
	"strings"

	"github.com/alexfalkowski/konfig/internal/provider/env"
	"github.com/alexfalkowski/konfig/internal/provider/ssm"
	"github.com/alexfalkowski/konfig/internal/provider/vault"
)

type transformer interface {
	Transform(ctx context.Context, value string) (string, error)
	IsMissing(err error) bool
}

type transformers map[string]transformer

// Transformer for provider.
type Transformer struct {
	ts transformers
}

// NewTransformer for provider.
func NewTransformer(et *env.Transformer, vt *vault.Transformer, st *ssm.Transformer) *Transformer {
	ts := transformers{
		"env":   et,
		"vault": vt,
		"ssm":   st,
	}

	return &Transformer{ts: ts}
}

// Transform for provider.
func (t *Transformer) Transform(ctx context.Context, value string) (string, error) {
	args := strings.Split(value, ":")
	if len(args) != 2 {
		return value, nil
	}

	tr, ok := t.ts[args[0]]
	if !ok {
		return value, nil
	}

	a, err := tr.Transform(ctx, args[1])
	if err != nil {
		if tr.IsMissing(err) {
			return value, nil
		}
	}

	return a, err
}
