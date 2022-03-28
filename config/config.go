package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/vcs"
	"gopkg.in/yaml.v3"
)

// Config for the service.
type Config struct {
	VCS           vcs.Config `yaml:"vcs"`
	config.Config `yaml:",inline"`
}

func (cfg *Config) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, cfg)
}
