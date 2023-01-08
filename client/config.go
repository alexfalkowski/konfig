package client

import (
	"time"
)

// Config for client.
type Config struct {
	Host        string        `yaml:"host" json:"host"`
	Timeout     time.Duration `yaml:"timeout" json:"timeout"`
	Application string        `yaml:"application" json:"application"`
	Version     string        `yaml:"version" json:"version"`
	Environment string        `yaml:"environment" json:"environment"`
	Continent   string        `yaml:"continent" json:"continent"`
	Country     string        `yaml:"country" json:"country"`
	Command     string        `yaml:"command" json:"command"`
	Kind        string        `yaml:"kind" json:"kind"`
	Mode        uint32        `yaml:"mode" json:"mode"`
}
