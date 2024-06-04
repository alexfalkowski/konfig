package ssm

import (
	"context"

	"github.com/alexfalkowski/go-service/env"
	sh "github.com/alexfalkowski/go-service/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for SSM.
type ClientParams struct {
	fx.In

	Config    *sh.Config
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	UserAgent env.UserAgent
}

// NewClient for SSM.
func NewClient(params ClientParams) (*ssm.Client, error) {
	client := sh.NewClient(
		sh.WithClientLogger(params.Logger), sh.WithClientTracer(params.Tracer),
		sh.WithClientMetrics(params.Meter), sh.WithClientUserAgent(string(params.UserAgent)),
	)

	ctx := context.Background()
	opts := []func(*config.LoadOptions) error{
		config.WithHTTPClient(client),
		config.WithRetryMaxAttempts(int(params.Config.Retry.Attempts)),
	}

	r := &resolver{EndpointResolverV2: ssm.NewDefaultEndpointResolverV2()}
	cfg, err := config.LoadDefaultConfig(ctx, opts...)

	return ssm.NewFromConfig(cfg, ssm.WithEndpointResolverV2(r)), err
}
