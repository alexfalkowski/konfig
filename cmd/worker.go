package cmd

import (
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/trace"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/konfig/config"
	"go.uber.org/fx"
)

// WorkerOptions for cmd.
var WorkerOptions = []fx.Option{
	fx.NopLogger, config.Module,
	logger.ZapModule, metrics.PrometheusModule,
	transport.HTTPServerModule, trace.JaegerOpenTracingModule,
}
