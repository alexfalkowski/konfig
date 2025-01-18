package s3

import (
	"context"
	"os"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	endpoints "github.com/aws/smithy-go/endpoints"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for S3.
type ClientParams struct {
	fx.In

	Config    *http.Config
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	UserAgent env.UserAgent
}

// NewClient for S3.
func NewClient(params ClientParams) (*s3.Client, error) {
	client, _ := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientUserAgent(params.UserAgent),
		http.WithClientTimeout(params.Config.Timeout),
	)

	ctx := context.Background()
	opts := []func(*config.LoadOptions) error{
		config.WithHTTPClient(client),
		config.WithRetryMaxAttempts(int(params.Config.Retry.Attempts)), //nolint:gosec
	}

	r := &resolver{EndpointResolverV2: s3.NewDefaultEndpointResolverV2()}
	cfg, err := config.LoadDefaultConfig(ctx, opts...)

	cl := s3.NewFromConfig(cfg, s3.WithEndpointResolverV2(r), func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return cl, err
}

type resolver struct {
	s3.EndpointResolverV2
}

func (r *resolver) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (endpoints.Endpoint, error) {
	u := os.Getenv("AWS_URL")
	if u != "" {
		params.Endpoint = &u
	}

	return r.EndpointResolverV2.ResolveEndpoint(ctx, params)
}
