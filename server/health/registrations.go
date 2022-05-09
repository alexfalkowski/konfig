package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/http/metrics/prometheus"
	"github.com/alexfalkowski/go-service/transport/http/trace/opentracing"
	"github.com/alexfalkowski/go-service/version"
	khealth "github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/source"
	"github.com/hashicorp/vault/api"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegistrationsParams for health.
type RegistrationsParams struct {
	fx.In

	HTTP    *http.Config
	Source  *source.Config
	Vault   *api.Config
	Logger  *zap.Logger
	Tracer  opentracing.Tracer
	Health  *khealth.Config
	Version version.Version
	Metrics *prometheus.ClientMetrics
}

// NewRegistrations for health.
func NewRegistrations(params RegistrationsParams) health.Registrations {
	client := http.NewClient(
		http.ClientParams{Version: params.Version, Config: params.HTTP},
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Metrics),
	)
	registrations := health.Registrations{
		server.NewRegistration("noop", params.Health.Duration, checker.NewNoopChecker()),
		server.NewRegistration("vault", params.Health.Duration, checker.NewHTTPChecker(params.Vault.Address, client)),
	}

	if params.Source.IsGit() {
		registrations = append(registrations, server.NewRegistration("git", params.Health.Duration, checker.NewHTTPChecker(params.Source.Git.URL, client)))
	}

	return registrations
}
