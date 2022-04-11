package source

import (
	"context"

	"github.com/alexfalkowski/konfig/source/git"
)

// Configurator for source.
type Configurator interface {
	GetConfig(ctx context.Context, app, ver, env, cmd string) ([]byte, error)
}

// NewConfigurator for source.
func NewConfigurator(cfg *git.Config) Configurator {
	return git.NewConfigurator(cfg)
}
