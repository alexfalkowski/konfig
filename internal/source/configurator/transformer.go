package configurator

import (
	"context"

	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/errors"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/konfig/internal/provider"
)

// NewTransformer for config.
func NewTransformer(pt *provider.Transformer, enc *encoding.Map, pool *sync.BufferPool) *Transformer {
	return &Transformer{pt: pt, enc: enc, pool: pool}
}

// Transformer for config.
type Transformer struct {
	pt   *provider.Transformer
	enc  *encoding.Map
	pool *sync.BufferPool
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

	buffer := t.pool.Get()
	defer t.pool.Put(buffer)

	buffer.Write(data)

	var result map[string]any

	m := t.enc.Get(kind)

	err = m.Decode(buffer, &result)
	runtime.Must(err)

	err = t.traverse(ctx, result)
	runtime.Must(err)

	buffer.Reset()

	err = m.Encode(buffer, result)
	runtime.Must(err)

	transformed = t.pool.Copy(buffer)

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
