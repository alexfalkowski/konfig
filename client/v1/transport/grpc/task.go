package grpc

import (
	"bytes"
	"context"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/client/task"
	kzap "github.com/alexfalkowski/konfig/client/v1/transport/grpc/logger/zap"
	gopentracing "github.com/alexfalkowski/konfig/client/v1/transport/grpc/trace/opentracing"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const configFile = "APP_CONFIG_FILE"

// TaskParams for gRPC.
type TaskParams struct {
	fx.In

	Client v1.ServiceClient
	Config *client.Config
	Tracer opentracing.Tracer
	Logger *zap.Logger
}

// NewTask for gRPC.
func NewTask(params TaskParams) task.Task {
	var clt task.Task = &Task{client: params.Client, cfg: params.Config}
	clt = kzap.NewClient(params.Logger, params.Config, clt)
	clt = gopentracing.NewClient(params.Config, params.Tracer, clt)

	return clt
}

// Task for gRPC.
type Task struct {
	client v1.ServiceClient
	cfg    *client.Config
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
	}

	resp, err := t.client.GetConfig(ctx, req)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(filepath.Clean(os.Getenv(configFile)))
	if err != nil || !bytes.Equal(data, resp.Config.Data) {
		return os.WriteFile(filepath.Clean(os.Getenv(configFile)), resp.Config.Data, 0o600)
	}

	return nil
}
