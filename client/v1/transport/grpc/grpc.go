package grpc

import (
	"context"

	sopentracing "github.com/alexfalkowski/go-service/trace/opentracing"
	tgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/client"
	kzap "github.com/alexfalkowski/konfig/client/v1/transport/grpc/logger/zap"
	"github.com/alexfalkowski/konfig/client/v1/transport/grpc/task"
	"github.com/alexfalkowski/konfig/client/v1/transport/grpc/trace/opentracing"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Config *tgrpc.Config
	Logger *zap.Logger
	Tracer sopentracing.TransportTracer
	Client *client.Config
}

// Register client.
func Register(lc fx.Lifecycle, params RegisterParams) {
	var (
		conn *grpc.ClientConn
		err  error
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn, err = tgrpc.NewClient(ctx, params.Client.Host,
				tgrpc.WithClientConfig(params.Config), tgrpc.WithClientLogger(params.Logger),
				tgrpc.WithClientTracer(params.Tracer),
			)
			if err != nil {
				return err
			}

			client := NewClient(v1.NewConfiguratorServiceClient(conn), params.Client, params.Logger)

			return client.Perform(ctx)
		},
		OnStop: func(ctx context.Context) error {
			if conn != nil {
				return conn.Close()
			}

			return nil
		},
	})
}

// NewClient for gRPC.
func NewClient(client v1.ConfiguratorServiceClient, cfg *client.Config, logger *zap.Logger) task.Client {
	var clt task.Client = &clt{client: client, cfg: cfg}
	clt = kzap.NewClient(logger, cfg, clt)
	clt = opentracing.NewClient(cfg, clt)

	return clt
}
