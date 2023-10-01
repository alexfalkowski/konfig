package ssm

import (
	"context"
	"os"

	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/telemetry/metrics/prometheus"
	"github.com/alexfalkowski/go-service/transport/http/telemetry/tracer"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for SSM.
type ClientParams struct {
	fx.In

	HTTPConfig *http.Config
	Logger     *zap.Logger
	Tracer     tracer.Tracer
	Metrics    *prometheus.ClientCollector
}

// NewClient for SSM.
func NewClient(params ClientParams) (*ssm.Client, error) {
	client := http.NewClient(
		http.ClientParams{Config: params.HTTPConfig},
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Metrics),
	)
	ctx := context.Background()

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...any) (aws.Endpoint, error) {
		url := os.Getenv("AWS_URL")
		if url != "" {
			return aws.Endpoint{PartitionID: "aws", URL: url, SigningRegion: region}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	opts := []func(*config.LoadOptions) error{
		config.WithEndpointResolverWithOptions(resolver),
		config.WithHTTPClient(client),
		config.WithRetryMaxAttempts(int(params.HTTPConfig.Retry.Attempts)),
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return ssm.NewFromConfig(cfg), nil
}
