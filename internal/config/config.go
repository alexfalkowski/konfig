package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/internal/health"
	"github.com/alexfalkowski/konfig/internal/source"
)

// Config for the service.
type Config struct {
	Source         *source.Config `yaml:"source,omitempty" json:"source,omitempty" toml:"source,omitempty"`
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func sourceConfig(cfg *Config) *source.Config {
	return cfg.Source
}

func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}
