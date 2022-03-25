package cmd

import (
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/trace"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/konfig/config"
	"github.com/alexfalkowski/konfig/health"
	ktransport "github.com/alexfalkowski/konfig/transport"
	"go.uber.org/fx"
)

// ServeOptions for cmd.
var ServeOptions = []fx.Option{
	fx.NopLogger, config.Module, health.Module,
	logger.ZapModule, metrics.PrometheusModule,
	transport.HTTPServerModule, ktransport.GRPCServerModule,
	trace.JaegerOpenTracingModule,
}
