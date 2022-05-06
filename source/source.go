package source

import (
	"github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
)

// NewConfigurator for source.
func NewConfigurator(cfg *Config) configurator.Configurator {
	if cfg.IsGit() {
		return git.NewConfigurator(&cfg.Git)
	}

	if cfg.IsFolder() {
		return folder.NewConfigurator(&cfg.Folder)
	}

	return nil
}
