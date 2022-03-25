package http

import (
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health/transport/http"
)

// NewHealthObserver for HTTP.
func NewHealthObserver(healthServer *server.Server) (*http.HealthObserver, error) {
	ob, err := healthServer.Observe("noop")
	if err != nil {
		return nil, err
	}

	return &http.HealthObserver{Observer: ob}, nil
}

// NewLivenessObserver for HTTP.
func NewLivenessObserver(healthServer *server.Server) (*http.LivenessObserver, error) {
	ob, err := healthServer.Observe("noop")
	if err != nil {
		return nil, err
	}

	return &http.LivenessObserver{Observer: ob}, nil
}

// NewReadinessObserver for HTTP.
func NewReadinessObserver(healthServer *server.Server) (*http.ReadinessObserver, error) {
	ob, err := healthServer.Observe("noop")
	if err != nil {
		return nil, err
	}

	return &http.ReadinessObserver{Observer: ob}, nil
}
