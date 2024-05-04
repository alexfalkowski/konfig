package client

import (
	"github.com/alexfalkowski/konfig/client/v1/config"
)

// IsEnabled for client.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

// Config for client.
type Config struct {
	V1 *config.Config `yaml:"v1,omitempty" json:"v1,omitempty" toml:"v1,omitempty"`
}
