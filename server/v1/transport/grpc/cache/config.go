package cache

import (
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/cache/redis"
)

// Config for cache.
type Config struct {
	Redis redis.Config `yaml:"redis"`
}
