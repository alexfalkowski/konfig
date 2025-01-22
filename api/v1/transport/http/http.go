package http

import (
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/konfig/api/config"
)

// Register for HTTP.
func Register(service *config.Configuration) {
	ch := &configHandler{service: service}
	rpc.Route("/v1/config", ch.GetConfig)

	sh := &secretsHandler{service: service}
	rpc.Route("/v1/secrets", sh.GetSecrets)
}

func handleError(err error) error {
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
