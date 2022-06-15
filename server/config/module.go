package config

import (
	"github.com/alexfalkowski/konfig/server/config/provider"
	"github.com/alexfalkowski/konfig/server/config/provider/env"
	"github.com/alexfalkowski/konfig/server/config/provider/ssm"
	sopentracing "github.com/alexfalkowski/konfig/server/config/provider/ssm/trace/opentracing"
	"github.com/alexfalkowski/konfig/server/config/provider/vault"
	vopentracing "github.com/alexfalkowski/konfig/server/config/provider/vault/trace/opentracing"

	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(env.NewTransformer),
		fx.Provide(vault.NewConfig),
		fx.Provide(vault.NewClient),
		fx.Provide(vault.NewTransformer),
		fx.Provide(vopentracing.NewTracer),
		fx.Provide(ssm.NewClient),
		fx.Provide(ssm.NewTransformer),
		fx.Provide(sopentracing.NewTracer),
		fx.Provide(NewTransformer),
		fx.Provide(provider.NewTransformer),
	)
)
