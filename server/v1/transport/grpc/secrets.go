package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetSecrets for gRPC.
func (s *Server) GetSecrets(ctx context.Context, req *v1.GetSecretsRequest) (*v1.GetSecretsResponse, error) {
	resp := &v1.GetSecretsResponse{}
	secs := req.GetSecrets()
	secrets := make(map[string][]byte, len(secs))

	for n, v := range secs {
		t, err := s.provider.Transform(ctx, v)
		if err != nil {
			ctx = meta.WithAttribute(ctx, "secretsError", meta.Error(err))
			resp.Meta = s.meta(ctx)

			return resp, status.Error(codes.Internal, "could not transform")
		}

		switch c := t.(type) {
		case string:
			secrets[n] = []byte(c)
		default:
			secrets[n] = []byte(v)
		}
	}

	resp.Meta = s.meta(ctx)
	resp.Secrets = secrets

	return resp, nil
}
