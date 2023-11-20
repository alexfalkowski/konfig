package auth

import (
	"github.com/alexfalkowski/auth/client"
)

// Config for auth.
type Config struct {
	Client client.Config `yaml:"client" json:"client" toml:"client"`
}
