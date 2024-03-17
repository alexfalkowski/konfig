package source

import (
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
	"github.com/alexfalkowski/konfig/source/configurator/s3"
)

// Config for source.
type Config struct {
	Kind   string         `yaml:"kind,omitempty" json:"kind,omitempty" toml:"kind,omitempty"`
	Git    *git.Config    `yaml:"git,omitempty" json:"git,omitempty" toml:"git,omitempty"`
	Folder *folder.Config `yaml:"folder,omitempty" json:"folder,omitempty" toml:"folder,omitempty"`
	S3     *s3.Config     `yaml:"s3,omitempty" json:"s3,omitempty" toml:"s3,omitempty"`
}

// IsGit configured.
func (c *Config) IsGit() bool {
	return c.Kind == "git"
}

// IsFolder configured.
func (c *Config) IsFolder() bool {
	return c.Kind == "folder"
}

// IsS3 configured.
func (c *Config) IsS3() bool {
	return c.Kind == "s3"
}
