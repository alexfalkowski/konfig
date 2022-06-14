package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/server/config/provider/ssm"
	"github.com/alexfalkowski/konfig/source"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func v1SourceConfig(cfg config.Configurator) *source.Config {
	return &cfg.(*Config).Server.V1.Source
}

func v1SSMConfig(cfg config.Configurator) *ssm.Config {
	return &cfg.(*Config).Server.V1.Provider.SSM
}

func clientConfig(cfg config.Configurator) *client.Config {
	return &cfg.(*Config).Client
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}
