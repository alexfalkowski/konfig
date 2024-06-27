package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/health"
	"github.com/alexfalkowski/go-service/time"
	"github.com/alexfalkowski/go-service/transport/http"
	h "github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/source"
	"github.com/hashicorp/vault/api"
	"go.uber.org/fx"
)

// RegistrationsParams for health.
type RegistrationsParams struct {
	fx.In

	HTTP      *http.Config
	Source    *source.Config
	Vault     *api.Config
	Health    *h.Config
	UserAgent env.UserAgent
}

// NewRegistrations for health.
func NewRegistrations(params RegistrationsParams) health.Registrations {
	client := http.NewClient(http.WithClientUserAgent(params.UserAgent))
	d := time.MustParseDuration(params.Health.Duration)
	registrations := health.Registrations{
		server.NewRegistration("noop", d, checker.NewNoopChecker()),
		server.NewRegistration("vault", d, checker.NewHTTPChecker(params.Vault.Address, client)),
	}

	return registrations
}
