package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/alexfalkowski/go-service/meta"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/go-redis/cache/v8"
)

const (
	cacheError = "server.cache_error"
	cached     = "server.cached"
)

var (
	ttl = 24 * time.Hour
)

// Server for redis.
type Server struct {
	cache *cache.Cache
	v1.ConfiguratorServiceServer
}

// GetConfig for redis.
func (s *Server) GetConfig(ctx context.Context, req *v1.GetConfigRequest) (*v1.GetConfigResponse, error) {
	key := key(req)

	var resp v1.GetConfigResponse

	err := s.cache.Get(ctx, key, &resp)
	if err == cache.ErrCacheMiss {
		ctx = meta.WithAttribute(ctx, cached, "false")

		resp, err := s.ConfiguratorServiceServer.GetConfig(ctx, req)
		if err != nil {
			return nil, err
		}

		if err := s.cache.Set(&cache.Item{Ctx: ctx, Key: key, Value: resp, TTL: ttl}); err != nil {
			meta.WithAttribute(ctx, cacheError, err.Error())
		}

		return resp, nil
	}

	if err != nil {
		ctx = meta.WithAttribute(ctx, cached, "false")
		ctx = meta.WithAttribute(ctx, cacheError, err.Error())

		return s.ConfiguratorServiceServer.GetConfig(ctx, req)
	}

	meta.WithAttribute(ctx, cached, "true")

	return &resp, nil
}

func key(req *v1.GetConfigRequest) string {
	return fmt.Sprintf("%s:%s:%s:%s", req.Application, req.Version, req.Environment, req.Command)
}
