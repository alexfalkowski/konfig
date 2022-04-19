package source

import (
	"context"

	"github.com/alexfalkowski/konfig/source/folder"
	"github.com/alexfalkowski/konfig/source/git"
)

// Configurator for source.
type Configurator interface {
	GetConfig(ctx context.Context, app, ver, env, cluster, cmd string) ([]byte, error)
}

// NewConfigurator for source.
func NewConfigurator(cfg *Config) Configurator {
	if cfg.IsGit() {
		return git.NewConfigurator(&cfg.Git)
	}

	if cfg.IsFolder() {
		return folder.NewConfigurator(&cfg.Folder)
	}

	return nil
}
