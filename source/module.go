package source

import (
	"github.com/alexfalkowski/konfig/source/configurator"
	gotel "github.com/alexfalkowski/konfig/source/configurator/git/otel"
	sotel "github.com/alexfalkowski/konfig/source/configurator/s3/otel"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(gotel.NewTracer),
		fx.Provide(sotel.NewTracer),
		fx.Provide(configurator.NewTransformer),
		fx.Provide(NewConfigurator),
	)
)
