package provider

import (
	"context"
	"strings"

	"github.com/alexfalkowski/konfig/internal/provider/env"
	"github.com/alexfalkowski/konfig/internal/provider/file"
	"github.com/alexfalkowski/konfig/internal/provider/ssm"
	"github.com/alexfalkowski/konfig/internal/provider/transformer"
	"github.com/alexfalkowski/konfig/internal/provider/vault"
	"go.uber.org/fx"
)

// TransformerParams for provider.
type TransformerParams struct {
	fx.In

	ENV   *env.Transformer
	Vault *vault.Transformer
	SSM   *ssm.Transformer
	File  *file.Transformer
}

// NewTransformer for provider.
func NewTransformer(params TransformerParams) *Transformer {
	ts := transformer.Transformers{
		"env":   params.ENV,
		"vault": params.Vault,
		"ssm":   params.SSM,
		"file":  params.File,
	}

	return &Transformer{ts: ts}
}

// Transformer for provider.
type Transformer struct {
	ts transformer.Transformers
}

// Transform for provider.
func (t *Transformer) Transform(ctx context.Context, value string) (string, error) {
	k, v, ok := strings.Cut(value, ":")
	if !ok {
		return value, nil
	}

	tr, ok := t.ts[k]
	if !ok {
		return value, nil
	}

	a, err := tr.Transform(ctx, v)
	if err != nil {
		if tr.IsMissing(err) {
			return value, nil
		}
	}

	return a, err
}
