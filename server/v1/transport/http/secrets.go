package http

import (
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/konfig/server/config"
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

	secretsHandler struct {
		service *config.Configuration
	}
)

func (h *secretsHandler) Handle(ctx rpc.Context, req *GetSecretsRequest) (*GetSecretsResponse, error) {
	resp := &GetSecretsResponse{}

	secrets, err := h.service.GetSecrets(ctx, req.Secrets)
	if err != nil {
		resp.Meta = meta.CamelStrings(ctx, "")

		return resp, handleError(err)
	}

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Secrets = secrets

	return resp, nil
}
