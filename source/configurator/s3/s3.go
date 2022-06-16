package s3

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/alexfalkowski/konfig/source/configurator/s3/opentracing"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewConfigurator for s3.
func NewConfigurator(cfg Config, tracer opentracing.Tracer, client *http.Client) *Configurator {
	return &Configurator{cfg: cfg, tracer: tracer, client: client}
}

// Configurator for s3.
type Configurator struct {
	cfg    Config
	tracer opentracing.Tracer
	client *http.Client
}

// GetConfig for s3.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, country, cmd string) ([]byte, error) {
	path := c.path(app, ver, env, continent, country, cmd)

	ctx, span := opentracing.StartSpanFromContext(ctx, c.tracer, "get-object", fmt.Sprintf("%s:%s", c.cfg.Bucket, path))
	defer span.Finish()

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...any) (aws.Endpoint, error) {
		if c.cfg.URL != "" {
			return aws.Endpoint{PartitionID: "aws", URL: c.cfg.URL, SigningRegion: c.cfg.Region}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	opts := []func(*config.LoadOptions) error{
		config.WithRegion(c.cfg.Region),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithHTTPClient(c.client),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.cfg.Access, c.cfg.Secret, "")),
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	out, err := client.GetObject(ctx, &s3.GetObjectInput{Bucket: &c.cfg.Bucket, Key: &path})
	if err != nil {
		meta.WithAttribute(ctx, "s3.get_object_error", err.Error())

		return nil, errors.ErrNotFound
	}

	data, err := io.ReadAll(out.Body)
	if err != nil {
		meta.WithAttribute(ctx, "s3.read_all_error", err.Error())

		return nil, errors.ErrNotFound
	}

	return data, nil
}

func (c *Configurator) path(app, ver, env, continent, country, cmd string) string {
	if continent == "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s/%s.config.yml", app, ver, env, cmd)
	}

	if continent != "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s/%s/%s.config.yml", app, ver, env, continent, cmd)
	}

	return fmt.Sprintf("%s/%s/%s/%s/%s/%s.config.yml", app, ver, env, continent, country, cmd)
}
