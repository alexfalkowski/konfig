package ssm

import (
	"context"
	"os"

	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/telemetry/tracer"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for SSM.
type ClientParams struct {
	fx.In

	Config *http.Config
	Logger *zap.Logger
	Tracer tracer.Tracer
	Meter  metric.Meter
}

// NewClient for SSM.
func NewClient(params ClientParams) (*ssm.Client, error) {
	client, err := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientUserAgent(params.Config.UserAgent),
	)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	resolver := aws.EndpointResolverWithOptionsFunc(func(_, region string, _ ...any) (aws.Endpoint, error) {
		url := os.Getenv("AWS_URL")
		if url != "" {
			return aws.Endpoint{PartitionID: "aws", URL: url, SigningRegion: region}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	opts := []func(*config.LoadOptions) error{
		config.WithEndpointResolverWithOptions(resolver),
		config.WithHTTPClient(client),
		config.WithRetryMaxAttempts(int(params.Config.Retry.Attempts)),
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return ssm.NewFromConfig(cfg), nil
}
