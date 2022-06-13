package source

import (
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
	"github.com/alexfalkowski/konfig/source/configurator/s3"
)

// Config for source.
type Config struct {
	Kind   string        `yaml:"kind"`
	Git    git.Config    `yaml:"git"`
	Folder folder.Config `yaml:"folder"`
	S3     s3.Config     `yaml:"s3"`
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
