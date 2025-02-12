package http

import (
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health/transport/http"
	aws "github.com/alexfalkowski/konfig/internal/aws/endpoint"
	"github.com/alexfalkowski/konfig/internal/source"
)

// NewHealthObserver for HTTP.
func NewHealthObserver(healthServer *server.Server, config *source.Config, endpoint aws.Endpoint) *http.HealthObserver {
	names := []string{"vault"}
	if source.IsEnabled(config) && config.IsGit() {
		names = append(names, "git")
	}

	if endpoint.IsSet() {
		names = append(names, "aws")
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
