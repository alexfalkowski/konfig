package v1

import (
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(grpc.NewServer),
		fx.Invoke(grpc.Register),
	)
)
