package config

import (
	av1 "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
	v1 "github.com/alexfalkowski/konfig/client/v1/config"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/source"
)

// NewConfigurator for config.
func NewConfigurator(i *cmd.InputConfig) (config.Configurator, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

func v1Client(cfg config.Configurator) *v1.Config {
	return cfg.(*Config).Client.V1
}

func v1AuthClientConfig(cfg config.Configurator) *av1.Config {
	return cfg.(*Config).Auth.Client.V1
}

func healthConfig(cfg config.Configurator) *health.Config {
	return cfg.(*Config).Health
}

func sourceConfig(cfg config.Configurator) *source.Config {
	return cfg.(*Config).Source
}
