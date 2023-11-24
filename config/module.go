package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/go-service/marshaller"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewConfigurator),
	marshaller.Module,
	config.ConfigModule,
	fx.Provide(v1Client), fx.Provide(v1AuthClientConfig),
	fx.Provide(sourceConfig), fx.Provide(healthConfig),
)
