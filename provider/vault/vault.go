package vault

import (
	"context"

	"github.com/alexfalkowski/konfig/provider/vault/telemetry/tracer"
	"github.com/hashicorp/vault/api"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

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

	sec, err := t.client.Logical().ReadWithContext(ctx, value)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return value, err
	}

	if sec == nil || sec.Data == nil {
		return value, nil
	}

	d := sec.Data["data"]
	if d == nil {
		return value, nil
	}

	md, ok := d.(map[string]any)
	if !ok || md["value"] == nil {
		return value, nil
	}

	return md["value"], nil
}
