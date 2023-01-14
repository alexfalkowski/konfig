package grpc

import (
	"context"
	"errors"
	"fmt"

	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/server/config"
	"github.com/alexfalkowski/konfig/server/v1/source"
	"github.com/alexfalkowski/konfig/source/configurator"
	kerrors "github.com/alexfalkowski/konfig/source/configurator/errors"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServerParams for gRPC.
type ServerParams struct {
	fx.In

	Configurator source.Configurator
	Transformer  *config.Transformer
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{conf: params.Configurator, trans: params.Transformer}
}

// Server for gRPC.
type Server struct {
	conf  source.Configurator
	trans *config.Transformer
	v1.UnimplementedServiceServer
}

// GetConfig for gRPC.
func (s *Server) GetConfig(ctx context.Context, req *v1.GetConfigRequest) (*v1.GetConfigResponse, error) {
	if req.Continent == "" {
		req.Continent = "*"
	}

	if req.Country == "" {
		req.Country = "*"
	}

	if req.Kind == "" {
		req.Kind = "yaml"
	}

	resp := &v1.GetConfigResponse{
		Config: &v1.Config{
			Application: req.Application,
			Version:     req.Version,
			Environment: req.Environment,
			Continent:   req.Continent,
			Country:     req.Country,
			Command:     req.Command,
			Kind:        req.Kind,
		},
	}

	if err := s.validateGetConfigRequest(req); err != nil {
		return resp, err
	}

	p := configurator.ConfigParams{
		Application: req.Application,
		Version:     req.Version,
		Environment: req.Environment,
		Continent:   req.Continent,
		Country:     req.Country,
		Command:     req.Command,
		Kind:        req.Kind,
	}

	c, err := s.conf.GetConfig(ctx, p)
	if err != nil {
		if errors.Is(err, kerrors.ErrNotFound) {
			return resp, status.Error(codes.NotFound, fmt.Sprintf("%s was not found", p))
		}

		return resp, status.Error(codes.Internal, "could get config")
	}

	data, err := s.trans.Transform(ctx, c)
	if err != nil {
		return resp, status.Error(codes.Internal, "could not transform")
	}

	resp.Config.Kind = c.Kind
	resp.Config.Data = data

	return resp, nil
}

func (s *Server) validateGetConfigRequest(req *v1.GetConfigRequest) error {
	if req.Application == "" {
		return status.Error(codes.InvalidArgument, "invalid application")
	}

	if req.Version == "" {
		return status.Error(codes.InvalidArgument, "invalid version")
	}

	if req.Environment == "" {
		return status.Error(codes.InvalidArgument, "invalid environment")
	}

	if req.Command == "" {
		return status.Error(codes.InvalidArgument, "invalid command")
	}

	return nil
}
