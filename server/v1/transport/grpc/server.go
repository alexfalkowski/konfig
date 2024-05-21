package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/provider"
	source "github.com/alexfalkowski/konfig/source/configurator"
	"go.uber.org/fx"
)

// ServerParams for gRPC.
type ServerParams struct {
	fx.In

	Configurator source.Configurator
	Transformer  *source.Transformer
	Provider     *provider.Transformer
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{
		conf: params.Configurator, transformer: params.Transformer, provider: params.Provider,
	}
}

// Server for gRPC.
type Server struct {
	v1.UnimplementedServiceServer
	conf        source.Configurator
	transformer *source.Transformer
	provider    *provider.Transformer
}

func (s *Server) meta(ctx context.Context) map[string]string {
	return meta.CamelStrings(ctx, "")
}
