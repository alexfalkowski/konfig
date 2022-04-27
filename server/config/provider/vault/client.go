package vault

import (
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/trace/opentracing"
	"github.com/hashicorp/vault/api"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for vault.
type ConfigParams struct {
	fx.In

	Config *http.Config
	Logger *zap.Logger
	Tracer opentracing.Tracer
}

// NewConfig for vault.
func NewConfig(params ConfigParams) *api.Config {
	client := http.NewClient(http.WithClientConfig(params.Config), http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer))
	config := api.DefaultConfig()

	config.HttpClient = client

	return config
}

// NewClient for vault.
func NewClient(cfg *api.Config) (*api.Client, error) {
	return api.NewClient(cfg)
}
