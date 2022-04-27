package cmd

import (
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/trace"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/konfig/config"
	"github.com/alexfalkowski/konfig/worker/health"
	"go.uber.org/fx"
)

// WorkerOptions for cmd.
var WorkerOptions = []fx.Option{
	fx.NopLogger, fx.Provide(NewVersion), config.Module, health.Module,
	logger.ZapModule, metrics.PrometheusModule,
	transport.GRPCServerModule, transport.GRPCJaegerModule,
	transport.HTTPServerModule, transport.HTTPJaegerModule,
	trace.JaegerOpenTracingModule,
}
