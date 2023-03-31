package provider

import (
	"github.com/alexfalkowski/konfig/provider/env"
	"github.com/alexfalkowski/konfig/provider/ssm"
	sotel "github.com/alexfalkowski/konfig/provider/ssm/otel"
	"github.com/alexfalkowski/konfig/provider/vault"
	votel "github.com/alexfalkowski/konfig/provider/vault/otel"

	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(env.NewTransformer),
		fx.Provide(vault.NewConfig),
		fx.Provide(vault.NewClient),
		fx.Provide(vault.NewTransformer),
		fx.Provide(votel.NewTracer),
		fx.Provide(ssm.NewClient),
		fx.Provide(ssm.NewTransformer),
		fx.Provide(sotel.NewTracer),
		fx.Provide(NewTransformer),
	)
)
