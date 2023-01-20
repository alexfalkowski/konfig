package git

import (
	"os"
)

// Config for git.
type Config struct {
	URL string `yaml:"url" json:"url" toml:"url"`
	Dir string `yaml:"dir" json:"dir" toml:"dir"`
}

// Token for git.
func (c *Config) Token() string {
	return os.Getenv("KONFIG_GIT_TOKEN")
}
