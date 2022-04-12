package source

import (
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(fx.Provide(NewConfigurator))
)
