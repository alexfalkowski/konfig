package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/api/config"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
)

// GetConfig for gRPC.
func (s *Server) GetConfig(ctx context.Context, req *v1.GetConfigRequest) (*v1.GetConfigResponse, error) {
	resp := &v1.GetConfigResponse{}

	cfg, err := config.NewConfig(
		req.GetApplication(),
		req.GetVersion(),
		req.GetEnvironment(),
		req.GetContinent(),
		req.GetCountry(),
		req.GetCommand(),
		req.GetKind(),
	)
	if err != nil {
		return resp, s.error(err)
	}

	resp.Config = &v1.Config{
		Application: cfg.Application(), Version: cfg.Version(),
		Environment: cfg.Environment(), Continent: cfg.Continent(),
		Country: cfg.Country(), Command: cfg.Command(), Kind: cfg.Kind(),
	}

	data, err := s.service.GetConfig(ctx, cfg)

	resp.Config.Data = data
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, s.error(err)
}
