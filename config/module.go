package config

import (
	"github.com/alexfalkowski/go-service/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(config.NewConfig[Config]),
	config.Module,
	fx.Decorate(decorateConfig),
	fx.Provide(sourceConfig),
	fx.Provide(healthConfig),
)
