package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexfalkowski/go-service/cache/redis"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/server/config"
	"github.com/alexfalkowski/konfig/source"
	kerrors "github.com/alexfalkowski/konfig/source/errors"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServerParams for gRPC.
type ServerParams struct {
	fx.In

	RedisConfig  *redis.Config
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
	if req.Cluster == "" {
		req.Cluster = "*"
	}

	resp := &v1.GetConfigResponse{
		Config: &v1.Config{
			Application: req.Application,
			Version:     req.Version,
			Environment: req.Environment,
			Cluster:     req.Cluster,
			Command:     req.Command,
			ContentType: "application/yaml",
		},
	}

	if err := s.validateGetConfigRequest(req); err != nil {
		return resp, err
	}

	data, err := s.conf.GetConfig(ctx, req.Application, req.Version, req.Environment, req.Cluster, req.Command)
	if err != nil && errors.Is(err, kerrors.ErrNotFound) {
		msg := fmt.Sprintf("%s/%s/%s/%s was not found", req.Application, req.Version, req.Environment, req.Command)

		return resp, status.Error(codes.NotFound, msg)
	}

	data, err = s.trans.Transform(ctx, data)
	if err != nil {
		return resp, status.Error(codes.Internal, "could not transform")
	}

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
