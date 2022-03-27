package grpc

import (
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health/transport/grpc"
)

// NewObserver for gRPC.
func NewObserver(healthServer *server.Server) (*grpc.Observer, error) {
	ob, _ := healthServer.Observe("noop")

	return &grpc.Observer{Observer: ob}, nil
}
