package grpc

import (
	"context"
	"fmt"
	"net/http"

	tgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/config"
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/cache/redis"
	"github.com/alexfalkowski/konfig/vcs"
	"github.com/go-redis/cache/v8"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	GRPCServer   *grpc.Server
	HTTPServer   *http.Server
	Config       *tgrpc.Config
	Logger       *zap.Logger
	Configurator vcs.Configurator
	Transformer  *config.Transformer
	Cache        *cache.Cache
}

// Register server.
func Register(lc fx.Lifecycle, params RegisterParams) {
	server := NewServer(params.Configurator, params.Transformer, params.Cache)
	v1.RegisterConfiguratorServiceServer(params.GRPCServer, server)

	var conn *grpc.ClientConn

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			cparams := &tgrpc.ClientParams{
				Host:   fmt.Sprintf("127.0.0.1:%s", params.Config.Port),
				Config: params.Config,
				Logger: params.Logger,
			}

			conn, _ = tgrpc.NewClient(ctx, cparams, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
			mux := params.HTTPServer.Handler.(*runtime.ServeMux)

			return v1.RegisterConfiguratorServiceHandler(ctx, mux, conn)
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})
}

// NewServer for gRPC.
func NewServer(conf vcs.Configurator, trans *config.Transformer, cache *cache.Cache) v1.ConfiguratorServiceServer {
	var server v1.ConfiguratorServiceServer = &Server{conf: conf, trans: trans}
	server = redis.NewServer(cache, server)

	return server
}
