package source

import (
	"errors"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
	cs3 "github.com/alexfalkowski/konfig/source/configurator/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/go-github/v63/github"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ErrNoConfigurator is defined in the config.
var ErrNoConfigurator = errors.New("no configurator")

// ConfiguratorParams for source.
type ConfiguratorParams struct {
	fx.In

	Config    *Config
	S3Client  *s3.Client
	GitClient *github.Client
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	UserAgent env.UserAgent
}

// NewConfigurator for source.
func NewConfigurator(params ConfiguratorParams) (configurator.Configurator, error) {
	config := params.Config
	if !IsEnabled(config) {
		return nil, ErrNoConfigurator
	}

	var configurator configurator.Configurator

	switch {
	case config.IsFolder():
		configurator = folder.NewConfigurator(params.Config.Folder, params.Tracer)
	case config.IsGit():
		configurator = git.NewConfigurator(params.GitClient, params.Config.Git, params.Tracer)
	case config.IsS3():
		configurator = cs3.NewConfigurator(params.S3Client, params.Config.S3, params.Tracer)
	default:
		return nil, ErrNoConfigurator
	}

	return configurator, nil
}
