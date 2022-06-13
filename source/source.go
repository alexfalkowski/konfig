package source

import (
	"errors"

	"github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
	"github.com/alexfalkowski/konfig/source/configurator/s3"
	"github.com/alexfalkowski/konfig/source/configurator/trace/opentracing"
	"go.uber.org/fx"
)

// ErrNoConfigurator is defined in the config.
var ErrNoConfigurator = errors.New("no configurator")

// ConfiguratorParams for source.
type ConfiguratorParams struct {
	fx.In

	Config *Config
	Tracer opentracing.Tracer
}

// NewConfigurator for source.
func NewConfigurator(params ConfiguratorParams) (configurator.Configurator, error) {
	var configurator configurator.Configurator

	switch {
	case params.Config.IsGit():
		configurator = git.NewConfigurator(params.Config.Git, params.Tracer)
	case params.Config.IsFolder():
		configurator = folder.NewConfigurator(params.Config.Folder)
	case params.Config.IsS3():
		configurator = s3.NewConfigurator(params.Config.S3, params.Tracer)
	default:
		return nil, ErrNoConfigurator
	}

	configurator = opentracing.NewConfigurator(configurator, params.Tracer)

	return configurator, nil
}
