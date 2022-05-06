package source

import (
	"github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
	"github.com/alexfalkowski/konfig/source/configurator/trace/opentracing"
	"go.uber.org/fx"
)

// ConfiguratorParams for source.
type ConfiguratorParams struct {
	fx.In

	Config *Config
	Tracer opentracing.Tracer
}

// NewConfigurator for source.
func NewConfigurator(params ConfiguratorParams) configurator.Configurator {
	var configurator configurator.Configurator

	if params.Config.IsGit() {
		configurator = git.NewConfigurator(&params.Config.Git)
	}

	if params.Config.IsFolder() {
		configurator = folder.NewConfigurator(&params.Config.Folder)
	}

	if configurator == nil {
		return nil
	}

	if params.Tracer != nil {
		configurator = opentracing.NewConfigurator(configurator, params.Tracer)
	}

	return configurator
}
