package grpc

import (
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/internal/api/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register for gRPC.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	v1.RegisterServiceServer(registrar, server)
}

// NewServer for gRPC.
func NewServer(service *config.Configuration) *Server {
	return &Server{service: service}
}

// Server for gRPC.
type Server struct {
	v1.UnimplementedServiceServer
	service *config.Configuration
}

func (s *Server) error(err error) error {
	if err == nil {
		return nil
	}

	if config.IsInvalidArgument(err) {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	if config.IsNotFound(err) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
