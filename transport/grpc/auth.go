package grpc

import (
	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/go-service/security/token"
	gt "github.com/alexfalkowski/go-service/transport/grpc/security/token"
	"google.golang.org/grpc"
)

// IsAuth for GRPC.
func IsAuth(c *token.Config) bool {
	return c != nil && c.Kind == "auth"
}

// UnaryServerInterceptor for gRPC.
func UnaryServerInterceptor(cfg *token.Config, tkn *client.Token) []grpc.UnaryServerInterceptor {
	if !IsAuth(cfg) {
		return nil
	}

	return []grpc.UnaryServerInterceptor{
		gt.UnaryServerInterceptor(tkn.Verifier("jwt", "konfig", "get-config")),
	}
}

// StreamServerInterceptor for gRPC.
func StreamServerInterceptor(cfg *token.Config, tkn *client.Token) []grpc.StreamServerInterceptor {
	if !IsAuth(cfg) {
		return nil
	}

	return []grpc.StreamServerInterceptor{
		gt.StreamServerInterceptor(tkn.Verifier("jwt", "konfig", "get-config")),
	}
}
