package opentracing

import (
	"context"

	"github.com/alexfalkowski/go-service/trace/opentracing"
	"github.com/alexfalkowski/konfig/source/configurator"
	otr "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/fx"
)

// Tracer for opentracing.
type Tracer otr.Tracer

// StartSpanFromContext for opentracing.
func StartSpanFromContext(ctx context.Context, tracer Tracer, operation, method string, opts ...otr.StartSpanOption) (context.Context, otr.Span) {
	return opentracing.StartSpanFromContext(ctx, tracer, "source", operation, method, opts...)
}

// NewTracer for opentracing.
func NewTracer(lc fx.Lifecycle, cfg *opentracing.Config) (Tracer, error) {
	return opentracing.NewTracer(opentracing.TracerParams{Lifecycle: lc, Name: "source", Config: cfg})
}

// Configurator for opentracing.
type Configurator struct {
	configurator configurator.Configurator
	tracer       Tracer
}

// NewProvider for opentracing.
func NewConfigurator(configurator configurator.Configurator, tracer Tracer) *Configurator {
	return &Configurator{configurator: configurator, tracer: tracer}
}

// GetConfig for opentracing.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, cluster, cmd string) ([]byte, error) {
	ctx, span := StartSpanFromContext(ctx, c.tracer, c.configurator.String(), "get-config")
	defer span.Finish()

	span.SetTag("configurator.app", app)
	span.SetTag("configurator.ver", ver)
	span.SetTag("configurator.env", env)
	span.SetTag("configurator.cluster", cluster)
	span.SetTag("configurator.cmd", cmd)

	bytes, err := c.configurator.GetConfig(ctx, app, ver, env, cluster, cmd)
	if err != nil {
		ext.Error.Set(span, true)
		span.LogFields(log.String("event", "error"), log.String("message", err.Error()))

		return nil, err
	}

	return bytes, nil
}

// String for opentracing.
func (c *Configurator) String() string {
	return "opentracing"
}
