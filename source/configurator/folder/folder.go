package folder

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/go-service/file"
	"github.com/alexfalkowski/go-service/meta"
	source "github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/errors"
)

// NewConfigurator for folder.
func NewConfigurator(cfg *Config) *Configurator {
	return &Configurator{cfg: cfg}
}

// Configurator for folder.
type Configurator struct {
	cfg *Config
}

// GetConfig for folder.
func (c *Configurator) GetConfig(ctx context.Context, params source.ConfigParams) (*source.Config, error) {
	if _, err := os.Stat(c.cfg.Dir); os.IsNotExist(err) {
		meta.WithAttribute(ctx, "folderDirError", meta.Error(err))

		return nil, err
	}

	p := c.path(params.Application, params.Version, params.Environment, params.Continent, params.Country, params.Command, params.Kind)
	path := filepath.Join(c.cfg.Dir, p)

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		meta.WithAttribute(ctx, "folderFileError", meta.Error(err))

		if os.IsNotExist(err) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return &source.Config{Kind: file.Extension(path), Data: data}, nil
}

func (c *Configurator) path(app, ver, env, continent, country, cmd, kind string) string {
	if continent == "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s/%s.%s", app, ver, env, cmd, kind)
	}

	if continent != "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s/%s/%s.%s", app, ver, env, continent, cmd, kind)
	}

	return fmt.Sprintf("%s/%s/%s/%s/%s/%s.%s", app, ver, env, continent, country, cmd, kind)
}
