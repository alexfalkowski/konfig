package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/health"
	"github.com/alexfalkowski/go-service/time"
	"github.com/alexfalkowski/go-service/transport/http"
	aws "github.com/alexfalkowski/konfig/internal/aws/endpoint"
	"github.com/alexfalkowski/konfig/internal/source"
	"github.com/hashicorp/vault/api"
	"go.uber.org/fx"
)

// RegistrationsParams for health.
type RegistrationsParams struct {
	fx.In

	HTTP      *http.Config
	Source    *source.Config
	Vault     *api.Config
	Health    *Config
	Endpoint  aws.Endpoint
	UserAgent env.UserAgent
}

// NewRegistrations for health.
func NewRegistrations(params RegistrationsParams) health.Registrations {
	rt, _ := http.NewRoundTripper(http.WithClientUserAgent(params.UserAgent))
	t := time.MustParseDuration(params.Health.Timeout)
	d := time.MustParseDuration(params.Health.Duration)
	registrations := health.Registrations{
		server.NewRegistration("noop", d, checker.NewNoopChecker()),
	}

	if params.Endpoint.IsSet() {
		registrations = append(registrations, server.NewRegistration("aws", d, checker.NewHTTPChecker(params.Endpoint.String(), rt, t)))
	}

	if params.Vault != nil {
		registrations = append(registrations, server.NewRegistration("vault", d, checker.NewHTTPChecker(params.Vault.Address, rt, t)))
	}

	return registrations
}
