package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/go-service/crypto"
	"github.com/alexfalkowski/go-service/token"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	crypto.Module,
	token.Module,
	fx.Provide(NewConfig),
	config.Module,
	fx.Decorate(decorateConfig),
	fx.Provide(sourceConfig),
	fx.Provide(healthConfig),
)
