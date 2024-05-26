package v1

import (
	"github.com/alexfalkowski/konfig/client/v1/transport/grpc"
	"github.com/alexfalkowski/konfig/client/v1/transport/grpc/security/token"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(token.NewGenerator),
	fx.Provide(grpc.NewServiceClient),
)
