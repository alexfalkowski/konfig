package http

import (
	"github.com/alexfalkowski/go-service/net/http/rpc"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/internal/api/v1/transport/grpc"
)

// Register for HTTP.
func Register(server *grpc.Server) {
	rpc.Route(v1.Service_GetConfig_FullMethodName, server.GetConfig)
	rpc.Route(v1.Service_GetSecrets_FullMethodName, server.GetSecrets)
}
