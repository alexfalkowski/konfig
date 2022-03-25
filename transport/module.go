package transport

import (
	tgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/konfig/transport/grpc"
	"go.uber.org/fx"
)

var (
	// GRPCServerModule for fx.
	GRPCServerModule = fx.Options(
		fx.Provide(tgrpc.NewServer),
		fx.Provide(grpc.UnaryServerInterceptor),
		fx.Provide(grpc.StreamServerInterceptor),
	)
)
