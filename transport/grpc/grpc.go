package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/security/token"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
)

// ClientOpts for gRPC.
type ClientOpts struct {
	Lifecycle fx.Lifecycle
	Client    *client.Config
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	Generator token.Generator
	UserAgent env.UserAgent
}

// NewClient for gRPC.
func NewClient(options ClientOpts) (*g.ClientConn, error) {
	sec, err := grpc.WithClientTLS(options.Client.TLS)
	if err != nil {
		return nil, err
	}

	opts := []grpc.ClientOption{
		grpc.WithClientLogger(options.Logger), grpc.WithClientTracer(options.Tracer),
		grpc.WithClientMetrics(options.Meter), grpc.WithClientRetry(options.Client.Retry),
		grpc.WithClientUserAgent(options.UserAgent), grpc.WithClientTimeout(options.Client.Timeout),
		grpc.WithClientTokenGenerator(options.Generator), sec,
	}

	conn, err := grpc.NewClient(options.Client.Host, opts...)

	options.Lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn, err
}
