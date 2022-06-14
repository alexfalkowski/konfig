package source

import (
	"errors"

	"github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	fopentracing "github.com/alexfalkowski/konfig/source/configurator/folder/opentracing"
	"github.com/alexfalkowski/konfig/source/configurator/git"
	gopentracing "github.com/alexfalkowski/konfig/source/configurator/git/opentracing"
	"github.com/alexfalkowski/konfig/source/configurator/s3"
	sopentracing "github.com/alexfalkowski/konfig/source/configurator/s3/opentracing"
	"go.uber.org/fx"
)

// ErrNoConfigurator is defined in the config.
var ErrNoConfigurator = errors.New("no configurator")

// ConfiguratorParams for source.
type ConfiguratorParams struct {
	fx.In

	Config       *Config
	FolderTracer fopentracing.Tracer
	GitTracer    gopentracing.Tracer
	S3Tracer     sopentracing.Tracer
}

// NewConfigurator for source.
func NewConfigurator(params ConfiguratorParams) (configurator.Configurator, error) {
	var configurator configurator.Configurator

	switch {
	case params.Config.IsFolder():
		configurator = folder.NewConfigurator(params.Config.Folder, params.FolderTracer)
	case params.Config.IsGit():
		configurator = git.NewConfigurator(params.Config.Git, params.GitTracer)
	case params.Config.IsS3():
		configurator = s3.NewConfigurator(params.Config.S3, params.S3Tracer)
	default:
		return nil, ErrNoConfigurator
	}

	return configurator, nil
}
