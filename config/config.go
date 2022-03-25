package config

import (
	"github.com/alexfalkowski/go-service/config"
	"gopkg.in/yaml.v3"
)

// Config for the service.
type Config struct {
	config.Config `yaml:",inline"`
}

func (cfg *Config) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, cfg)
}
