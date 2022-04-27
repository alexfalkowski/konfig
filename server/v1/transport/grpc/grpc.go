package grpc

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/cache/redis"
	sgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	shttp "github.com/alexfalkowski/go-service/transport/http"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/server/config"
	"github.com/alexfalkowski/konfig/source"
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
	Tracer       opentracing.Tracer
	Configurator source.Configurator
	Transformer  *config.Transformer
}

// Register server.
func Register(lc fx.Lifecycle, params RegisterParams) {
	sparams := ServerParams{
		RedisConfig:  params.RedisConfig,
		Configurator: params.Configurator,
		Transformer:  params.Transformer,
	}
	server := NewServer(sparams)

	v1.RegisterServiceServer(params.GRPCServer, server)

	var conn *grpc.ClientConn

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn, _ = sgrpc.NewClient(ctx, fmt.Sprintf("127.0.0.1:%s", params.GRPCConfig.Port),
				sgrpc.WithClientConfig(params.GRPCConfig), sgrpc.WithClientLogger(params.Logger),
				sgrpc.WithClientTracer(params.Tracer), sgrpc.WithClientDialOption(grpc.WithBlock()),
			)

			return v1.RegisterServiceHandler(ctx, params.HTTPServer.Mux, conn)
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})
}

// ServerParams for gRPC.
type ServerParams struct {
	RedisConfig  *redis.Config
	Configurator source.Configurator
	Transformer  *config.Transformer
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{conf: params.Configurator, trans: params.Transformer}
}
