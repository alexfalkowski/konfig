package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/alexfalkowski/go-service/telemetry/tracer"
	ks "github.com/alexfalkowski/konfig/internal/aws/s3"
	"github.com/alexfalkowski/konfig/internal/source/configurator/errors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for s3.
func NewConfigurator(client *s3.Client, config *Config, tracer trace.Tracer) *Configurator {
	return &Configurator{client: client, config: config, tracer: tracer}
}

// Configurator for s3.
type Configurator struct {
	client *s3.Client
	config *Config
	tracer trace.Tracer
}

// GetConfig for s3.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, country, cmd, kind string) ([]byte, error) {
	ctx, span := c.span(ctx)
	defer span.End()

	path := c.path(app, ver, env, continent, country, cmd, kind)

	out, err := c.client.GetObject(ctx, &s3.GetObjectInput{Bucket: &c.config.Bucket, Key: &path})
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		if ks.IsNotFound(err) {
			return nil, fmt.Errorf("%w: %w", err, errors.ErrNotFound)
		}

		return nil, err
	}

	tracer.Meta(ctx, span)

	return io.ReadAll(out.Body)
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
