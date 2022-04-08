package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/vcs/git"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func gitConfig(cfg config.Configurator) *git.Config {
	return &cfg.(*Config).Server.VCS.Git
}

func clientConfig(cfg config.Configurator) *client.Config {
	return &cfg.(*Config).Client
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}
