package folder

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/konfig/source/configurator/errors"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for folder.
func NewConfigurator(config *Config, tracer trace.Tracer) *Configurator {
	return &Configurator{config: config, tracer: tracer}
}

// Configurator for folder.
type Configurator struct {
	config *Config
	tracer trace.Tracer
}

// GetConfig for folder.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, country, cmd, kind string) ([]byte, error) {
	ctx, span := c.span(ctx)
	defer span.End()

	if _, err := os.Stat(c.config.Dir); os.IsNotExist(err) {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		return nil, err
	}

	p := c.path(app, ver, env, continent, country, cmd, kind)
	path := filepath.Join(c.config.Dir, p)

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		if os.IsNotExist(err) {
			meta.WithAttribute(ctx, "folderError", meta.Error(err))

			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	tracer.Meta(ctx, span)

	return data, nil
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
