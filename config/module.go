package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfig/config/provider"
	"github.com/alexfalkowski/konfig/config/provider/env"
	"github.com/alexfalkowski/konfig/config/provider/vault"
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
		fx.Provide(gitConfig),
		fx.Provide(clientConfig),
		fx.Provide(healthConfig),
		fx.Provide(env.NewTransformer),
		fx.Provide(vault.NewConfig),
		fx.Provide(vault.NewClient),
		fx.Provide(vault.NewTransformer),
		fx.Provide(NewTransformer),
		fx.Provide(provider.NewTransformer),
	)
)
