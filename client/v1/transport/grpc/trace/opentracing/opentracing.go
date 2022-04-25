package opentracing

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/time"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/client/v1/transport/grpc/task"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

// Client for opentracing.
type Client struct {
	cfg *client.Config
	task.Client
}

// NewClient for zap.
func NewClient(cfg *client.Config, task task.Client) *Client {
	return &Client{cfg: cfg, Client: task}
}

// Perform logger for client.
func (c *Client) Perform(ctx context.Context) error {
	start := time.Now().UTC()
	tracer := opentracing.GlobalTracer()
	operationName := fmt.Sprintf("sync %s/%s/%s/%s", c.cfg.Application, c.cfg.Version, c.cfg.Environment, c.cfg.Command)
	opts := []opentracing.StartSpanOption{
		opentracing.Tag{Key: "client.start_time", Value: start.Format(time.RFC3339)},
		opentracing.Tag{Key: "client.application", Value: c.cfg.Application},
		opentracing.Tag{Key: "client.version", Value: c.cfg.Version},
		opentracing.Tag{Key: "client.environment", Value: c.cfg.Environment},
		opentracing.Tag{Key: "client.command", Value: c.cfg.Command},
		opentracing.Tag{Key: "component", Value: "client"},
		ext.SpanKindRPCClient,
	}

	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, operationName, opts...)
	defer span.Finish()

	err := c.Client.Perform(ctx)

	span.SetTag("client.duration", time.ToMilliseconds(time.Since(start)))

	if err != nil {
		ext.Error.Set(span, true)
		span.LogFields(log.String("event", "error"), log.String("message", err.Error()))
	}

	return err
}
