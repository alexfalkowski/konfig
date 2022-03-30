package server

import (
	"github.com/alexfalkowski/konfig/vcs"
)

// Config for server.
type Config struct {
	VCS vcs.Config `yaml:"vcs"`
}
