package grpc

import (
	tgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/ratelimit"
	"github.com/alexfalkowski/go-service/transport/meta"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor for gRPC.
func UnaryServerInterceptor(cfg *tgrpc.Config) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		ratelimit.UnaryServerInterceptor(&cfg.RateLimit, meta.UserAgent),
	}
}

// StreamServerInterceptor for gRPC.
func StreamServerInterceptor(cfg *tgrpc.Config) []grpc.StreamServerInterceptor {
	return []grpc.StreamServerInterceptor{
		ratelimit.StreamServerInterceptor(&cfg.RateLimit, meta.UserAgent),
	}
}
