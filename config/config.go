package config

import (
	av1 "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/auth"
	"github.com/alexfalkowski/konfig/client"
	v1 "github.com/alexfalkowski/konfig/client/v1/config"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/source"
)

// NewConfigurator for config.
func NewConfig(i *cmd.InputConfig) (*Config, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

// IsEnabled for config.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

// Config for the service.
type Config struct {
	Auth           *auth.Config   `yaml:"auth,omitempty" json:"auth,omitempty" toml:"auth,omitempty"`
	Source         *source.Config `yaml:"source,omitempty" json:"source,omitempty" toml:"source,omitempty"`
	Client         *client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	if !IsEnabled(cfg) {
		return nil
	}

	return cfg.Config
}

func v1Client(cfg *Config) *v1.Config {
	if !IsEnabled(cfg) || !client.IsEnabled(cfg.Client) {
		return nil
	}

	return cfg.Client.V1
}

func v1AuthClientConfig(cfg *Config) *av1.Config {
	if !IsEnabled(cfg) || !auth.IsEnabled(cfg.Auth) {
		return nil
	}

	return cfg.Auth.Client.V1
}

func healthConfig(cfg *Config) *health.Config {
	if !IsEnabled(cfg) {
		return nil
	}

	return cfg.Health
}

func sourceConfig(cfg *Config) *source.Config {
	if !IsEnabled(cfg) {
		return nil
	}

	return cfg.Source
}
