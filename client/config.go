package client

import (
	"time"
)

// Config for client.
type Config struct {
	Host        string        `yaml:"host"`
	Timeout     time.Duration `yaml:"timeout"`
	Application string        `yaml:"application"`
	Version     string        `yaml:"version"`
	Environment string        `yaml:"environment"`
	Continent   string        `yaml:"continent"`
	Country     string        `yaml:"country"`
	Command     string        `yaml:"command"`
}
