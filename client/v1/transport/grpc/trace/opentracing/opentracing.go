package opentracing

import (
	"context"
	"fmt"
	"time"

	stime "github.com/alexfalkowski/go-service/time"
	gopentracing "github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	"github.com/alexfalkowski/konfig/client/task"
	"github.com/alexfalkowski/konfig/client/v1/config"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

// Client for opentracing.
type Client struct {
	cfg    *config.Config
	tracer gopentracing.Tracer
	task.Task
}

// NewClient for zap.
func NewClient(cfg *config.Config, tracer gopentracing.Tracer, task task.Task) *Client {
	return &Client{cfg: cfg, tracer: tracer, Task: task}
}

// Perform logger for client.
func (c *Client) Perform(ctx context.Context) error {
	start := time.Now().UTC()
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

	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, c.tracer, operationName, opts...)
	defer span.Finish()

	err := c.Task.Perform(ctx)

	span.SetTag("client.duration", stime.ToMilliseconds(time.Since(start)))

	if err != nil {
		ext.Error.Set(span, true)
		span.LogFields(log.String("event", "error"), log.String("message", err.Error()))
	}

	return err
}
