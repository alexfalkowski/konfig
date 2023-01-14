package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health"
	"github.com/alexfalkowski/go-service/transport/http"
	khealth "github.com/alexfalkowski/konfig/health"
	v1 "github.com/alexfalkowski/konfig/server/v1/config"
	"github.com/hashicorp/vault/api"
	"go.uber.org/fx"
)

// RegistrationsParams for health.
type RegistrationsParams struct {
	fx.In

	HTTP     *http.Config
	V1Config *v1.Config
	Vault    *api.Config
	Health   *khealth.Config
}

// NewRegistrations for health.
func NewRegistrations(params RegistrationsParams) health.Registrations {
	client := http.NewClient(http.ClientParams{Config: params.HTTP})
	registrations := health.Registrations{
		server.NewRegistration("noop", params.Health.Duration, checker.NewNoopChecker()),
		server.NewRegistration("vault", params.Health.Duration, checker.NewHTTPChecker(params.Vault.Address, client)),
	}
	v1s := params.V1Config.Source

	if v1s.IsGit() {
		registrations = append(registrations, server.NewRegistration("v1-git", params.Health.Duration, checker.NewHTTPChecker(v1s.Git.URL, client)))
	}

	return registrations
}
