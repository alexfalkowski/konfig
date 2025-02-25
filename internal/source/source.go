package source

import (
	"errors"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/konfig/internal/source/folder"
	"github.com/alexfalkowski/konfig/internal/source/git"
	cs3 "github.com/alexfalkowski/konfig/internal/source/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/go-github/v69/github"
	"go.uber.org/fx"
)

// ErrNoConfigurator is defined in the config.
var ErrNoConfigurator = errors.New("no configurator")

// ConfiguratorParams for source.
type ConfiguratorParams struct {
	fx.In
	Tracer    *tracer.Tracer
	Meter     *metrics.Meter
	FS        os.FileSystem
	Config    *Config
	S3Client  *s3.Client
	GitClient *github.Client
	Logger    *logger.Logger
	UserAgent env.UserAgent
}

// NewConfigurator for source.
func NewConfigurator(params ConfiguratorParams) (Configurator, error) {
	config := params.Config
	if !IsEnabled(config) {
		return nil, ErrNoConfigurator
	}

	var configurator Configurator

	switch {
	case config.IsFolder():
		configurator = folder.NewConfigurator(params.Config.Folder, params.FS, params.Tracer)
	case config.IsGit():
		configurator = git.NewConfigurator(params.GitClient, params.Config.Git, params.Tracer)
	case config.IsS3():
		configurator = cs3.NewConfigurator(params.S3Client, params.Config.S3, params.Tracer)
	default:
		return nil, ErrNoConfigurator
	}

	return configurator, nil
}
