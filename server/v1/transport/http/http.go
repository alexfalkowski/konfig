package http

import (
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/konfig/server/service"
	"go.uber.org/fx"
)

type (
	// RegisterParams for HTTP.
	RegisterParams struct {
		fx.In

		Marshaller *marshaller.Map
		Mux        http.ServeMux
		Service    *service.Service
	}

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
func Register(params RegisterParams) error {
	s := &Server{service: params.Service}

	gh := http.NewHandler[GetConfigRequest](params.Mux, params.Marshaller, &configErrorer{})
	if err := gh.Handle("POST", "/v1/config", s.GetConfig); err != nil {
		return err
	}

	sh := http.NewHandler[GetSecretsRequest](params.Mux, params.Marshaller, &secretsErrorer{})
	if err := sh.Handle("POST", "/v1/secrets", s.GetSecrets); err != nil {
		return err
	}

	return nil
}
