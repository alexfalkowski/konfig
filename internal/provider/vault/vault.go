package vault

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/hashicorp/vault/api"
	"go.opentelemetry.io/otel/trace"
)

var errMissing = errors.New("missing value")

// Transformer for vault.
type Transformer struct {
	client *api.Client
	tracer *tracer.Tracer
}

// NewTransformer for vault.
func NewTransformer(client *api.Client, tracer *tracer.Tracer) *Transformer {
	return &Transformer{client: client, tracer: tracer}
}

// Transform for vault.
func (t *Transformer) Transform(ctx context.Context, value string) (string, error) {
	ctx, span := t.tracer.Start(ctx, operationName("transform"), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tracer.WithTraceID(ctx, span)

	sec, err := t.client.Logical().ReadWithContext(ctx, value)
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		return value, err
	}

	tracer.Meta(ctx, span)

	if sec == nil {
		return value, errMissing
	}

	d := sec.Data["data"]
	if d == nil {
		return value, errMissing
	}

	md, ok := d.(map[string]any)
	if !ok || md["value"] == nil {
		return value, errMissing
	}

	return md["value"].(string), nil //nolint:forcetypeassert
}

// IsMissing value for vault.
func (t *Transformer) IsMissing(err error) bool {
	return errors.Is(err, errMissing)
}

func operationName(name string) string {
	return tracer.OperationName("vault", name)
}
