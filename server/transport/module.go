package transport

import (
	tgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/konfig/server/transport/grpc"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(tgrpc.NewServer),
		fx.Provide(grpc.UnaryServerInterceptor),
		fx.Provide(grpc.StreamServerInterceptor),
	)
)
