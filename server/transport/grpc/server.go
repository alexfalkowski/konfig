package grpc

import (
	tgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/ratelimit"
	"github.com/alexfalkowski/go-service/transport/meta"
	"github.com/dgraph-io/ristretto"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor for gRPC.
func UnaryServerInterceptor(cfg *tgrpc.Config, cache *ristretto.Cache) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		ratelimit.UnaryServerInterceptor(&cfg.RateLimit, cache, meta.UserAgent),
	}
}

// StreamServerInterceptor for gRPC.
func StreamServerInterceptor(cfg *tgrpc.Config, cache *ristretto.Cache) []grpc.StreamServerInterceptor {
	return []grpc.StreamServerInterceptor{
		ratelimit.StreamServerInterceptor(&cfg.RateLimit, cache, meta.UserAgent),
	}
}
