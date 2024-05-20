package folder

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/go-service/file"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	source "github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/errors"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for folder.
func NewConfigurator(cfg *Config, t trace.Tracer) *Configurator {
	return &Configurator{cfg: cfg, tracer: t}
}

// Configurator for folder.
type Configurator struct {
	cfg    *Config
	tracer trace.Tracer
}

// GetConfig for folder.
func (c *Configurator) GetConfig(ctx context.Context, params source.ConfigParams) (*source.Config, error) {
	ctx, span := c.span(ctx)
	defer span.End()

	if _, err := os.Stat(c.cfg.Dir); os.IsNotExist(err) {
		tracer.Error(err, span)

		return nil, err
	}

	p := c.path(params.Application, params.Version, params.Environment, params.Continent, params.Country, params.Command, params.Kind)
	path := filepath.Join(c.cfg.Dir, p)

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		tracer.Error(err, span)

		if os.IsNotExist(err) {
			meta.WithAttribute(ctx, "folderFileError", meta.Error(err))

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

func (c *Configurator) span(ctx context.Context) (context.Context, trace.Span) {
	ctx, span := c.tracer.Start(ctx, operationName("get config"), trace.WithSpanKind(trace.SpanKindClient))
	ctx = tracer.WithTraceID(ctx, span)

	return ctx, span
}

func operationName(name string) string {
	return tracer.OperationName("s3", name)
}
