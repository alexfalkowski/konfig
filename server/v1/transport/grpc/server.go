package grpc

import (
	"context"
	"os"
)

// Server for gRPC.
type Server struct {
	UnimplementedConfiguratorServer
}

// NewServer for gRPC.
func NewServer() ConfiguratorServer {
	return &Server{}
}

// GetConfig for gRPC.
func (s *Server) GetConfig(ctx context.Context, req *GetConfigRequest) (*GetConfigResponse, error) {
	bs, _ := os.ReadFile(os.Getenv("CONFIG_FILE"))

	resp := &GetConfigResponse{
		Application: req.Application,
		Environment: req.Environment,
		ContentType: "application/yaml",
		Data:        bs,
	}

	return resp, nil
}
