package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/token"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	token.Module,
	fx.Provide(NewConfig),
	config.Module,
	fx.Decorate(decorateConfig),
	fx.Provide(sourceConfig),
	fx.Provide(healthConfig),
	fx.Provide(tokenConfig),
)
