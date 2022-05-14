package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health"
	"github.com/alexfalkowski/go-service/transport/http"
	khealth "github.com/alexfalkowski/konfig/health"
	"github.com/alexfalkowski/konfig/source"
	"github.com/hashicorp/vault/api"
	"go.uber.org/fx"
)

// RegistrationsParams for health.
type RegistrationsParams struct {
	fx.In

	HTTP   *http.Config
	Source *source.Config
	Vault  *api.Config
	Health *khealth.Config
}

// NewRegistrations for health.
func NewRegistrations(params RegistrationsParams) health.Registrations {
	client := http.NewClient(http.ClientParams{Config: params.HTTP})
	registrations := health.Registrations{
		server.NewRegistration("noop", params.Health.Duration, checker.NewNoopChecker()),
		server.NewRegistration("vault", params.Health.Duration, checker.NewHTTPChecker(params.Vault.Address, client)),
	}

	if params.Source.IsGit() {
		registrations = append(registrations, server.NewRegistration("git", params.Health.Duration, checker.NewHTTPChecker(params.Source.Git.URL, client)))
	}

	return registrations
}
