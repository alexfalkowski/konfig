package grpc

import (
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Conn *grpc.ClientConn
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) v1.ServiceClient {
	return v1.NewServiceClient(params.Conn)
}
