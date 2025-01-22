package http

import (
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/konfig/api/config"
)

// Register for HTTP.
func Register(handler *Handler) {
	rpc.Route("/v1/config", handler.GetConfig)
	rpc.Route("/v1/secrets", handler.GetSecrets)
}

// NewServer for HTTP.
func NewHandler(service *config.Configuration) *Handler {
	return &Handler{service: service}
}

// Handler for HTTP.
type Handler struct {
	service *config.Configuration
}

func (h *Handler) error(err error) error {
	if err == nil {
		return nil
	}

	if config.IsInvalidArgument(err) {
		return status.Error(http.StatusBadRequest, err.Error())
	}

	if config.IsNotFound(err) {
		return status.Error(http.StatusNotFound, err.Error())
	}

	return err
}
