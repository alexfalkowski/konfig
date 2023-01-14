package config

import (
	"github.com/alexfalkowski/konfig/source"
)

// Config for v1.
type Config struct {
	Source source.Config `yaml:"source" json:"source" toml:"source"`
}
