package http

import (
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	nh "github.com/alexfalkowski/go-service/net/http"
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

func (h *secretsHandler) Handle(ctx nh.Context, req *GetSecretsRequest) (*GetSecretsResponse, error) {
	secrets, err := h.service.GetSecrets(ctx, req.Secrets)
	resp := &GetSecretsResponse{
		Meta:    meta.CamelStrings(ctx, ""),
		Secrets: secrets,
	}

	return resp, err
}

func (h *secretsHandler) Status(err error) int {
	if service.IsInvalidArgument(err) {
		return http.StatusBadRequest
	}

	if service.IsNotFound(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
