package grpc

import (
	"context"

	"github.com/alexfalkowski/auth/client"
	t "github.com/alexfalkowski/go-service/security/token"
	"github.com/alexfalkowski/go-service/transport/grpc"
	gt "github.com/alexfalkowski/go-service/transport/grpc/security/token"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	v1c "github.com/alexfalkowski/konfig/client/v1/config"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	GRPCConfig   *grpc.Config
	TokenConfig  *t.Config
	ClientConfig *v1c.Config
	Logger       *zap.Logger
	Tracer       tracer.Tracer
	Meter        metric.Meter
	Token        *client.Token
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	opts := []grpc.ClientOption{
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer), grpc.WithClientMetrics(params.Meter),
		grpc.WithClientRetry(&params.GRPCConfig.Retry), grpc.WithClientUserAgent(params.GRPCConfig.UserAgent),
	}

	if params.TokenConfig.Kind == "auth" {
		opts = append(opts, grpc.WithClientDialOption(g.WithPerRPCCredentials(gt.NewPerRPCCredentials(params.Token.Generator("jwt", "konfig")))))
	}

	conn, err := grpc.NewClient(context.Background(), params.ClientConfig.Host, opts...)
	if err != nil {
		return nil, err
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

	return v1.NewServiceClient(conn), nil
}
