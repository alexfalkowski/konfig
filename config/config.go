package config

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/source"
)

// NewConfigurator for config.
func NewConfig(i *cmd.InputConfig) (*Config, error) {
	c := &Config{}

	return c, i.Decode(c)
}

// IsEnabled for config.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

// Config for the service.
type Config struct {
	Source         *source.Config `yaml:"source,omitempty" json:"source,omitempty" toml:"source,omitempty"`
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	if !IsEnabled(cfg) {
		return nil
	}

	return cfg.Config
}

func sourceConfig(cfg *Config) *source.Config {
	if !IsEnabled(cfg) {
		return nil
	}

	return cfg.Source
}

func healthConfig(cfg *Config) *health.Config {
	if !IsEnabled(cfg) {
		return nil
	}

	return cfg.Health
}
