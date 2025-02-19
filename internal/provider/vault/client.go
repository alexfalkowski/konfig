package vault

import (
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/id"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/hashicorp/vault/api"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// ConfigParams for vault.
type ConfigParams struct {
	fx.In
	Tracer    trace.Tracer
	Meter     metric.Meter
	ID        id.Generator
	Config    *http.Config
	Logger    *logger.Logger
	UserAgent env.UserAgent
}

// NewConfig for vault.
func NewConfig(params ConfigParams) *api.Config {
	client, _ := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientUserAgent(params.UserAgent),
		http.WithClientTimeout(params.Config.Timeout), http.WithClientID(params.ID),
	)

	config := api.DefaultConfig()
	config.HttpClient = client

	return config
}

// NewClient for vault.
func NewClient(cfg *api.Config) (*api.Client, error) {
	return api.NewClient(cfg)
}
