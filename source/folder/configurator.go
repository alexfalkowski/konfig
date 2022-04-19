package folder

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/source/errors"
)

// Configurator for folder.
type Configurator struct {
	cfg *Config
}

// GetConfig for folder.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, cluster, cmd string) ([]byte, error) {
	var path string

	if cluster == "*" {
		path = filepath.Join(c.cfg.Dir, fmt.Sprintf("%s/%s/%s/%s.config.yml", app, ver, env, cmd))
	} else {
		path = filepath.Join(c.cfg.Dir, fmt.Sprintf("%s/%s/%s/%s/%s.config.yml", app, ver, env, cluster, cmd))
	}

	data, err := os.ReadFile(path)
	if err != nil {
		meta.WithAttribute(ctx, "folder.file_error", err.Error())

		return nil, errors.ErrNotFound
	}

	return data, nil
}
