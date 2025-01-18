package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	cmd.Module,
	fx.Provide(NewVersion),
)
