package provider

import (
	"github.com/alexfalkowski/konfig/internal/provider/env"
	"github.com/alexfalkowski/konfig/internal/provider/file"
	"github.com/alexfalkowski/konfig/internal/provider/ssm"
	"github.com/alexfalkowski/konfig/internal/provider/vault"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(env.NewTransformer),
	fx.Provide(vault.NewConfig),
	fx.Provide(vault.NewClient),
	fx.Provide(vault.NewTransformer),
	fx.Provide(ssm.NewTransformer),
	fx.Provide(file.NewTransformer),
	fx.Provide(NewTransformer),
)
