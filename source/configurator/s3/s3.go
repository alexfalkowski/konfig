package s3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/alexfalkowski/go-service/file"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/konfig/aws"
	source "github.com/alexfalkowski/konfig/source/configurator"
	ke "github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for s3.
func NewConfigurator(cfg *Config, t trace.Tracer, client *http.Client) *Configurator {
	return &Configurator{cfg: cfg, tracer: t, client: client}
}

// Configurator for s3.
type Configurator struct {
	cfg    *Config
	tracer trace.Tracer
	client *http.Client
}

// GetConfig for s3.
func (c *Configurator) GetConfig(ctx context.Context, params source.ConfigParams) (*source.Config, error) {
	path := c.path(params.Application, params.Version, params.Environment, params.Continent, params.Country, params.Command, params.Kind)

	ctx, span := c.span(ctx)
	defer span.End()

	opts := []func(*config.LoadOptions) error{
		config.WithEndpointResolverWithOptions(aws.EndpointResolver()),
		config.WithHTTPClient(c.client),
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	out, err := client.GetObject(ctx, &s3.GetObjectInput{Bucket: &c.cfg.Bucket, Key: &path})
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
