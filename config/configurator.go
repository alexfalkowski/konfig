package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/vcs"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func vcsConfig(cfg config.Configurator) *vcs.Config {
	return &cfg.(*Config).Server.VCS
}

func clientConfig(cfg config.Configurator) *client.Config {
	return &cfg.(*Config).Client
}
