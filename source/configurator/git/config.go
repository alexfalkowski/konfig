package git

import (
	"os"
	"path/filepath"
	"strings"
)

type (
	// Token for git.
	Token string

	// Config for git.
	Config struct {
		Token Token  `yaml:"token,omitempty" json:"token,omitempty" toml:"token,omitempty"`
		URL   string `yaml:"url,omitempty" json:"url,omitempty" toml:"url,omitempty"`
	}
)

// GetToken for git.
func (c *Config) GetToken() (string, error) {
	k, err := os.ReadFile(filepath.Clean(string(c.Token)))

	return strings.TrimSpace(string(k)), err
}
