package git

import (
	"os"
)

// NewConfigurator for git.
func NewConfigurator(cfg *Config) *Configurator {
	if cfg.Token == "" {
		cfg.Token = os.Getenv("KONFIG_GIT_TOKEN")
	}

	return &Configurator{cfg: cfg}
}
