package config

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
	cv1 "github.com/alexfalkowski/konfig/client/v1/config"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/source"
)

// NewConfigurator for config.
func NewConfigurator(i *cmd.InputConfig) (config.Configurator, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

func v1Client(cfg config.Configurator) *cv1.Config {
	return &cfg.(*Config).Client.V1
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}

func sourceConfig(cfg config.Configurator) *source.Config {
	return &cfg.(*Config).Source
}
