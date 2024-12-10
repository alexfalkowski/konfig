package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/go-service/crypto"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	crypto.Module,
	fx.Provide(config.NewConfig[Config]),
	config.Module,
	fx.Decorate(decorateConfig),
	fx.Provide(sourceConfig),
	fx.Provide(healthConfig),
)
