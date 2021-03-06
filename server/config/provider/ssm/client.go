package ssm

import (
	"context"

	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/metrics/prometheus"
	"github.com/alexfalkowski/go-service/transport/http/trace/opentracing"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for SSM.
type ClientParams struct {
	fx.In

	Config     *Config
	HTTPConfig *http.Config
	Logger     *zap.Logger
	Tracer     opentracing.Tracer
	Metrics    *prometheus.ClientMetrics
}

// NewClient for SSM.
func NewClient(params ClientParams) (*ssm.Client, error) {
	client := http.NewClient(
		http.ClientParams{Config: params.HTTPConfig},
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Metrics),
	)
	ctx := context.Background()
	region := params.Config.Region

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...any) (aws.Endpoint, error) {
		if params.Config.URL != "" {
			return aws.Endpoint{PartitionID: "aws", URL: params.Config.URL, SigningRegion: region}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	opts := []func(*config.LoadOptions) error{
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithHTTPClient(client),
		config.WithRetryMaxAttempts(int(params.HTTPConfig.Retry.Attempts)),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(params.Config.Access, params.Config.Secret, "")),
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return ssm.NewFromConfig(cfg), nil
}
