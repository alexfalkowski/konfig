package env

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"go.opentelemetry.io/otel/trace"
)

var errMissing = errors.New("missing value")

// Transformer for env.
type Transformer struct {
	tracer trace.Tracer
}

// NewTransformer for env.
func NewTransformer(tracer trace.Tracer) *Transformer {
	return &Transformer{tracer: tracer}
}

// Transform for env.
func (t *Transformer) Transform(ctx context.Context, value string) (string, error) {
	ctx, span := t.tracer.Start(ctx, operationName("transform"), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	tracer.Meta(ctx, span)

	v := os.GetVariable(value)
	if v != "" {
		return v, nil
	}

	return value, errMissing
}

// IsMissing value for env.
func (t *Transformer) IsMissing(err error) bool {
	return errors.Is(err, errMissing)
}

func operationName(name string) string {
	return tracer.OperationName("env", name)
}
