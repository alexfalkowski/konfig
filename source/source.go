package source

import (
	"errors"

	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/konfig/source/configurator"
	"github.com/alexfalkowski/konfig/source/configurator/folder"
	"github.com/alexfalkowski/konfig/source/configurator/git"
	"github.com/alexfalkowski/konfig/source/configurator/s3"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
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
	Tracer     trace.Tracer
	Meter      metric.Meter
}

// NewConfigurator for source.
func NewConfigurator(params ConfiguratorParams) (configurator.Configurator, error) {
	client := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientUserAgent(params.HTTPConfig.UserAgent),
	)

	var configurator configurator.Configurator

	switch {
	case params.Config.IsFolder():
		configurator = folder.NewConfigurator(params.Config.Folder)
	case params.Config.IsGit():
		configurator = git.NewConfigurator(params.Config.Git, params.Tracer, client)
	case params.Config.IsS3():
		configurator = s3.NewConfigurator(params.Config.S3, params.Tracer, client)
	default:
		return nil, ErrNoConfigurator
	}

	return configurator, nil
}
