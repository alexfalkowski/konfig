package grpc

import (
	"context"

	sgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/metrics/prometheus"
	"github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	"github.com/alexfalkowski/go-service/version"
	"github.com/alexfalkowski/konfig/client"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ClientConnParams for gRPC.
type ClientConnParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    *sgrpc.Config
	Logger    *zap.Logger
	Tracer    opentracing.Tracer
	Client    *client.Config
	Version   version.Version
	Metrics   *prometheus.ClientMetrics
}

// NewClientConn for gRPC.
func NewClientConn(params ClientConnParams) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), params.Client.Timeout)
	defer cancel()

	conn, err := sgrpc.NewClient(
		sgrpc.ClientParams{Context: ctx, Host: params.Client.Host, Version: params.Version, Config: params.Config},
		sgrpc.WithClientLogger(params.Logger), sgrpc.WithClientTracer(params.Tracer), sgrpc.WithClientDialOption(grpc.WithBlock()),
		sgrpc.WithClientMetrics(params.Metrics),
	)
	if err != nil {
		return nil, err
	}

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})

	return conn, err
}
