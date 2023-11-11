package cmd

import (
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/konfig/config"
	"github.com/alexfalkowski/konfig/provider"
	"github.com/alexfalkowski/konfig/server/health"
	v1 "github.com/alexfalkowski/konfig/server/v1"
	"github.com/alexfalkowski/konfig/source"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, runtime.Module, Module, debug.Module,
	metrics.Module, telemetry.Module, health.Module,
	transport.Module, config.Module, provider.Module,
	source.Module, v1.Module,
}
