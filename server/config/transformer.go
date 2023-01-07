package config

import (
	"context"

	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/server/config/errors"
	"github.com/alexfalkowski/konfig/server/config/provider"
	source "github.com/alexfalkowski/konfig/source/configurator"
)

// Transformer for config.
type Transformer struct {
	pt *provider.Transformer
	f  *marshaller.Factory
}

// NewTransformer for config.
func NewTransformer(pt *provider.Transformer, f *marshaller.Factory) *Transformer {
	return &Transformer{pt: pt, f: f}
}

// Transform config.
func (t *Transformer) Transform(ctx context.Context, c *source.Config) ([]byte, error) {
	m, err := t.f.Create(c.Kind)
	if err != nil {
		return nil, err
	}

	cfg := map[string]any{}
	if err := m.Unmarshal(c.Data, cfg); err != nil {
		meta.WithAttribute(ctx, "config.unmarshal_error", err.Error())

		return nil, errors.ErrUnmarshalError
	}

	if err := t.traverse(ctx, cfg); err != nil {
		meta.WithAttribute(ctx, "config.traverse_error", err.Error())

		return nil, errors.ErrTraverseError
	}

	data, err := m.Marshal(cfg)
	if err != nil {
		meta.WithAttribute(ctx, "config.marshal_error", err.Error())

		return nil, errors.ErrMarshalError
	}

	return data, nil
}

func (t *Transformer) traverse(ctx context.Context, cfg map[string]any) error {
	for key, val := range cfg {
		switch v := val.(type) {
		case string:
			vt, err := t.pt.Transform(ctx, v)
			if err != nil {
				return err
			}

			cfg[key] = vt
		case map[string]any:
			if err := t.traverse(ctx, v); err != nil {
				return err
			}
		}
	}

	return nil
}
