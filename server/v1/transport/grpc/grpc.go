package grpc

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/cache/redis"
	sgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	shttp "github.com/alexfalkowski/go-service/transport/http"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/config"
	tredis "github.com/alexfalkowski/konfig/server/v1/transport/grpc/cache/redis"
	"github.com/alexfalkowski/konfig/vcs"
	"github.com/go-redis/cache/v8"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	GRPCServer   *grpc.Server
	HTTPServer   *shttp.Server
	GRPCConfig   *sgrpc.Config
	RedisConfig  *redis.Config
	Logger       *zap.Logger
	Configurator vcs.Configurator
	Transformer  *config.Transformer
	Cache        *cache.Cache
}

// Register server.
func Register(lc fx.Lifecycle, params RegisterParams) {
	server := NewServer(params.Configurator, params.Transformer, params.RedisConfig, params.Cache)
	v1.RegisterConfiguratorServiceServer(params.GRPCServer, server)

	var conn *grpc.ClientConn

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn, _ = sgrpc.NewClient(ctx, fmt.Sprintf("127.0.0.1:%s", params.GRPCConfig.Port), params.GRPCConfig, params.Logger,
				sgrpc.WithClientDialOption(grpc.WithBlock()),
			)

			return v1.RegisterConfiguratorServiceHandler(ctx, params.HTTPServer.Mux, conn)
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})
}

// NewServer for gRPC.
func NewServer(conf vcs.Configurator, trans *config.Transformer, cfg *redis.Config, cache *cache.Cache) v1.ConfiguratorServiceServer {
	var server v1.ConfiguratorServiceServer = &Server{conf: conf, trans: trans}
	server = tredis.NewServer(cfg, cache, server)

	return server
}
