package provider

import (
	"github.com/alexfalkowski/konfig/provider/env"
	"github.com/alexfalkowski/konfig/provider/ssm"
	"github.com/alexfalkowski/konfig/provider/vault"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(env.NewTransformer),
	fx.Provide(vault.NewConfig),
	fx.Provide(vault.NewClient),
	fx.Provide(vault.NewTransformer),
	fx.Provide(ssm.NewClient),
	fx.Provide(ssm.NewTransformer),
	fx.Provide(NewTransformer),
)
