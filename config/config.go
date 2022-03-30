package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/server"
	"gopkg.in/yaml.v3"
)

// Config for the service.
type Config struct {
	Server        server.Config `yaml:"server"`
	Client        client.Config `yaml:"client"`
	config.Config `yaml:",inline"`
}

func (cfg *Config) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, cfg)
}
