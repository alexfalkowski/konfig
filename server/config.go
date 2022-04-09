package server

import (
	v1 "github.com/alexfalkowski/konfig/server/v1"
	"github.com/alexfalkowski/konfig/vcs"
)

// Config for server.
type Config struct {
	VCS vcs.Config `yaml:"vcs"`
	V1  v1.Config  `yaml:"v1"`
}
