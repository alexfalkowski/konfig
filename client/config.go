package client

import (
	"github.com/alexfalkowski/konfig/client/v1/config"
)

// Config for client.
type Config struct {
	V1 *config.Config `yaml:"v1,omitempty" json:"v1,omitempty" toml:"v1,omitempty"`
}
