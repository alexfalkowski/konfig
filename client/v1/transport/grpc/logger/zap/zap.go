package zap

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/time"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/client/v1/transport/grpc/task"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Client for zap.
type Client struct {
	logger *zap.Logger
	cfg    *client.Config
	task.Client
}

// NewClient for zap.
func NewClient(logger *zap.Logger, cfg *client.Config, task task.Client) *Client {
	return &Client{logger: logger, cfg: cfg, Client: task}
}

// Perform logger for client.
func (c *Client) Perform(ctx context.Context) error {
	start := time.Now().UTC()
	err := c.Client.Perform(ctx)
	fields := []zapcore.Field{
		zap.Int64("client.duration", time.ToMilliseconds(time.Since(start))),
		zap.String("client.start_time", start.Format(time.RFC3339)),
		zap.String("client.application", c.cfg.Application),
		zap.String("client.version", c.cfg.Version),
		zap.String("client.environment", c.cfg.Environment),
		zap.String("client.command", c.cfg.Command),
		zap.String("span.kind", "client"),
		zap.String("component", "client"),
	}

	for k, v := range meta.Attributes(ctx) {
		fields = append(fields, zap.String(k, v))
	}

	if d, ok := ctx.Deadline(); ok {
		fields = append(fields, zap.String("client.deadline", d.UTC().Format(time.RFC3339)))
	}

	if err != nil {
		fields = append(fields, zap.Error(err))
		c.logger.Error("finished call with error", fields...)

		return err
	}

	c.logger.Info("finished call with success", fields...)

	return nil
}
