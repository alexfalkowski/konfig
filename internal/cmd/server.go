package cmd

import (
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
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	module.Module, debug.Module, feature.Module,
	transport.Module, telemetry.Module,
	health.Module, config.Module,
	provider.Module, source.Module, token.Module,
	v1.Module, Module,
}
