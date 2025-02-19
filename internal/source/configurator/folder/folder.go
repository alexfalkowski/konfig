package folder

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/konfig/internal/source/configurator/errors"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for folder.
func NewConfigurator(config *Config, fs os.FileSystem, tracer *tracer.Tracer) *Configurator {
	return &Configurator{config: config, fs: fs, tracer: tracer}
}

// Configurator for folder.
type Configurator struct {
	config *Config
	fs     os.FileSystem
	tracer *tracer.Tracer
}

// GetConfig for folder.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, country, cmd, kind string) ([]byte, error) {
	ctx, span := c.span(ctx)
	defer span.End()

	if !c.fs.PathExists(c.config.Dir) {
		err := fmt.Errorf("%s: %w", c.config.Dir, errors.ErrInvalidFolder)

		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		return nil, err
	}

	p := c.path(app, ver, env, continent, country, cmd, kind)
	path := filepath.Join(c.config.Dir, p)

	data, err := c.fs.ReadFile(path)
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		if c.fs.IsNotExist(err) {
			return nil, fmt.Errorf("%w: %w", err, errors.ErrNotFound)
		}

		return nil, err
	}

	tracer.Meta(ctx, span)

	return []byte(data), nil
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

//nolint:spancheck
func (c *Configurator) span(ctx context.Context) (context.Context, trace.Span) {
	ctx, span := c.tracer.Start(ctx, operationName("get config"), trace.WithSpanKind(trace.SpanKindClient))
	ctx = tracer.WithTraceID(ctx, span)

	return ctx, span
}

func operationName(name string) string {
	return tracer.OperationName("s3", name)
}
