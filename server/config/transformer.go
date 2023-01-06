package config

import (
	"context"

	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/server/config/errors"
	"github.com/alexfalkowski/konfig/server/config/provider"
)

// Transformer for config.
type Transformer struct {
	pt *provider.Transformer
	m  *marshaller.YAML
}

// NewTransformer for config.
func NewTransformer(pt *provider.Transformer, m *marshaller.YAML) *Transformer {
	return &Transformer{pt: pt, m: m}
}

// Transform config.
func (t *Transformer) Transform(ctx context.Context, bytes []byte) ([]byte, error) {
	cfg := config.Map{}
	if err := t.m.Unmarshal(bytes, cfg); err != nil {
		meta.WithAttribute(ctx, "config.unmarshal_error", err.Error())

		return nil, errors.ErrUnmarshalError
	}

	if err := t.traverse(ctx, cfg); err != nil {
		meta.WithAttribute(ctx, "config.traverse_error", err.Error())

		return nil, errors.ErrTraverseError
	}

	data, err := t.m.Marshal(cfg)
	if err != nil {
		meta.WithAttribute(ctx, "config.marshal_error", err.Error())

		return nil, errors.ErrMarshalError
	}

	return data, nil
}

func (t *Transformer) traverse(ctx context.Context, cfg config.Map) error {
	for key, val := range cfg {
		switch v := val.(type) {
		case string:
			vt, err := t.pt.Transform(ctx, v)
			if err != nil {
				return err
			}

			cfg[key] = vt
		case config.Map:
			if err := t.traverse(ctx, v); err != nil {
				return err
			}
		}
	}

	return nil
}
