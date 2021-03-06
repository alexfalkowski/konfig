package grpc

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/metrics/prometheus"
	"github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	"github.com/alexfalkowski/go-service/transport/http"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/server/config"
	source "github.com/alexfalkowski/konfig/source/configurator"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle       fx.Lifecycle
	GRPCServer      *grpc.Server
	HTTPServer      *http.Server
	GRPCConfig      *grpc.Config
	TransportConfig *transport.Config
	Logger          *zap.Logger
	Tracer          opentracing.Tracer
	Configurator    source.Configurator
	Transformer     *config.Transformer
	Metrics         *prometheus.ClientMetrics
}

// Register server.
func Register(params RegisterParams) error {
	ctx := context.Background()
	server := NewServer(ServerParams{Configurator: params.Configurator, Transformer: params.Transformer})

	v1.RegisterServiceServer(params.GRPCServer.Server, server)

	conn, err := grpc.NewClient(
		grpc.ClientParams{Context: ctx, Host: fmt.Sprintf("127.0.0.1:%s", params.TransportConfig.Port), Config: params.GRPCConfig},
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer), grpc.WithClientMetrics(params.Metrics),
	)
	if err != nil {
		return err
	}

	if err := v1.RegisterServiceHandler(ctx, params.HTTPServer.Mux, conn); err != nil {
		return err
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn.ResetConnectBackoff()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})

	return nil
}
