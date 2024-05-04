package auth

import (
	"github.com/alexfalkowski/auth/client"
)

// IsEnabled for auth.
func IsEnabled(cfg *Config) bool {
	return cfg != nil && cfg.Client != nil
}

// Config for auth.
type Config struct {
	Client *client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
}
