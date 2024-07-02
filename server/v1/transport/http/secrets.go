package http

import (
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/konfig/server/service"
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
		service *service.Service
	}
)

func (h *secretsHandler) Handle(ctx http.Context, req *GetSecretsRequest) (*GetSecretsResponse, error) {
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
