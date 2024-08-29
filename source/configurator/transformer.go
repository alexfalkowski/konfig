package configurator

import (
	"bytes"
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/provider"
)

var (
	// ErrDecodeError in config.
	ErrDecodeError = errors.New("decode issue")

	// ErrEncodeError in config.
	ErrEncodeError = errors.New("encode issue")

	// ErrTraverseError in config.
	ErrTraverseError = errors.New("traverse issue")
)

// Transformer for config.
type Transformer struct {
	pt  *provider.Transformer
	enc *encoding.Map
}

// NewTransformer for config.
func NewTransformer(pt *provider.Transformer, enc *encoding.Map) *Transformer {
	return &Transformer{pt: pt, enc: enc}
}

// Transform config.
func (t *Transformer) Transform(ctx context.Context, kind string, data []byte) ([]byte, error) {
	m := t.enc.Get(kind)
	c := map[string]any{}

	if err := m.Decode(bytes.NewReader(data), &c); err != nil {
		meta.WithAttribute(ctx, "configDecodeError", meta.Error(err))

		return nil, ErrDecodeError
	}

	if err := t.traverse(ctx, c); err != nil {
		meta.WithAttribute(ctx, "configTraverseError", meta.Error(err))

		return nil, ErrTraverseError
	}

	var b bytes.Buffer
	if err := m.Encode(&b, c); err != nil {
		meta.WithAttribute(ctx, "configEncodeError", meta.Error(err))

		return nil, ErrEncodeError
	}

	return b.Bytes(), nil
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
