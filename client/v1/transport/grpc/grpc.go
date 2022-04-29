package grpc

import (
	"context"

	"github.com/alexfalkowski/konfig/client/task"
	"go.uber.org/fx"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Task task.Task
}

// Register client.
func Register(lc fx.Lifecycle, params RegisterParams) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return params.Task.Perform(ctx)
		},
	})
}
