package http

import (
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/konfig/server/config"
)

// Register for HTTP.
func Register(service *config.Configuration) {
	ch := &configHandler{service: service}
	rpc.Unary("/v1/config", ch.GetConfig)

	sh := &secretsHandler{service: service}
	rpc.Unary("/v1/secrets", sh.GetSecrets)
}

func handleError(err error) error {
	if config.IsInvalidArgument(err) {
		return rpc.Error(http.StatusBadRequest, err.Error())
	}

	if config.IsNotFound(err) {
		return rpc.Error(http.StatusNotFound, err.Error())
	}

	return err
}
