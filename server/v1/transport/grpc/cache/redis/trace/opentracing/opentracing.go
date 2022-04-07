package opentracing

import (
	"github.com/alexfalkowski/go-service/cache/redis"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
)

// NewServer for opentracing.
func NewServer(cfg *redis.Config, server v1.ConfiguratorServiceServer) *Server {
	return &Server{cfg: cfg, ConfiguratorServiceServer: server}
}
