package otel

import (
	"github.com/alexfalkowski/go-service/otel"
	"github.com/alexfalkowski/go-service/version"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// Tracer for otel.
type Tracer trace.Tracer

// NewTracer for otel.
func NewTracer(lc fx.Lifecycle, cfg *otel.Config, version version.Version) (Tracer, error) {
	return otel.NewTracer(otel.TracerParams{Lifecycle: lc, Name: "vault", Config: cfg, Version: version})
}
