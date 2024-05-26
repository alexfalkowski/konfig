package v1

import (
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc"
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/security/token"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(token.NewVerifier),
	fx.Provide(grpc.NewServer),
	fx.Invoke(grpc.Register),
)
