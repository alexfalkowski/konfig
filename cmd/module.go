package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(

	cmd.Module,
	env.Module,
	fx.Provide(NewVersion),
)
