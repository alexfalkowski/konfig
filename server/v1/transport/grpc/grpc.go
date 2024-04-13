package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/http"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	v1c "github.com/alexfalkowski/konfig/client/v1/config"
	g "github.com/alexfalkowski/konfig/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
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
	Tracer       trace.Tracer
	Meter        metric.Meter
	Server       v1.ServiceServer
}

// Register server.
func Register(params RegisterParams) error {
	v1.RegisterServiceServer(params.GRPCServer.Server, params.Server)

	opts := g.ClientOpts{
		Lifecycle:    params.Lifecycle,
		ClientConfig: params.ClientConfig.Config,
		Logger:       params.Logger,
		Tracer:       params.Tracer,
		Meter:        params.Meter,
	}

	conn, err := g.NewClient(opts)
	if err != nil {
		return err
	}

	return v1.RegisterServiceHandler(context.Background(), params.HTTPServer.Mux, conn)
}
