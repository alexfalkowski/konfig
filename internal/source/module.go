package source

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewTransformer),
	fx.Provide(NewConfigurator),
)
