package grpc

import (
	"context"

	ac "github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/security"
	"github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/security/token"
	"github.com/alexfalkowski/go-service/transport/grpc"
	gt "github.com/alexfalkowski/go-service/transport/grpc/security/token"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
)

// ClientOpts for gRPC.
type ClientOpts struct {
	Lifecycle    fx.Lifecycle
	ClientConfig *client.Config
	TokenConfig  *token.Config
	Logger       *zap.Logger
	Tracer       trace.Tracer
	Meter        metric.Meter
	Token        *ac.Token
}

// NewClient for gRPC.
func NewClient(options ClientOpts) (*g.ClientConn, error) {
	sec, err := grpc.WithClientTLS(options.ClientConfig.TLS)
	if err != nil {
		return nil, err
	}

	opts := []grpc.ClientOption{
		grpc.WithClientLogger(options.Logger), grpc.WithClientTracer(options.Tracer),
		grpc.WithClientMetrics(options.Meter), grpc.WithClientRetry(options.ClientConfig.Retry),
		grpc.WithClientUserAgent(options.ClientConfig.UserAgent), sec,
	}

	if security.IsAuth(options.TokenConfig) {
		kind, name := "jwt", os.ExecutableName()
		opts = append(opts, grpc.WithClientDialOption(g.WithPerRPCCredentials(gt.NewPerRPCCredentials(options.Token.Generator(kind, name)))))
	}

	conn, err := grpc.NewClient(options.ClientConfig.Host, opts...)

	options.Lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn, err
}
