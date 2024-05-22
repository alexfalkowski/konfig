package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/transport/grpc"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	v1c "github.com/alexfalkowski/konfig/client/v1/config"
	g "github.com/alexfalkowski/konfig/transport/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	GRPC      *grpc.Server
	Mux       *runtime.ServeMux
	Client    *v1c.Config
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	Server    v1.ServiceServer
}

// Register server.
func Register(params RegisterParams) error {
	v1.RegisterServiceServer(params.GRPC.Server(), params.Server)

	opts := g.ClientOpts{
		Lifecycle: params.Lifecycle,
		Client:    params.Client.Config,
		Logger:    params.Logger,
		Tracer:    params.Tracer,
		Meter:     params.Meter,
	}

	conn, err := g.NewClient(opts)
	if err != nil {
		return err
	}

	return v1.RegisterServiceHandler(context.Background(), params.Mux, conn)
}
