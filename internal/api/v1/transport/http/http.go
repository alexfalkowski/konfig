package http

import (
	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/konfig/internal/api/v1/transport/grpc"
)

// Register for HTTP.
func Register(server *grpc.Server) {
	rpc.Route("/v1/config", server.GetConfig)
	rpc.Route("/v1/secrets", server.GetSecrets)
}
