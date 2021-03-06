package v1

import (
	"github.com/alexfalkowski/konfig/server/config/provider"
	"github.com/alexfalkowski/konfig/source"
)

// Config for v1.
type Config struct {
	Provider provider.Config `yaml:"provider"`
	Source   source.Config   `yaml:"source"`
}
