package http

import (
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health/transport/http"
	v1 "github.com/alexfalkowski/konfig/server/v1/config"
)

// NewHealthObserver for HTTP.
func NewHealthObserver(healthServer *server.Server, v1Config *v1.Config) *http.HealthObserver {
	names := []string{"vault"}
	if v1Config.Source.IsGit() {
		names = append(names, "v1-git")
	}

	return &http.HealthObserver{Observer: healthServer.Observe(names...)}
}

// NewLivenessObserver for HTTP.
func NewLivenessObserver(healthServer *server.Server) *http.LivenessObserver {
	return &http.LivenessObserver{Observer: healthServer.Observe("noop")}
}

// NewReadinessObserver for HTTP.
func NewReadinessObserver(healthServer *server.Server) *http.ReadinessObserver {
	return &http.ReadinessObserver{Observer: healthServer.Observe("noop")}
}
