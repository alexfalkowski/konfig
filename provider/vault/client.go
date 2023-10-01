package vault

import (
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/telemetry/metrics/prometheus"
	"github.com/alexfalkowski/go-service/transport/http/telemetry/tracer"
	"github.com/hashicorp/vault/api"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for vault.
type ConfigParams struct {
	fx.In

	Config  *http.Config
	Logger  *zap.Logger
	Tracer  tracer.Tracer
	Metrics *prometheus.ClientCollector
}

// NewConfig for vault.
func NewConfig(params ConfigParams) *api.Config {
	client := http.NewClient(
		http.ClientParams{Config: params.Config},
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Metrics),
	)
	config := api.DefaultConfig()

	config.HttpClient = client

	return config
}

// NewClient for vault.
func NewClient(cfg *api.Config) (*api.Client, error) {
	return api.NewClient(cfg)
}
