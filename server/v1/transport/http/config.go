package http

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/server/service"
)

type (
	// GetConfigRequest for a specific application.
	GetConfigRequest struct {
		Application string `json:"application,omitempty"`
		Version     string `json:"version,omitempty"`
		Environment string `json:"environment,omitempty"`
		Continent   string `json:"continent,omitempty"`
		Country     string `json:"country,omitempty"`
		Command     string `json:"command,omitempty"`
		Kind        string `json:"kind,omitempty"`
	}

	// GetConfigResponse for a specific application.
	GetConfigResponse struct {
		Meta   map[string]string `json:"meta,omitempty"`
		Error  *Error            `json:"error,omitempty"`
		Config *Config           `json:"config,omitempty"`
	}

	// Config for a specific application.
	Config struct {
		Application string `json:"application,omitempty"`
		Version     string `json:"version,omitempty"`
		Environment string `json:"environment,omitempty"`
		Continent   string `json:"continent,omitempty"`
		Country     string `json:"country,omitempty"`
		Command     string `json:"command,omitempty"`
		Kind        string `json:"kind,omitempty"`
		Data        []byte `json:"data,omitempty"`
	}

	configErrorer struct{}
)

// GetConfig for HTTP.
func (s *Server) GetConfig(ctx context.Context, req *GetConfigRequest) (*GetConfigResponse, error) {
	resp := &GetConfigResponse{}

	cfg, err := service.NewConfig(req.Application, req.Version,
		req.Environment, req.Continent, req.Country,
		req.Command, req.Kind)
	if err != nil {
		return resp, err
	}

	data, err := s.service.GetConfig(ctx, cfg)

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Config = &Config{
		Application: cfg.Application(), Version: cfg.Version(),
		Environment: cfg.Environment(), Continent: cfg.Continent(), Country: cfg.Country(),
		Command: cfg.Command(), Kind: cfg.Kind(), Data: data,
	}

	return resp, err
}

func (*configErrorer) Error(ctx context.Context, err error) *GetConfigResponse {
	return &GetConfigResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (*configErrorer) Status(err error) int {
	if service.IsInvalidArgument(err) {
		return http.StatusBadRequest
	}

	if service.IsNotFoundError(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
