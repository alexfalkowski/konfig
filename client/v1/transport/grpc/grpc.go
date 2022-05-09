package grpc

import (
	"context"

	"github.com/alexfalkowski/konfig/client/task"
	"go.uber.org/fx"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Task      task.Task
}

// Register client.
func Register(params RegisterParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return params.Task.Perform(ctx)
		},
	})
}
