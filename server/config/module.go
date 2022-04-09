package config

import (
	"github.com/alexfalkowski/konfig/server/config/provider"
	"github.com/alexfalkowski/konfig/server/config/provider/env"
	"github.com/alexfalkowski/konfig/server/config/provider/vault"

	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(env.NewTransformer),
		fx.Provide(vault.NewConfig),
		fx.Provide(vault.NewClient),
		fx.Provide(vault.NewTransformer),
		fx.Provide(NewTransformer),
		fx.Provide(provider.NewTransformer),
	)
)
