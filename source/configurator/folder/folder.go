package folder

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/alexfalkowski/konfig/source/configurator/folder/opentracing"
)

// NewConfigurator for folder.
func NewConfigurator(cfg Config, tracer opentracing.Tracer) *Configurator {
	return &Configurator{cfg: cfg, tracer: tracer}
}

// Configurator for folder.
type Configurator struct {
	cfg    Config
	tracer opentracing.Tracer
}

// GetConfig for folder.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, cluster, cmd string) ([]byte, error) {
	var path string

	if cluster == "*" {
		path = filepath.Join(c.cfg.Dir, fmt.Sprintf("%s/%s/%s/%s.config.yml", app, ver, env, cmd))
	} else {
		path = filepath.Join(c.cfg.Dir, fmt.Sprintf("%s/%s/%s/%s/%s.config.yml", app, ver, env, cluster, cmd))
	}

	_, span := opentracing.StartSpanFromContext(ctx, c.tracer, "read-file", path)
	defer span.Finish()

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		meta.WithAttribute(ctx, "folder.file_error", err.Error())

		return nil, errors.ErrNotFound
	}

	return data, nil
}
