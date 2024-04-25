package grpc

import (
	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/go-service/os"
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
	its := []grpc.UnaryServerInterceptor{}

	if IsAuth(cfg) {
		its = append(its, gt.UnaryServerInterceptor(tkn.Verifier("jwt", os.ExecutableName(), "get-config")))
	}

	return its
}

// StreamServerInterceptor for gRPC.
func StreamServerInterceptor(cfg *token.Config, tkn *client.Token) []grpc.StreamServerInterceptor {
	its := []grpc.StreamServerInterceptor{}

	if IsAuth(cfg) {
		its = append(its, gt.StreamServerInterceptor(tkn.Verifier("jwt", os.ExecutableName(), "get-config")))
	}

	return its
}
