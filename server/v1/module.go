package v1

import (
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc"
	"go.uber.org/fx"
)

var (
	// ServerModule for fx.
	ServerModule = fx.Options(fx.Invoke(grpc.Register))
)
