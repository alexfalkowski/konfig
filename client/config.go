package client

import (
	"github.com/alexfalkowski/konfig/client/v1/config"
)

// Config for client.
type Config struct {
	V1 config.Config `yaml:"v1" json:"v1" toml:"v1"`
}
