package source

import (
	"github.com/alexfalkowski/konfig/source/configurator"
	gt "github.com/alexfalkowski/konfig/source/configurator/git/telemetry/tracer"
	st "github.com/alexfalkowski/konfig/source/configurator/s3/telemetry/tracer"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(gt.NewTracer),
		fx.Provide(st.NewTracer),
		fx.Provide(configurator.NewTransformer),
		fx.Provide(NewConfigurator),
	)
)
