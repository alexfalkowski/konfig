package provider

import (
	"github.com/alexfalkowski/konfig/provider/env"
	"github.com/alexfalkowski/konfig/provider/ssm"
	st "github.com/alexfalkowski/konfig/provider/ssm/telemetry/tracer"
	"github.com/alexfalkowski/konfig/provider/vault"
	vt "github.com/alexfalkowski/konfig/provider/vault/telemetry/tracer"

	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(env.NewTransformer),
		fx.Provide(vault.NewConfig),
		fx.Provide(vault.NewClient),
		fx.Provide(vault.NewTransformer),
		fx.Provide(vt.NewTracer),
		fx.Provide(ssm.NewClient),
		fx.Provide(ssm.NewTransformer),
		fx.Provide(st.NewTracer),
		fx.Provide(NewTransformer),
	)
)
