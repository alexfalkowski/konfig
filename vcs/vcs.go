package vcs

import (
	"context"

	"github.com/alexfalkowski/konfig/vcs/git"
)

// Configurator for vcs.
type Configurator interface {
	GetConfig(ctx context.Context, app, ver, env, cmd string) ([]byte, error)
}

// NewConfigurator for vcs.
func NewConfigurator(cfg *git.Config) Configurator {
	return git.NewConfigurator(cfg)
}
