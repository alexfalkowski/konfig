package v1

import (
	"github.com/alexfalkowski/konfig/aws"
	"github.com/alexfalkowski/konfig/git"
	"github.com/alexfalkowski/konfig/server/config"
	"github.com/alexfalkowski/konfig/server/security/token"
	"github.com/alexfalkowski/konfig/server/v1/transport/grpc"
	"github.com/alexfalkowski/konfig/server/v1/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	config.Module,
	aws.Module,
	git.Module,
	token.Module,
	fx.Provide(grpc.NewServer),
	fx.Invoke(grpc.Register),
	fx.Invoke(http.Register),
)
