package source

import (
	"errors"

	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/metrics/prometheus"
	hotel "github.com/alexfalkowski/go-service/transport/http/otel"
	"github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
	gotel "github.com/alexfalkowski/konfig/source/configurator/git/otel"
	"github.com/alexfalkowski/konfig/source/configurator/s3"
	sotel "github.com/alexfalkowski/konfig/source/configurator/s3/otel"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ErrNoConfigurator is defined in the config.
var ErrNoConfigurator = errors.New("no configurator")

// ConfiguratorParams for source.
type ConfiguratorParams struct {
	fx.In

	Config     *Config
	HTTPConfig *http.Config
	Logger     *zap.Logger
	HTTPTracer hotel.Tracer
	Metrics    *prometheus.ClientCollector
	GitTracer  gotel.Tracer
	S3Tracer   sotel.Tracer
}

// NewConfigurator for source.
func NewConfigurator(params ConfiguratorParams) (configurator.Configurator, error) {
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
