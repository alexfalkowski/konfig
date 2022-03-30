package v1

import (
	"github.com/alexfalkowski/konfig/client/v1/transport/grpc"
	"go.uber.org/fx"
)

var (
	// ClientModule for fx.
	ClientModule = fx.Options(fx.Invoke(grpc.Register))
)
