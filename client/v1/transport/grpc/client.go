package grpc

import (
	"context"

	sgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/metrics/prometheus"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/client/v1/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    *sgrpc.Config
	Logger    *zap.Logger
	Tracer    tracer.Tracer
	Client    *config.Config
	Metrics   *prometheus.ClientCollector
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), params.Client.Timeout)
	defer cancel()

	conn, err := sgrpc.NewClient(
		sgrpc.ClientParams{Context: ctx, Host: params.Client.Host, Config: params.Config},
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

	return v1.NewServiceClient(conn), nil
}
