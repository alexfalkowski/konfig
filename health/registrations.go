package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health"
	schecker "github.com/alexfalkowski/go-service/health/checker"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/konfig/vcs/git"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Params for health.
type Params struct {
	fx.In

	HTTP   *http.Config
	Git    *git.Config
	Redis  *redis.Ring
	Logger *zap.Logger
	Health *Config
}

// NewRegistrations for health.
func NewRegistrations(params Params) health.Registrations {
	client := http.NewClient(params.HTTP, params.Logger)
	registrations := health.Registrations{
		server.NewRegistration("noop", params.Health.Duration, checker.NewNoopChecker()),
		server.NewRegistration("git", params.Health.Duration, checker.NewHTTPChecker(params.Git.URL, client)),
		server.NewRegistration("redis", params.Health.Duration, schecker.NewRedisChecker(params.Redis, params.Health.Timeout)),
	}

	return registrations
}
