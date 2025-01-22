package http

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
)

type (
	// GetSecretsRequest a map of name and secret.
	GetSecretsRequest struct {
		Secrets map[string]string `json:"secrets,omitempty"`
	}

	// GetSecretsResponse a map of meta and secrets.
	GetSecretsResponse struct {
		Meta    map[string]string `json:"meta,omitempty"`
		Secrets map[string][]byte `json:"secrets,omitempty"`
	}
)

// GetSecrets for HTTP.
func (h *Handler) GetSecrets(ctx context.Context, req *GetSecretsRequest) (*GetSecretsResponse, error) {
	resp := &GetSecretsResponse{}
	secrets, err := h.service.GetSecrets(ctx, req.Secrets)

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Secrets = secrets

	return resp, h.error(err)
}
