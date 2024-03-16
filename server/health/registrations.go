package health

import (
	"time"

	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health"
	"github.com/alexfalkowski/go-service/transport/http"
	h "github.com/alexfalkowski/konfig/health"
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
	Health *h.Config
}

// NewRegistrations for health.
func NewRegistrations(params RegistrationsParams) (health.Registrations, error) {
	d, err := time.ParseDuration(params.Health.Duration)
	if err != nil {
		return nil, err
	}

	client, err := http.NewClient(http.WithClientUserAgent(params.HTTP.UserAgent))
	if err != nil {
		return nil, err
	}

	registrations := health.Registrations{
		server.NewRegistration("noop", d, checker.NewNoopChecker()),
		server.NewRegistration("vault", d, checker.NewHTTPChecker(params.Vault.Address, client)),
	}
	s := params.Source

	if s.IsGit() {
		registrations = append(registrations, server.NewRegistration("git", d, checker.NewHTTPChecker(s.Git.URL, client)))
	}

	return registrations, nil
}
