package cmd

import (
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/trace"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/konfig/config"
	kconfig "github.com/alexfalkowski/konfig/server/config"
	"github.com/alexfalkowski/konfig/server/health"
	v1 "github.com/alexfalkowski/konfig/server/v1"
	"github.com/alexfalkowski/konfig/source"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, config.Module, kconfig.Module, health.Module,
	logger.ZapModule, metrics.PrometheusModule,
	transport.GRPCServerModule, transport.GRPCJaegerModule,
	transport.HTTPServerModule, transport.HTTPJaegerModule,
	trace.JaegerOpenTracingModule,
	source.Module, v1.Module,
}
