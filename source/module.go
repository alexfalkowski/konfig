package source

import (
	"github.com/alexfalkowski/konfig/source/configurator/trace/opentracing"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(opentracing.NewTracer),
		fx.Provide(NewConfigurator),
	)
)
