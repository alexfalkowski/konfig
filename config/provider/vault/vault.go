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
func (t *Transformer) Transform(ctx context.Context, value string) (string, error) {
	s, err := t.client.Logical().ReadWithContext(ctx, value)
	if err != nil {
		return "", err
	}

	return s.Data["data"].(map[string]any)["value"].(string), nil
}
