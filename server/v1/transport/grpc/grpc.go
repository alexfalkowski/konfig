package grpc

import (
	"context"

	shttp "github.com/alexfalkowski/go-service/transport/http"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	GRPCServer *grpc.Server
	HTTPServer *shttp.Server
	Conn       *grpc.ClientConn
	Server     v1.ServiceServer
}

// Register server.
func Register(lc fx.Lifecycle, params RegisterParams) {
	v1.RegisterServiceServer(params.GRPCServer, params.Server)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return v1.RegisterServiceHandler(ctx, params.HTTPServer.Mux, params.Conn)
		},
	})
}
