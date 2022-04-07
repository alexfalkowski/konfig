package health

import (
	"time"

	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health"
	schecker "github.com/alexfalkowski/go-service/health/checker"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/konfig/vcs"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Params for health.
type Params struct {
	fx.In

	HTTP   *http.Config
	VCS    *vcs.Config
	Redis  *redis.Ring
	Logger *zap.Logger
}

// NewRegistrations for health.
func NewRegistrations(params Params) health.Registrations {
	duration := 1 * time.Minute
	client := http.NewClient(params.HTTP, params.Logger)
	registrations := health.Registrations{
		server.NewRegistration("noop", duration, checker.NewNoopChecker()),
		server.NewRegistration("git", duration, checker.NewHTTPChecker(params.VCS.Git.URL, client)),
		server.NewRegistration("redis", duration, schecker.NewRedisChecker(params.Redis, 1*time.Second)),
	}

	return registrations
}
