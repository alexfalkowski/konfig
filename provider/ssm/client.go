package ssm

import (
	"context"

	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/konfig/aws"
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

	Config *http.Config
	Logger *zap.Logger
	Tracer trace.Tracer
	Meter  metric.Meter
}

// NewClient for SSM.
func NewClient(params ClientParams) (*ssm.Client, error) {
	client := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientUserAgent(params.Config.UserAgent),
	)

	ctx := context.Background()
	opts := []func(*config.LoadOptions) error{
		config.WithEndpointResolverWithOptions(aws.EndpointResolver()),
		config.WithHTTPClient(client),
		config.WithRetryMaxAttempts(int(params.Config.Retry.Attempts)),
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return ssm.NewFromConfig(cfg), nil
}
