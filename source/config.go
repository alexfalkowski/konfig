package source

import (
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
)

// Config for source.
type Config struct {
	Type   string        `yaml:"type"`
	Git    git.Config    `yaml:"git"`
	Folder folder.Config `yaml:"folder"`
}

// IsGit configured.
func (c *Config) IsGit() bool {
	return c.Type == "git"
}

// IsFolder configured.
func (c *Config) IsFolder() bool {
	return c.Type == "folder"
}
