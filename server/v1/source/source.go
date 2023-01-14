package source

import (
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/metrics/prometheus"
	hopentracing "github.com/alexfalkowski/go-service/transport/http/trace/opentracing"
	"github.com/alexfalkowski/konfig/server/v1/config"
	"github.com/alexfalkowski/konfig/source"
	"github.com/alexfalkowski/konfig/source/configurator"
	gopentracing "github.com/alexfalkowski/konfig/source/configurator/git/opentracing"
	sopentracing "github.com/alexfalkowski/konfig/source/configurator/s3/opentracing"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Configurator for source.
type Configurator configurator.Configurator

// ConfiguratorParams for source.
type ConfiguratorParams struct {
	fx.In

	Config     *config.Config
	HTTPConfig *http.Config
	Logger     *zap.Logger
	HTTPTracer hopentracing.Tracer
	Metrics    *prometheus.ClientMetrics
	GitTracer  gopentracing.Tracer
	S3Tracer   sopentracing.Tracer
}

// NewConfigurator for source.
func NewConfigurator(params ConfiguratorParams) (Configurator, error) {
	p := &source.ConfiguratorParams{
		Config:     &params.Config.Source,
		HTTPConfig: params.HTTPConfig,
		Logger:     params.Logger,
		HTTPTracer: params.HTTPTracer,
		Metrics:    params.Metrics,
		GitTracer:  params.GitTracer,
		S3Tracer:   params.S3Tracer,
	}

	return source.NewConfigurator(p)
}
