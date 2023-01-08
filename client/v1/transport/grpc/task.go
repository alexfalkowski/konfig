package grpc

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/client/cmd"
	"github.com/alexfalkowski/konfig/client/task"
	kzap "github.com/alexfalkowski/konfig/client/v1/transport/grpc/logger/zap"
	gopentracing "github.com/alexfalkowski/konfig/client/v1/transport/grpc/trace/opentracing"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// TaskParams for gRPC.
type TaskParams struct {
	fx.In

	Client       v1.ServiceClient
	Config       *client.Config
	Tracer       opentracing.Tracer
	Logger       *zap.Logger
	OutputConfig *cmd.OutputConfig
}

// NewTask for gRPC.
func NewTask(params TaskParams) task.Task {
	var clt task.Task = &Task{client: params.Client, cfg: params.Config, out: params.OutputConfig}
	clt = kzap.NewClient(params.Logger, params.Config, clt)
	clt = gopentracing.NewClient(params.Config, params.Tracer, clt)

	return clt
}

// Task for gRPC.
type Task struct {
	client v1.ServiceClient
	cfg    *client.Config
	out    *cmd.OutputConfig
}

// Perform getting config.
func (t *Task) Perform(ctx context.Context) error {
	req := &v1.GetConfigRequest{
		Application: t.cfg.Application,
		Version:     t.cfg.Version,
		Environment: t.cfg.Environment,
		Continent:   t.cfg.Continent,
		Country:     t.cfg.Country,
		Command:     t.cfg.Command,
		Kind:        t.cfg.Kind,
	}

	resp, err := t.client.GetConfig(ctx, req)
	if err != nil {
		return err
	}

	return t.out.Write(resp.Config.Data, fs.FileMode(t.cfg.Mode))
}
