package redis

import (
	"github.com/alexfalkowski/go-service/cache/redis"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/cache/redis/trace/opentracing"

	"github.com/go-redis/cache/v8"
)

// NewServer for redis.
func NewServer(config *Config, redisConfig *redis.Config, cache *cache.Cache, server v1.ConfiguratorServiceServer) v1.ConfiguratorServiceServer {
	var s v1.ConfiguratorServiceServer = &Server{config: config, cache: cache, ConfiguratorServiceServer: server}
	s = opentracing.NewServer(redisConfig, s)

	return s
}
