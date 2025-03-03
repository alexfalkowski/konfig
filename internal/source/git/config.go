package git

import (
	"github.com/alexfalkowski/go-service/os"
)

// Config for git.
type Config struct {
	Token      Token  `yaml:"token,omitempty" json:"token,omitempty" toml:"token,omitempty"`
	Owner      string `yaml:"owner,omitempty" json:"owner,omitempty" toml:"owner,omitempty"`
	Repository string `yaml:"repository,omitempty" json:"repository,omitempty" toml:"repository,omitempty"`
}

// GetToken for git.
func (c *Config) GetToken() ([]byte, error) {
	return os.ReadFile(c.Token.String())
}

// Token for git.
type Token string

// String for Token.
func (t Token) String() string {
	return string(t)
}
