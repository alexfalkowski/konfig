package s3

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/alexfalkowski/go-service/telemetry/tracer"
	ke "github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for s3.
func NewConfigurator(client *s3.Client, cfg *Config, t trace.Tracer) *Configurator {
	return &Configurator{client: client, cfg: cfg, tracer: t}
}

// Configurator for s3.
type Configurator struct {
	client *s3.Client
	cfg    *Config
	tracer trace.Tracer
}

// GetConfig for s3.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, country, cmd, kind string) ([]byte, error) {
	path := c.path(app, ver, env, continent, country, cmd, kind)

	ctx, span := c.span(ctx)
	defer span.End()

	out, err := c.client.GetObject(ctx, &s3.GetObjectInput{Bucket: &c.cfg.Bucket, Key: &path})
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		var nerr *types.NoSuchKey
		if errors.As(err, &nerr) {
			return nil, ke.ErrNotFound
		}

		return nil, err
	}

	data, err := io.ReadAll(out.Body)
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

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

func (c *Configurator) span(ctx context.Context) (context.Context, trace.Span) {
	ctx, span := c.tracer.Start(ctx, operationName("get config"), trace.WithSpanKind(trace.SpanKindClient))
	ctx = tracer.WithTraceID(ctx, span)

	return ctx, span
}

func operationName(name string) string {
	return tracer.OperationName("s3", name)
}
