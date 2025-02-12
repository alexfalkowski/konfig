package source

import (
	"github.com/alexfalkowski/konfig/internal/source/configurator"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(configurator.NewTransformer),
	fx.Provide(NewConfigurator),
)
