package grpc

import (
	"context"

	sgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/client"
	kzap "github.com/alexfalkowski/konfig/client/v1/transport/grpc/logger/zap"
	"github.com/alexfalkowski/konfig/client/v1/transport/grpc/task"
	gopentracing "github.com/alexfalkowski/konfig/client/v1/transport/grpc/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Config *sgrpc.Config
	Logger *zap.Logger
	Tracer opentracing.Tracer
	Client *client.Config
}

// Register client.
func Register(lc fx.Lifecycle, params RegisterParams) {
	var conn *grpc.ClientConn

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn, _ = sgrpc.NewClient(ctx, params.Client.Host,
				sgrpc.WithClientConfig(params.Config), sgrpc.WithClientLogger(params.Logger),
				sgrpc.WithClientTracer(params.Tracer), sgrpc.WithClientDialOption(grpc.WithBlock()),
			)

			client := NewClient(v1.NewServiceClient(conn), params.Client, params.Logger)

			return client.Perform(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})
}

// NewClient for gRPC.
func NewClient(client v1.ServiceClient, cfg *client.Config, logger *zap.Logger) task.Client {
	var clt task.Client = &clt{client: client, cfg: cfg}
	clt = kzap.NewClient(logger, cfg, clt)
	clt = gopentracing.NewClient(cfg, clt)

	return clt
}
