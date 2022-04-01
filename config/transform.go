package config

import (
	"context"

	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/config/errors"
	"github.com/alexfalkowski/konfig/config/provider"
)

// Transform the config.
func Transform(ctx context.Context, bytes []byte) ([]byte, error) {
	cfg := config.Map{}
	if err := config.UnmarshalFromBytes(bytes, cfg); err != nil {
		meta.WithAttribute(ctx, "config.unmarshal_error", err.Error())

		return nil, errors.ErrUnmarshalError
	}

	traverse(cfg)

	data, err := config.MarshalToBytes(cfg)
	if err != nil {
		meta.WithAttribute(ctx, "config.marshal_error", err.Error())

		return nil, errors.ErrMarshalError
	}

	return data, nil
}

func traverse(cfg config.Map) {
	for key, val := range cfg {
		switch v := val.(type) {
		case string:
			t := provider.NewTransformer(v)
			if t != nil {
				cfg[key] = t.Transform()
			}
		case config.Map:
			traverse(v)
		}
	}
}
