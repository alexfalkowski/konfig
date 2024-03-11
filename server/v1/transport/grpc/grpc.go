package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	"github.com/alexfalkowski/go-service/transport/http"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	v1c "github.com/alexfalkowski/konfig/client/v1/config"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	GRPCServer   *grpc.Server
	HTTPServer   *http.Server
	ClientConfig *v1c.Config
	Logger       *zap.Logger
	Tracer       tracer.Tracer
	Meter        metric.Meter
	Server       v1.ServiceServer
}

// Register server.
func Register(params RegisterParams) error {
	ctx := context.Background()

	v1.RegisterServiceServer(params.GRPCServer.Server, params.Server)

	opts := []grpc.ClientOption{
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer),
		grpc.WithClientMetrics(params.Meter), grpc.WithClientRetry(&params.ClientConfig.Retry),
		grpc.WithClientUserAgent(params.ClientConfig.UserAgent),
	}

	if params.ClientConfig.Security.Enabled {
		sec, err := grpc.WithClientSecure(params.ClientConfig.Security)
		if err != nil {
			return err
		}

		opts = append(opts, sec)
	}

	conn, err := grpc.NewClient(ctx, params.ClientConfig.Host, opts...)
	if err != nil {
		return err
	}

	if err := v1.RegisterServiceHandler(ctx, params.HTTPServer.Mux, conn); err != nil {
		return err
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			conn.ResetConnectBackoff()

			return nil
		},
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return nil
}
