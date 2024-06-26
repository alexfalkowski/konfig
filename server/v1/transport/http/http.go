package http

import (
	"github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/konfig/server/service"
)

// Error for HTTP.
type Error struct {
	Message string `json:"message,omitempty"`
}

// Register for HTTP.
func Register(service *service.Service) {
	http.Handle("/v1/config", &configHandler{service: service})
	http.Handle("/v1/secrets", &secretsHandler{service: service})
}