package config

import (
	"time"

	"github.com/alexfalkowski/go-service/retry"
)

// Config for client.
type Config struct {
	Host        string        `yaml:"host" json:"host" toml:"host"`
	Timeout     time.Duration `yaml:"timeout" json:"timeout" toml:"timeout"`
	Retry       retry.Config  `yaml:"retry" json:"retry" toml:"retry"`
	UserAgent   string        `yaml:"user_agent" json:"user_agent" toml:"user_agent"`
	Application string        `yaml:"application" json:"application" toml:"application"`
	Version     string        `yaml:"version" json:"version" toml:"version"`
	Environment string        `yaml:"environment" json:"environment" toml:"environment"`
	Continent   string        `yaml:"continent" json:"continent" toml:"continent"`
	Country     string        `yaml:"country" json:"country" toml:"country"`
	Command     string        `yaml:"command" json:"command" toml:"command"`
	Kind        string        `yaml:"kind" json:"kind" toml:"kind"`
	Mode        uint32        `yaml:"mode" json:"mode" toml:"mode"`
}
