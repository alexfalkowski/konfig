package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/server"
)

// Config for the service.
type Config struct {
	Server        server.Config `yaml:"server" json:"server" toml:"server"`
	Client        client.Config `yaml:"client" json:"client" toml:"client"`
	Health        health.Config `yaml:"health" json:"health" toml:"health"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
