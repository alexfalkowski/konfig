package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/auth"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/source"
)

// Config for the service.
type Config struct {
	Auth          *auth.Config   `yaml:"auth,omitempty" json:"auth,omitempty" toml:"auth,omitempty"`
	Source        *source.Config `yaml:"source,omitempty" json:"source,omitempty" toml:"source,omitempty"`
	Client        *client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Health        *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
