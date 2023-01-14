package server

import (
	"github.com/alexfalkowski/konfig/server/v1/config"
)

// Config for server.
type Config struct {
	V1 config.Config `yaml:"v1" json:"v1" toml:"v1"`
}
