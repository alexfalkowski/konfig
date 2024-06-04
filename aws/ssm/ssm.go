package ssm

import (
	"context"
	"os"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	endpoints "github.com/aws/smithy-go/endpoints"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for SSM.
type ClientParams struct {
	fx.In

	Config    *http.Config
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	UserAgent env.UserAgent
}

// NewClient for SSM.
func NewClient(params ClientParams) (*ssm.Client, error) {
	client := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientUserAgent(string(params.UserAgent)),
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

type resolver struct {
	ssm.EndpointResolverV2
}

func (r *resolver) ResolveEndpoint(ctx context.Context, params ssm.EndpointParameters) (endpoints.Endpoint, error) {
	u := os.Getenv("AWS_URL")
	if u != "" {
		params.Endpoint = &u
	}

	return r.EndpointResolverV2.ResolveEndpoint(ctx, params)
}
