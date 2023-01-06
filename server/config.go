package server

import (
	v1 "github.com/alexfalkowski/konfig/server/v1"
)

// Config for server.
type Config struct {
	V1 v1.Config `yaml:"v1" json:"v1"`
}
