package config

import (
	"github.com/alexfalkowski/go-service/config"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(ConfiguratorModule, config.UnmarshalModule, ConfigModule)

	// ConfiguratorModule for fx.
	ConfiguratorModule = fx.Provide(NewConfigurator)

	// ConfigModule for fx.
	ConfigModule = fx.Options(
		config.ConfigModule,
		fx.Provide(vcsConfig),
	)
)
