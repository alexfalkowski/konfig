package http

import (
	"context"
	"net/http"

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
		Error   *Error            `json:"error,omitempty"`
		Secrets map[string][]byte `json:"secrets,omitempty"`
	}

	secretsErrorer struct{}
)

// GetSecrets for HTTP.
func (s *Server) GetSecrets(ctx context.Context, req *GetSecretsRequest) (*GetSecretsResponse, error) {
	secrets, err := s.service.GetSecrets(ctx, req.Secrets)
	resp := &GetSecretsResponse{
		Meta:    meta.CamelStrings(ctx, ""),
		Secrets: secrets,
	}

	return resp, err
}

func (*secretsErrorer) Error(ctx context.Context, err error) *GetSecretsResponse {
	return &GetSecretsResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (*secretsErrorer) Status(_ error) int {
	return http.StatusInternalServerError
}
