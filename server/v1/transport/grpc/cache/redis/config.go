package redis

import (
	"time"
)

// Config for redis.
type Config struct {
	TTL time.Duration `yaml:"ttl"`
}
