package grpc

import (
	"github.com/alexfalkowski/auth/client"
	t "github.com/alexfalkowski/go-service/security/token"
	gt "github.com/alexfalkowski/go-service/transport/grpc/security/token"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor for gRPC.
func UnaryServerInterceptor(cfg *t.Config, tkn *client.Token) []grpc.UnaryServerInterceptor {
	if cfg.Kind != "auth" {
		return nil
	}

	return []grpc.UnaryServerInterceptor{
		gt.UnaryServerInterceptor(tkn.Verifier("jwt", "konfig", "get-config")),
	}
}

// StreamServerInterceptor for gRPC.
func StreamServerInterceptor(cfg *t.Config, tkn *client.Token) []grpc.StreamServerInterceptor {
	if cfg.Kind != "auth" {
		return nil
	}

	return []grpc.StreamServerInterceptor{
		gt.StreamServerInterceptor(tkn.Verifier("jwt", "konfig", "get-config")),
	}
}
