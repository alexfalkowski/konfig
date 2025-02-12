package health

import (
	"github.com/alexfalkowski/go-service/health"
	"github.com/alexfalkowski/konfig/internal/health/transport/grpc"
	"github.com/alexfalkowski/konfig/internal/health/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(http.NewHealthObserver), fx.Provide(http.NewLivenessObserver), fx.Provide(http.NewReadinessObserver),
	fx.Provide(grpc.NewObserver), fx.Provide(NewRegistrations), health.Module,
)
