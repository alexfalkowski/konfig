package health

import (
	"time"

	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health"
	"github.com/alexfalkowski/go-service/transport/http"
	"go.uber.org/zap"
)

// NewRegistrations for health.
func NewRegistrations(httpCfg *http.Config, logger *zap.Logger) health.Registrations {
	duration := 10 * time.Millisecond // nolint:gomnd

	nc := checker.NewNoopChecker()
	nr := server.NewRegistration("noop", duration, nc)

	return health.Registrations{nr}
}
