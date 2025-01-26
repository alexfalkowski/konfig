package configurator

import (
	"bytes"
	"context"

	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/errors"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/konfig/provider"
)

// NewTransformer for config.
func NewTransformer(pt *provider.Transformer, enc *encoding.Map) *Transformer {
	return &Transformer{pt: pt, enc: enc}
}

// Transformer for config.
type Transformer struct {
	pt  *provider.Transformer
	enc *encoding.Map
}

// Transform config.
//
//nolint:nonamedreturns
func (t *Transformer) Transform(ctx context.Context, kind string, data []byte) (transformed []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Prefix("transform", runtime.ConvertRecover(r))
		}
	}()

	var (
		c map[string]any
		b bytes.Buffer
	)

	m := t.enc.Get(kind)

	err = m.Decode(bytes.NewReader(data), &c)
	runtime.Must(err)

	err = t.traverse(ctx, c)
	runtime.Must(err)

	err = m.Encode(&b, c)
	runtime.Must(err)

	transformed = b.Bytes()

	return
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
