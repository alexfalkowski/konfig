package config

import (
	"github.com/alexfalkowski/go-service/client"
)

type (
	// Configuration for client.
	Configuration struct {
		Application string `yaml:"application,omitempty" json:"application,omitempty" toml:"application,omitempty"`
		Version     string `yaml:"version,omitempty" json:"version,omitempty" toml:"version,omitempty"`
		Environment string `yaml:"environment,omitempty" json:"environment,omitempty" toml:"environment,omitempty"`
		Continent   string `yaml:"continent,omitempty" json:"continent,omitempty" toml:"continent,omitempty"`
		Country     string `yaml:"country,omitempty" json:"country,omitempty" toml:"country,omitempty"`
		Command     string `yaml:"command,omitempty" json:"command,omitempty" toml:"command,omitempty"`
		Kind        string `yaml:"kind,omitempty" json:"kind,omitempty" toml:"kind,omitempty"`
		Mode        uint32 `yaml:"mode,omitempty" json:"mode,omitempty" toml:"mode,omitempty"`
	}

	// Secrets for client.
	Secrets struct {
		Files map[string]string `yaml:"files,omitempty" json:"files,omitempty" toml:"files,omitempty"`
		Path  string            `yaml:"path,omitempty" json:"path,omitempty" toml:"path,omitempty"`
		Mode  uint32            `yaml:"mode,omitempty" json:"mode,omitempty" toml:"mode,omitempty"`
	}

	// Config for client.
	Config struct {
		*client.Config `yaml:",inline" json:",inline" toml:",inline"`
		Configuration  *Configuration `yaml:"config,omitempty" json:"config,omitempty" toml:"config,omitempty"`
		Secrets        *Secrets       `yaml:"secrets,omitempty" json:"secrets,omitempty" toml:"secrets,omitempty"`
	}
)
