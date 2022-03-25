package grpc

import (
	"context"
	"fmt"
	"net/http"

	tgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	GRPCServer *grpc.Server
	HTTPServer *http.Server
	Config     *tgrpc.Config
	Logger     *zap.Logger
}

// Register server.
func Register(lc fx.Lifecycle, params RegisterParams) error {
	server := NewServer()
	RegisterConfiguratorServer(params.GRPCServer, server)

	var conn *grpc.ClientConn

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error

			cparams := &tgrpc.ClientParams{
				Host:   fmt.Sprintf("127.0.0.1:%s", params.Config.Port),
				Config: params.Config,
				Logger: params.Logger,
			}

			conn, err = tgrpc.NewClient(ctx, cparams, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return err
			}

			mux := params.HTTPServer.Handler.(*runtime.ServeMux)

			return RegisterConfiguratorHandler(ctx, mux, conn)
		},
		OnStop: func(ctx context.Context) error {
			if conn != nil {
				return conn.Close()
			}

			return nil
		},
	})

	return nil
}
