package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
)

// GetSecrets for gRPC.
func (s *Server) GetSecrets(ctx context.Context, req *v1.GetSecretsRequest) (*v1.GetSecretsResponse, error) {
	resp := &v1.GetSecretsResponse{}
	secrets, err := s.service.GetSecrets(ctx, req.GetSecrets())

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Secrets = secrets

	return resp, s.error(err)
}
