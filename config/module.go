package config

import (
	"github.com/alexfalkowski/go-service/config"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(NewConfigurator),
		config.UnmarshalModule,
		config.ConfigModule,
		fx.Provide(v1SourceConfig),
		fx.Provide(clientConfig),
		fx.Provide(healthConfig),
		fx.Provide(v1RedisConfig),
	)
)
