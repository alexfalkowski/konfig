package redis

import (
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/go-redis/cache/v8"
)

// NewServer for redis.
func NewServer(cache *cache.Cache, server v1.ConfiguratorServiceServer) *Server {
	return &Server{cache: cache, ConfiguratorServiceServer: server}
}
