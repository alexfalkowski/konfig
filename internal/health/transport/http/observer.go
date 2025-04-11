package http

import (
	health "github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health/transport/http"
	aws "github.com/alexfalkowski/konfig/internal/aws/endpoint"
	"github.com/alexfalkowski/konfig/internal/source"
	"github.com/hashicorp/vault/api"
)

// NewHealthObserver for HTTP.
func NewHealthObserver(server *health.Server, config *source.Config, endpoint aws.Endpoint, client *api.Client) *http.HealthObserver {
	names := []string{"noop"}

	if source.IsEnabled(config) && config.IsGit() {
		names = append(names, "git")
	}

	if endpoint.IsSet() {
		names = append(names, "aws")
	}

	if client != nil {
		names = append(names, "vault")
	}

	return &http.HealthObserver{Observer: server.Observe(names...)}
}

// NewLivenessObserver for HTTP.
func NewLivenessObserver(server *health.Server) *http.LivenessObserver {
	return &http.LivenessObserver{Observer: server.Observe("noop")}
}

// NewReadinessObserver for HTTP.
func NewReadinessObserver(server *health.Server) *http.ReadinessObserver {
	return &http.ReadinessObserver{Observer: server.Observe("noop")}
}
