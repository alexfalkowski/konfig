package git

import (
	"os"
)

// Config for git.
type Config struct {
	URL string `yaml:"url,omitempty" json:"url,omitempty" toml:"url,omitempty"`
}

// Token for git.
func (c *Config) Token() string {
	return os.Getenv("KONFIG_GIT_TOKEN")
}
