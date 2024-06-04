package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexfalkowski/go-service/meta"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	source "github.com/alexfalkowski/konfig/source/configurator"
	ce "github.com/alexfalkowski/konfig/source/configurator/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetConfig for gRPC.
func (s *Server) GetConfig(ctx context.Context, req *v1.GetConfigRequest) (*v1.GetConfigResponse, error) {
	if req.GetContinent() == "" {
		req.Continent = "*"
	}

	if req.GetCountry() == "" {
		req.Country = "*"
	}

	if req.GetKind() == "" {
		req.Kind = "yaml"
	}

	resp := &v1.GetConfigResponse{
		Config: &v1.Config{
			Application: req.GetApplication(), Version: req.GetVersion(),
			Environment: req.GetEnvironment(), Continent: req.GetContinent(), Country: req.GetCountry(),
			Command: req.GetCommand(), Kind: req.GetKind(),
		},
	}

	if err := s.validateGetConfigRequest(req); err != nil {
		resp.Meta = s.meta(ctx)

		return resp, err
	}

	p := source.ConfigParams{
		Application: req.GetApplication(), Version: req.GetVersion(),
		Environment: req.GetEnvironment(), Continent: req.GetContinent(), Country: req.GetCountry(),
		Command: req.GetCommand(), Kind: req.GetKind(),
	}

	c, err := s.conf.GetConfig(ctx, p)
	if err != nil {
		if errors.Is(err, ce.ErrNotFound) {
			resp.Meta = s.meta(ctx)

			return resp, status.Error(codes.NotFound, fmt.Sprintf("%s was not found", p))
		}

		ctx = meta.WithAttribute(ctx, "configError", meta.Error(err))
		resp.Meta = s.meta(ctx)

		return resp, status.Error(codes.Internal, "could not get config")
	}

	data, err := s.transformer.Transform(ctx, c)
	if err != nil {
		ctx = meta.WithAttribute(ctx, "configError", meta.Error(err))
		resp.Meta = s.meta(ctx)

		return resp, status.Error(codes.Internal, "could not transform")
	}

	resp.Config.Kind = c.Kind
	resp.Config.Data = data
	resp.Meta = s.meta(ctx)

	return resp, nil
}

func (s *Server) validateGetConfigRequest(req *v1.GetConfigRequest) error {
	if req.GetApplication() == "" {
		return status.Error(codes.InvalidArgument, "invalid application")
	}

	if req.GetVersion() == "" {
		return status.Error(codes.InvalidArgument, "invalid version")
	}

	if req.GetEnvironment() == "" {
		return status.Error(codes.InvalidArgument, "invalid environment")
	}

	if req.GetCommand() == "" {
		return status.Error(codes.InvalidArgument, "invalid command")
	}

	return nil
}