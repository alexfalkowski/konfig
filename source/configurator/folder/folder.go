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
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, cmd string) ([]byte, error) {
	var path string

	if continent == "*" {
		path = filepath.Join(c.cfg.Dir, fmt.Sprintf("%s/%s/%s/%s.config.yml", app, ver, env, cmd))
	} else {
		path = filepath.Join(c.cfg.Dir, fmt.Sprintf("%s/%s/%s/%s/%s.config.yml", app, ver, env, continent, cmd))
	}

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		meta.WithAttribute(ctx, "folder.file_error", err.Error())

		return nil, errors.ErrNotFound
	}

	return data, nil
}
