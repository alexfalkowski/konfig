package opentracing

import (
	"context"

	"github.com/alexfalkowski/go-service/trace/opentracing"
	"github.com/alexfalkowski/go-service/version"
	otr "github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
)

// Tracer for opentracing.
type Tracer otr.Tracer

// StartSpanFromContext for opentracing.
func StartSpanFromContext(ctx context.Context, tracer Tracer, operation, method string, opts ...otr.StartSpanOption) (context.Context, otr.Span) {
	return opentracing.StartSpanFromContext(ctx, tracer, "git", operation, method, opts...)
}

// NewTracer for opentracing.
func NewTracer(lc fx.Lifecycle, cfg *opentracing.Config, version version.Version) (Tracer, error) {
	return opentracing.NewTracer(opentracing.TracerParams{Lifecycle: lc, Name: "git", Config: cfg, Version: version})
}
