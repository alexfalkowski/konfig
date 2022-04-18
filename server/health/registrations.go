package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health"
	sopentracing "github.com/alexfalkowski/go-service/trace/opentracing"
	"github.com/alexfalkowski/go-service/transport/http"
	khealth "github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/source"
	"github.com/hashicorp/vault/api"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Params for health.
type Params struct {
	fx.In

	HTTP   *http.Config
	Source *source.Config
	Vault  *api.Config
	Logger *zap.Logger
	Tracer sopentracing.TransportTracer
	Health *khealth.Config
}

// NewRegistrations for health.
func NewRegistrations(params Params) health.Registrations {
	client := http.NewClient(http.WithClientConfig(params.HTTP), http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer))
	registrations := health.Registrations{
		server.NewRegistration("noop", params.Health.Duration, checker.NewNoopChecker()),
		server.NewRegistration("vault", params.Health.Duration, checker.NewHTTPChecker(params.Vault.Address, client)),
	}

	if params.Source.IsGit() {
		registrations = append(registrations, server.NewRegistration("git", params.Health.Duration, checker.NewHTTPChecker(params.Source.Git.URL, client)))
	}

	return registrations
}
