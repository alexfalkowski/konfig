package grpc

import (
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health/transport/grpc"
)

// NewObserver for gRPC.
func NewObserver(healthServer *server.Server) *grpc.Observer {
	return &grpc.Observer{Observer: healthServer.Observe("noop")}
}
