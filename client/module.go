package client

import (
	v1 "github.com/alexfalkowski/konfig/client/v1"
	"go.uber.org/fx"
)

var (
	// ConfigModule for fx.
	ConfigModule = fx.Options(
		v1.Module,
		fx.Provide(NewClient),
		fx.Invoke(GetConfig),
	)

	// SecretsModule for fx.
	SecretsModule = fx.Options(
		v1.Module,
		fx.Provide(NewClient),
		fx.Invoke(WriteSecrets),
	)
)
