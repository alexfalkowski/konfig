package source

import (
	"errors"

	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/metrics/prometheus"
	hopentracing "github.com/alexfalkowski/go-service/transport/http/trace/opentracing"
	"github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
	gopentracing "github.com/alexfalkowski/konfig/source/configurator/git/opentracing"
	"github.com/alexfalkowski/konfig/source/configurator/s3"
	sopentracing "github.com/alexfalkowski/konfig/source/configurator/s3/opentracing"
	"go.uber.org/zap"
)

// ErrNoConfigurator is defined in the config.
var ErrNoConfigurator = errors.New("no configurator")

// ConfiguratorParams for source.
type ConfiguratorParams struct {
	Config     *Config
	HTTPConfig *http.Config
	Logger     *zap.Logger
	HTTPTracer hopentracing.Tracer
	Metrics    *prometheus.ClientMetrics
	GitTracer  gopentracing.Tracer
	S3Tracer   sopentracing.Tracer
}

// NewConfigurator for source.
func NewConfigurator(params *ConfiguratorParams) (configurator.Configurator, error) {
	client := http.NewClient(
		http.ClientParams{Config: params.HTTPConfig},
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.HTTPTracer),
		http.WithClientMetrics(params.Metrics),
	)

	var configurator configurator.Configurator

	switch {
	case params.Config.IsFolder():
		configurator = folder.NewConfigurator(params.Config.Folder)
	case params.Config.IsGit():
		configurator = git.NewConfigurator(params.Config.Git, params.GitTracer, client)
	case params.Config.IsS3():
		configurator = s3.NewConfigurator(params.Config.S3, params.S3Tracer, client)
	default:
		return nil, ErrNoConfigurator
	}

	return configurator, nil
}
