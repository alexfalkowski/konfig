package http

import (
	"github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/konfig/server/service"
)

type (
	// Server for HTTP.
	Server struct {
		service *service.Service
	}

	// Error for HTTP.
	Error struct {
		Message string `json:"message,omitempty"`
	}
)

// Register for HTTP.
func Register(service *service.Service) {
	s := &Server{service: service}

	http.Handler("POST /v1/config", &configErrorer{}, s.GetConfig)
	http.Handler("POST /v1/secrets", &secretsErrorer{}, s.GetSecrets)
}
