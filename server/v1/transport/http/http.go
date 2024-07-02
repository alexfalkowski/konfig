package http

import (
	"net/http"

	nh "github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/konfig/server/service"
)

// Register for HTTP.
func Register(service *service.Service) {
	nh.Handle("/v1/config", &configHandler{service: service})
	nh.Handle("/v1/secrets", &secretsHandler{service: service})
}

func handleError(err error) error {
	if service.IsInvalidArgument(err) {
		return nh.Error(http.StatusBadRequest, err.Error())
	}

	if service.IsNotFound(err) {
		return nh.Error(http.StatusNotFound, err.Error())
	}

	return err
}
