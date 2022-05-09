package config

import (
	"github.com/alexfalkowski/konfig/server/config/provider"
	"github.com/alexfalkowski/konfig/server/config/provider/env"
	"github.com/alexfalkowski/konfig/server/config/provider/vault"
	"github.com/alexfalkowski/konfig/server/config/provider/vault/trace/opentracing"

	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(env.NewTransformer),
		fx.Provide(vault.NewConfig),
		fx.Provide(vault.NewClient),
		fx.Provide(vault.NewTransformer),
		fx.Provide(opentracing.NewTracer),
		fx.Provide(NewTransformer),
		fx.Provide(provider.NewTransformer),
	)
)
