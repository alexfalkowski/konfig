package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewConfigurator for s3.
func NewConfigurator(cfg *Config) *Configurator {
	return &Configurator{cfg: cfg}
}

// Configurator for s3.
type Configurator struct {
	cfg *Config
}

// GetConfig for s3.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, cluster, cmd string) ([]byte, error) {
	// nolint:staticcheck
	resolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if c.cfg.URL != "" {
			return aws.Endpoint{PartitionID: "aws", URL: c.cfg.URL, SigningRegion: c.cfg.Region}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(c.cfg.Region), config.WithEndpointResolver(resolver))
	if err != nil {
		meta.WithAttribute(ctx, "s3.config_error", err.Error())

		return nil, errors.ErrNotFound
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	var path string

	if cluster == "*" {
		path = fmt.Sprintf("%s/%s/%s/%s.config.yml", app, ver, env, cmd)
	} else {
		path = fmt.Sprintf("%s/%s/%s/%s/%s.config.yml", app, ver, env, cluster, cmd)
	}

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
