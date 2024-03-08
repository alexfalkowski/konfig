package config

import (
	"github.com/alexfalkowski/go-service/client"
)

// Config for client.
type Config struct {
	Application   string `yaml:"application" json:"application" toml:"application"`
	Version       string `yaml:"version" json:"version" toml:"version"`
	Environment   string `yaml:"environment" json:"environment" toml:"environment"`
	Continent     string `yaml:"continent" json:"continent" toml:"continent"`
	Country       string `yaml:"country" json:"country" toml:"country"`
	Command       string `yaml:"command" json:"command" toml:"command"`
	Kind          string `yaml:"kind" json:"kind" toml:"kind"`
	Mode          uint32 `yaml:"mode" json:"mode" toml:"mode"`
	client.Config `yaml:",inline" json:",inline" toml:",inline"`
}
