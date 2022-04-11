package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/cache/redis"
	"github.com/alexfalkowski/konfig/source/git"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func v1GitConfig(cfg config.Configurator) *git.Config {
	return &cfg.(*Config).Server.V1.Source.Git
}

func clientConfig(cfg config.Configurator) *client.Config {
	return &cfg.(*Config).Client
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}

func v1RedisConfig(cfg config.Configurator) *redis.Config {
	return &cfg.(*Config).Server.V1.Cache.Redis
}
