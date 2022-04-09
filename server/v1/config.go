package v1

import (
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/cache"
	"github.com/alexfalkowski/konfig/vcs"
)

// Config for v1.
type Config struct {
	Cache cache.Config `yaml:"cache"`
	VCS   vcs.Config   `yaml:"vcs"`
}
