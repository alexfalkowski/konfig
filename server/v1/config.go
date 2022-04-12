package v1

import (
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/cache"
	"github.com/alexfalkowski/konfig/source"
)

// Config for v1.
type Config struct {
	Cache  cache.Config  `yaml:"cache"`
	Source source.Config `yaml:"source"`
}
