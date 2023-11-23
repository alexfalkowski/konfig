package vault

import (
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/telemetry/tracer"
	"github.com/hashicorp/vault/api"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for vault.
type ConfigParams struct {
	fx.In

	Config *http.Config
	Logger *zap.Logger
	Tracer tracer.Tracer
	Meter  metric.Meter
}

// NewConfig for vault.
func NewConfig(params ConfigParams) (*api.Config, error) {
	client, err := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientUserAgent(params.Config.UserAgent),
	)
	if err != nil {
		return nil, err
	}

	config := api.DefaultConfig()

	config.HttpClient = client

	return config, nil
}

// NewClient for vault.
func NewClient(cfg *api.Config) (*api.Client, error) {
	return api.NewClient(cfg)
}
