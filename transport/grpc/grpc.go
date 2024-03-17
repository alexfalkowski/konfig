package grpc

import (
	"context"

	ac "github.com/alexfalkowski/auth/client"
	c "github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/security/token"
	"github.com/alexfalkowski/go-service/transport/grpc"
	gt "github.com/alexfalkowski/go-service/transport/grpc/security/token"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
)

// ClientOpts for gRPC.
type ClientOpts struct {
	Lifecycle    fx.Lifecycle
	ClientConfig c.Config
	TokenConfig  *token.Config
	Logger       *zap.Logger
	Tracer       tracer.Tracer
	Meter        metric.Meter
	Token        *ac.Token
}

// NewClient for gRPC.
func NewClient(ctx context.Context, options ClientOpts) (*g.ClientConn, error) {
	opts := []grpc.ClientOption{
		grpc.WithClientLogger(options.Logger), grpc.WithClientTracer(options.Tracer),
		grpc.WithClientMetrics(options.Meter), grpc.WithClientRetry(options.ClientConfig.Retry),
		grpc.WithClientUserAgent(options.ClientConfig.UserAgent),
	}

	if IsAuth(options.TokenConfig) {
		opts = append(opts, grpc.WithClientDialOption(g.WithPerRPCCredentials(gt.NewPerRPCCredentials(options.Token.Generator("jwt", "konfig")))))
	}

	sec := options.ClientConfig.Security
	if sec != nil && sec.Enabled {
		sec, err := grpc.WithClientSecure(options.ClientConfig.Security)
		if err != nil {
			return nil, err
		}

		opts = append(opts, sec)
	}

	conn, err := grpc.NewClient(ctx, options.ClientConfig.Host, opts...)
	if err != nil {
		return nil, err
	}

	options.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			conn.ResetConnectBackoff()

			return nil
		},
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn, nil
}
