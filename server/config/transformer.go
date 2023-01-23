package config

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/provider"
	source "github.com/alexfalkowski/konfig/source/configurator"
)

var (
	// ErrUnmarshalError in config.
	ErrUnmarshalError = errors.New("unmarshal issue")

	// ErrMarshalError in config.
	ErrMarshalError = errors.New("marshal issue")

	// ErrTraverseError in config.
	ErrTraverseError = errors.New("traverse issue")
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
func (t *Transformer) Transform(ctx context.Context, cfg *source.Config) ([]byte, error) {
	m, err := t.f.Create(cfg.Kind)
	if err != nil {
		return nil, err
	}

	c := map[string]any{}
	if err := m.Unmarshal(cfg.Data, &c); err != nil {
		meta.WithAttribute(ctx, "config.unmarshal_error", err.Error())

		return nil, ErrUnmarshalError
	}

	if err := t.traverse(ctx, c); err != nil {
		meta.WithAttribute(ctx, "config.traverse_error", err.Error())

		return nil, ErrTraverseError
	}

	data, err := m.Marshal(c)
	if err != nil {
		meta.WithAttribute(ctx, "config.marshal_error", err.Error())

		return nil, ErrMarshalError
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
