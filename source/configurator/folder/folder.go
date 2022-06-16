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
	path := filepath.Join(c.cfg.Dir, c.path(app, ver, env, continent, country, cmd))

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		meta.WithAttribute(ctx, "folder.file_error", err.Error())

		return nil, errors.ErrNotFound
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
