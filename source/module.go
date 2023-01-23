package source

import (
	"github.com/alexfalkowski/konfig/source/configurator"
	gotr "github.com/alexfalkowski/konfig/source/configurator/git/opentracing"
	sotr "github.com/alexfalkowski/konfig/source/configurator/s3/opentracing"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(gotr.NewTracer),
		fx.Provide(sotr.NewTracer),
		fx.Provide(configurator.NewTransformer),
		fx.Provide(NewConfigurator),
	)
)
