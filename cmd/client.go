package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/transport"
	ccmd "github.com/alexfalkowski/konfig/client/cmd"
	v1 "github.com/alexfalkowski/konfig/client/v1"
	"github.com/alexfalkowski/konfig/config"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	fx.NopLogger, fx.Provide(NewVersion), fx.Provide(ccmd.NewOutputConfig),
	marshaller.Module, cmd.Module,
	config.Module, logger.ZapModule,
	transport.GRPCModule, v1.Module,
}
