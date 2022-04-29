package grpc

import (
	"context"
	"fmt"

	sgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	"github.com/alexfalkowski/go-service/version"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ClientConnParams for gRPC.
type ClientConnParams struct {
	fx.In

	Config  *sgrpc.Config
	Logger  *zap.Logger
	Tracer  opentracing.Tracer
	Version version.Version
}

// NewClientConn for gRPC.
func NewClientConn(lc fx.Lifecycle, params ClientConnParams) (*grpc.ClientConn, error) {
	p := sgrpc.ClientParams{
		Context: context.Background(), Host: fmt.Sprintf("127.0.0.1:%s", params.Config.Port),
		Version: params.Version, Config: params.Config,
	}

	conn, err := sgrpc.NewClient(p, sgrpc.WithClientLogger(params.Logger), sgrpc.WithClientTracer(params.Tracer))
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})

	return conn, err
}
