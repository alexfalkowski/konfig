package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/server"
)

// Config for the service.
type Config struct {
	Server        server.Config `yaml:"server"`
	Client        client.Config `yaml:"client"`
	Health        health.Config `yaml:"health"`
	config.Config `yaml:",inline"`
}
