package source

import (
	gopentracing "github.com/alexfalkowski/konfig/source/configurator/git/opentracing"
	sopentracing "github.com/alexfalkowski/konfig/source/configurator/s3/opentracing"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(gopentracing.NewTracer),
		fx.Provide(sopentracing.NewTracer),
		fx.Provide(NewConfigurator),
	)
)
