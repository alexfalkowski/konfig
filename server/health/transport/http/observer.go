package http

import (
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health/transport/http"
)

// NewHealthObserver for HTTP.
func NewHealthObserver(healthServer *server.Server) *http.HealthObserver {
	return &http.HealthObserver{Observer: healthServer.Observe("noop")}
}

// NewLivenessObserver for HTTP.
func NewLivenessObserver(healthServer *server.Server) *http.LivenessObserver {
	return &http.LivenessObserver{Observer: healthServer.Observe("noop")}
}

// NewReadinessObserver for HTTP.
func NewReadinessObserver(healthServer *server.Server) *http.ReadinessObserver {
	return &http.ReadinessObserver{Observer: healthServer.Observe("noop")}
}
