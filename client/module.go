package client

import (
	v1 "github.com/alexfalkowski/konfig/client/v1"
	"go.uber.org/fx"
)

var (
	// ClientModule for fx.
	ClientModule = fx.Options(
		v1.Module,
		fx.Provide(NewClient),
	)

	// CommandModule for fx.
	CommandModule = fx.Options(
		ClientModule,
		fx.Provide(NewOutputConfig),
		fx.Invoke(RunCommand),
	)
)
