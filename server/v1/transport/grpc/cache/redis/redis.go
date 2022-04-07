package redis

import (
	"github.com/alexfalkowski/go-service/cache/redis"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/cache/redis/trace/opentracing"

	"github.com/go-redis/cache/v8"
)

// NewServer for redis.
func NewServer(cfg *redis.Config, cache *cache.Cache, server v1.ConfiguratorServiceServer) v1.ConfiguratorServiceServer {
	var s v1.ConfiguratorServiceServer = &Server{cache: cache, ConfiguratorServiceServer: server}
	s = opentracing.NewServer(cfg, s)

	return s
}
