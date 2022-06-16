package folder

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/source/configurator/errors"
)

// NewConfigurator for folder.
func NewConfigurator(cfg Config) *Configurator {
	return &Configurator{cfg: cfg}
}

// Configurator for folder.
type Configurator struct {
	cfg Config
}

// GetConfig for folder.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, country, cmd string) ([]byte, error) {
	if _, err := os.Stat(c.cfg.Dir); os.IsNotExist(err) {
		meta.WithAttribute(ctx, "folder.dir_error", err.Error())

		return nil, err
	}

	path := filepath.Join(c.cfg.Dir, c.path(app, ver, env, continent, country, cmd))

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		meta.WithAttribute(ctx, "folder.file_error", err.Error())

		if os.IsNotExist(err) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return data, nil
}

func (c *Configurator) path(app, ver, env, continent, country, cmd string) string {
	if continent == "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s/%s.config.yml", app, ver, env, cmd)
	}

	if continent != "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s/%s/%s.config.yml", app, ver, env, continent, cmd)
	}

	return fmt.Sprintf("%s/%s/%s/%s/%s/%s.config.yml", app, ver, env, continent, country, cmd)
}
