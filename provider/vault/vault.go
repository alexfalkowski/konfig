package vault

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/meta"
	tm "github.com/alexfalkowski/go-service/transport/meta"
	"github.com/alexfalkowski/konfig/provider/vault/telemetry/tracer"
	"github.com/hashicorp/vault/api"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var errMissing = errors.New("missing value")

// Transformer for vault.
type Transformer struct {
	client *api.Client
	tracer tracer.Tracer
}

// NewTransformer for vault.
func NewTransformer(client *api.Client, t tracer.Tracer) *Transformer {
	return &Transformer{client: client, tracer: t}
}

// Transform for vault.
func (t *Transformer) Transform(ctx context.Context, value string) (any, error) {
	ctx, span := t.tracer.Start(ctx, "transform", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToValuer(span.SpanContext().TraceID()))

	sec, err := t.client.Logical().ReadWithContext(ctx, value)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return value, err
	}

	if sec == nil || sec.Data == nil {
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

	return md["value"], nil
}

// IsMissing value for vault.
func (t *Transformer) IsMissing(err error) bool {
	return errors.Is(err, errMissing)
}
