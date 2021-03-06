package git

import (
	"os"
)

// Config for git.
type Config struct {
	URL   string `yaml:"url"`
	Dir   string `yaml:"dir"`
	Token string `yaml:"token"`
}

// GetToken that is specified in config or from KONFIG_GIT_TOKEN env variable.
func (c *Config) GetToken() string {
	if c.Token == "" {
		return os.Getenv("KONFIG_GIT_TOKEN")
	}

	return c.Token
}
