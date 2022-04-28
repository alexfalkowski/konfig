package vault

import (
	"context"

	"github.com/hashicorp/vault/api"
)

// Transformer for vault.
type Transformer struct {
	client *api.Client
}

// NewTransformer for vault.
func NewTransformer(client *api.Client) *Transformer {
	return &Transformer{client: client}
}

// Transform for vault.
func (t *Transformer) Transform(ctx context.Context, value string) (any, error) {
	sec, err := t.client.Logical().ReadWithContext(ctx, value)
	if err != nil {
		return value, err
	}

	if sec == nil || sec.Data == nil {
		return value, nil
	}

	d := sec.Data["data"]
	if d == nil {
		return value, nil
	}

	md, ok := d.(map[string]any)
	if !ok || md["value"] == nil {
		return value, nil
	}

	return md["value"], nil
}
