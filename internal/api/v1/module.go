package v1

import (
	"github.com/alexfalkowski/konfig/internal/api/config"
	"github.com/alexfalkowski/konfig/internal/api/v1/transport/grpc"
	"github.com/alexfalkowski/konfig/internal/api/v1/transport/http"
	"github.com/alexfalkowski/konfig/internal/aws"
	"github.com/alexfalkowski/konfig/internal/git"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	config.Module,
	aws.Module,
	git.Module,
	fx.Provide(grpc.NewServer),
	fx.Invoke(grpc.Register),
	fx.Invoke(http.Register),
)
