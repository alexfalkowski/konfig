package ssm

import (
	"context"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/id"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/konfig/internal/aws/endpoint"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// ConfigParams for SSM.
type ClientParams struct {
	fx.In
	Tracer    trace.Tracer
	Meter     metric.Meter
	ID        id.Generator
	Endpoint  endpoint.Endpoint
	Config    *http.Config
	Logger    *logger.Logger
	UserAgent env.UserAgent
}

// NewClient for SSM.
func NewClient(params ClientParams) (*ssm.Client, error) {
	client, _ := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientUserAgent(params.UserAgent),
		http.WithClientTimeout(params.Config.Timeout), http.WithClientID(params.ID),
	)

	ctx := context.Background()
	opts := []func(*config.LoadOptions) error{
		config.WithHTTPClient(client),
		config.WithRetryMaxAttempts(int(params.Config.Retry.Attempts)), //nolint:gosec
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	cl := ssm.NewFromConfig(cfg, func(o *ssm.Options) {
		if params.Endpoint.IsSet() {
			o.BaseEndpoint = aws.String(string(params.Endpoint))
		}
	})

	return cl, err
}
