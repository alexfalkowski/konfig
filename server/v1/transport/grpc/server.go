package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexfalkowski/konfig/vcs"
	verrors "github.com/alexfalkowski/konfig/vcs/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server for gRPC.
type Server struct {
	conf vcs.Configurator
	UnimplementedConfiguratorServer
}

// GetConfig for gRPC.
func (s *Server) GetConfig(ctx context.Context, req *GetConfigRequest) (*GetConfigResponse, error) {
	resp := &GetConfigResponse{
		Application: req.Application,
		Version:     req.Version,
		Environment: req.Environment,
		Command:     req.Command,
		ContentType: "application/yaml",
	}

	if err := s.validateGetConfigRequest(req); err != nil {
		return resp, err
	}

	data, err := s.conf.GetConfig(ctx, req.Application, req.Version, req.Environment, req.Command)
	if err != nil {
		if errors.Is(err, verrors.ErrNotFound) {
			msg := fmt.Sprintf("%s/%s/%s/%s was not found", req.Application, req.Version, req.Environment, req.Command)

			return resp, status.Error(codes.NotFound, msg)
		}

		return resp, err
	}

	resp.Data = data

	return resp, nil
}

func (s *Server) validateGetConfigRequest(req *GetConfigRequest) error {
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
