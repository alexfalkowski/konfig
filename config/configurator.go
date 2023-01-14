package config

import (
	"github.com/alexfalkowski/go-service/config"
	cv1 "github.com/alexfalkowski/konfig/client/v1/config"
	"github.com/alexfalkowski/konfig/health"
	sv1 "github.com/alexfalkowski/konfig/server/v1/config"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func v1Server(cfg config.Configurator) *sv1.Config {
	return &cfg.(*Config).Server.V1
}

func v1Client(cfg config.Configurator) *cv1.Config {
	return &cfg.(*Config).Client.V1
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}
