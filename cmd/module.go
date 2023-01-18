package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		cmd.Module,
		fx.Provide(NewVersion),
	)
)
