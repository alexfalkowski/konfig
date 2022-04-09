package v1

import (
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/cache"
)

// Config for v1.
type Config struct {
	Cache cache.Config `yaml:"cache"`
}
