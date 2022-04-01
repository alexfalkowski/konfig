package config

import (
	"context"

	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/config/errors"
	"github.com/alexfalkowski/konfig/config/provider"
)

// Transformer for config.
type Transformer struct {
	pt *provider.Transformer
}

// NewTransformer for config.
func NewTransformer(pt *provider.Transformer) *Transformer {
	return &Transformer{pt: pt}
}

// Transform config.
func (t *Transformer) Transform(ctx context.Context, bytes []byte) ([]byte, error) {
	cfg := config.Map{}
	if err := config.UnmarshalFromBytes(bytes, cfg); err != nil {
		meta.WithAttribute(ctx, "config.unmarshal_error", err.Error())

		return nil, errors.ErrUnmarshalError
	}

	if err := t.traverse(cfg); err != nil {
		meta.WithAttribute(ctx, "config.traverse_error", err.Error())

		return nil, errors.ErrTraverseError
	}

	data, err := config.MarshalToBytes(cfg)
	if err != nil {
		meta.WithAttribute(ctx, "config.marshal_error", err.Error())

		return nil, errors.ErrMarshalError
	}

	return data, nil
}

func (t *Transformer) traverse(cfg config.Map) error {
	for key, val := range cfg {
		switch v := val.(type) {
		case string:
			vt, err := t.pt.Transform(v)
			if err != nil {
				return err
			}

			cfg[key] = vt
		case config.Map:
			if err := t.traverse(v); err != nil {
				return err
			}
		}
	}

	return nil
}
