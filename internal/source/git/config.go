package git

import (
	"github.com/alexfalkowski/go-service/os"
)

type (
	// Token for git.
	Token string

	// Config for git.
	Config struct {
		Token      Token  `yaml:"token,omitempty" json:"token,omitempty" toml:"token,omitempty"`
		Owner      string `yaml:"owner,omitempty" json:"owner,omitempty" toml:"owner,omitempty"`
		Repository string `yaml:"repository,omitempty" json:"repository,omitempty" toml:"repository,omitempty"`
	}
)

// GetToken for git.
func (c *Config) GetToken() (string, error) {
	return os.ReadFile(string(c.Token))
}
