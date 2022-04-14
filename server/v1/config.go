package v1

import (
	"github.com/alexfalkowski/konfig/source"
)

// Config for v1.
type Config struct {
	Source source.Config `yaml:"source"`
}
