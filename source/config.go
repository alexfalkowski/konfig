package source

import (
	"github.com/alexfalkowski/konfig/source/git"
)

// Config for source.
type Config struct {
	Git git.Config `yaml:"git"`
}
