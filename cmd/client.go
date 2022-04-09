package cmd

import (
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/trace"
	v1 "github.com/alexfalkowski/konfig/client/v1"
	"github.com/alexfalkowski/konfig/config"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	fx.NopLogger, config.Module, logger.ZapModule,
	trace.JaegerOpenTracingModule, v1.Module,
}
