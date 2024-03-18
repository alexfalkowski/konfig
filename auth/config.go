package auth

import (
	"github.com/alexfalkowski/auth/client"
)

// Config for auth.
type Config struct {
	Client *client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
}
