package v1

import (
	"github.com/alexfalkowski/konfig/aws"
	"github.com/alexfalkowski/konfig/git"
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc"
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc/security/token"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	aws.Module,
	git.Module,
	fx.Provide(token.NewVerifier),
	fx.Provide(grpc.NewServer),
	fx.Invoke(grpc.Register),
)
