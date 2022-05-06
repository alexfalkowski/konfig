package git

import (
	"os"

	"github.com/alexfalkowski/konfig/source/configurator/trace/opentracing"
)

// NewConfigurator for git.
func NewConfigurator(cfg *Config, tracer opentracing.Tracer) *Configurator {
	if cfg.Token == "" {
		cfg.Token = os.Getenv("KONFIG_GIT_TOKEN")
	}

	return &Configurator{cfg: cfg, tracer: tracer}
}
