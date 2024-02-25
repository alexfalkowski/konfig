package s3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alexfalkowski/go-service/file"
	"github.com/alexfalkowski/go-service/meta"
	tm "github.com/alexfalkowski/go-service/transport/meta"
	source "github.com/alexfalkowski/konfig/source/configurator"
	cerrors "github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/alexfalkowski/konfig/source/configurator/s3/telemetry/tracer"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for s3.
func NewConfigurator(cfg Config, t tracer.Tracer, client *http.Client) *Configurator {
	return &Configurator{cfg: cfg, tracer: t, client: client}
}

// Configurator for s3.
type Configurator struct {
	cfg    Config
	tracer tracer.Tracer
	client *http.Client
}

// GetConfig for s3.
func (c *Configurator) GetConfig(ctx context.Context, params source.ConfigParams) (*source.Config, error) {
	path := c.path(params.Application, params.Version, params.Environment, params.Continent, params.Country, params.Command, params.Kind)

	ctx, span := c.span(ctx)
	defer span.End()

	resolver := aws.EndpointResolverWithOptionsFunc(func(_, region string, _ ...any) (aws.Endpoint, error) {
		url := os.Getenv("AWS_URL")
		if url != "" {
			return aws.Endpoint{PartitionID: "aws", URL: url, SigningRegion: region}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	opts := []func(*config.LoadOptions) error{
		config.WithEndpointResolverWithOptions(resolver),
		config.WithHTTPClient(c.client),
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	out, err := client.GetObject(ctx, &s3.GetObjectInput{Bucket: &c.cfg.Bucket, Key: &path})
	if err != nil {
		meta.WithAttribute(ctx, "s3GetObjectError", err.Error())

		var nerr *types.NoSuchKey
		if errors.As(err, &nerr) {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)

			return nil, cerrors.ErrNotFound
		}

		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return nil, err
	}

	data, err := io.ReadAll(out.Body)
	if err != nil {
		meta.WithAttribute(ctx, "s3ReadAllError", err.Error())

		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

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
	ctx, span := c.tracer.Start(ctx, "get-config", trace.WithSpanKind(trace.SpanKindClient))
	ctx = tm.WithTraceID(ctx, span.SpanContext().TraceID().String())

	return ctx, span
}
