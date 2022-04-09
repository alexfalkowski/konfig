package vcs

import (
	"github.com/alexfalkowski/konfig/vcs/git"
)

// Config for vcs.
type Config struct {
	Git git.Config `yaml:"git"`
}
