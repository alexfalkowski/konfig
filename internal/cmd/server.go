package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	v1 "github.com/alexfalkowski/konfig/internal/api/v1"
	"github.com/alexfalkowski/konfig/internal/config"
	"github.com/alexfalkowski/konfig/internal/health"
	"github.com/alexfalkowski/konfig/internal/provider"
	"github.com/alexfalkowski/konfig/internal/source"
	"github.com/alexfalkowski/konfig/internal/token"
)

// RegisterServer for cmd.
func RegisterServer(command *cmd.Command) {
	flags := cmd.NewFlagSet("server")
	flags.AddInput("env:KONFIG_CONFIG_FILE")

	command.AddServer("server", "Start konfig server", flags,
		module.Module, debug.Module, feature.Module,
		transport.Module, telemetry.Module,
		health.Module, config.Module,
		provider.Module, source.Module, token.Module,
		v1.Module, cmd.Module,
	)
}
