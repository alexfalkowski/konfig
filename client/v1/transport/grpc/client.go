package grpc

import (
	"github.com/alexfalkowski/auth/client"
	t "github.com/alexfalkowski/go-service/security/token"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	v1c "github.com/alexfalkowski/konfig/client/v1/config"
	"github.com/alexfalkowski/konfig/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	TokenConfig  *t.Config
	ClientConfig *v1c.Config
	Logger       *zap.Logger
	Tracer       trace.Tracer
	Meter        metric.Meter
	Token        *client.Token
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	opts := grpc.ClientOpts{
		Lifecycle:    params.Lifecycle,
		ClientConfig: params.ClientConfig.Config,
		TokenConfig:  params.TokenConfig,
		Logger:       params.Logger,
		Tracer:       params.Tracer,
		Meter:        params.Meter,
		Token:        params.Token,
	}

	conn, err := grpc.NewClient(opts)
	if err != nil {
		return nil, err
	}

	return v1.NewServiceClient(conn), nil
}
