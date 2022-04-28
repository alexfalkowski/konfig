package grpc

import (
	"context"

	sgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	"github.com/alexfalkowski/go-service/version"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/client/task"
	kzap "github.com/alexfalkowski/konfig/client/v1/transport/grpc/logger/zap"
	gopentracing "github.com/alexfalkowski/konfig/client/v1/transport/grpc/trace/opentracing"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ClientConnParams for gRPC.
type ClientConnParams struct {
	fx.In

	Config  *sgrpc.Config
	Logger  *zap.Logger
	Tracer  opentracing.Tracer
	Client  *client.Config
	Version version.Version
}

// NewClientConn for gRPC.
func NewClientConn(lc fx.Lifecycle, params ClientConnParams) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), params.Client.Timeout)
	defer cancel()

	conn, err := sgrpc.NewClient(
		sgrpc.ClientParams{Context: ctx, Host: params.Client.Host, Version: params.Version, Config: params.Config},
		sgrpc.WithClientLogger(params.Logger), sgrpc.WithClientTracer(params.Tracer), sgrpc.WithClientDialOption(grpc.WithBlock()),
	)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})

	return conn, err
}

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Conn   *grpc.ClientConn
	Logger *zap.Logger
	Tracer opentracing.Tracer
	Client *client.Config
}

// Register client.
func Register(lc fx.Lifecycle, params RegisterParams) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			client := NewClient(v1.NewServiceClient(params.Conn), params.Client, params.Tracer, params.Logger)

			return client.Perform(ctx)
		},
	})
}

// NewClient for gRPC.
func NewClient(client v1.ServiceClient, cfg *client.Config, tracer opentracing.Tracer, logger *zap.Logger) task.Task {
	var clt task.Task = &Client{client: client, cfg: cfg}
	clt = kzap.NewClient(logger, cfg, clt)
	clt = gopentracing.NewClient(cfg, tracer, clt)

	return clt
}
