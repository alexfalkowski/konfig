package ssm

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/alexfalkowski/konfig/provider/ssm/telemetry/tracer"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// Secret from SSM.
type Secret struct {
	Data map[string]any `json:"data"`
}

// Transformer for SSM.
type Transformer struct {
	client *ssm.Client
	tracer tracer.Tracer
}

// NewTransformer for SSM.
func NewTransformer(client *ssm.Client, t tracer.Tracer) *Transformer {
	return &Transformer{client: client, tracer: t}
}

// Transform for SSM.
func (t *Transformer) Transform(ctx context.Context, value string) (any, error) {
	ctx, span := t.tracer.Start(ctx, "transform", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	out, err := t.client.GetParameter(ctx, &ssm.GetParameterInput{Name: &value})
	if err != nil {
		var perr *types.ParameterNotFound
		if errors.As(err, &perr) {
			return value, nil
		}

		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return value, err
	}

	var sec Secret

	if err := json.Unmarshal([]byte(*out.Parameter.Value), &sec); err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return value, err
	}

	v := sec.Data["value"]
	if v == nil {
		return value, nil
	}

	return v, nil
}
