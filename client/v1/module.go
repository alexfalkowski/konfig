package v1

import (
	"github.com/alexfalkowski/konfig/client/v1/transport/grpc"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(fx.Invoke(grpc.Register))
)
