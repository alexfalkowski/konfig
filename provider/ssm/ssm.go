package ssm

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/meta"
	tm "github.com/alexfalkowski/go-service/transport/meta"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var errMissing = errors.New("missing value")

// Secret from SSM.
type Secret struct {
	Data map[string]any `json:"data"`
}

// Transformer for SSM.
type Transformer struct {
	client *ssm.Client
	json   *marshaller.JSON
	tracer trace.Tracer
}

// NewTransformer for SSM.
func NewTransformer(client *ssm.Client, json *marshaller.JSON, tracer trace.Tracer) *Transformer {
	return &Transformer{client: client, json: json, tracer: tracer}
}

// Transform for SSM.
func (t *Transformer) Transform(ctx context.Context, value string) (any, error) {
	ctx, span := t.tracer.Start(ctx, "transform", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToValuer(span.SpanContext().TraceID()))

	out, err := t.client.GetParameter(ctx, &ssm.GetParameterInput{Name: &value})
	if err != nil {
		var perr *types.ParameterNotFound
		if errors.As(err, &perr) {
			return value, errMissing
		}

		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return value, err
	}

	var sec Secret

	if err := t.json.Unmarshal([]byte(*out.Parameter.Value), &sec); err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return value, err
	}

	v := sec.Data["value"]
	if v == nil {
		return value, errMissing
	}

	return v, nil
}

// IsMissing value for SSM.
func (t *Transformer) IsMissing(err error) bool {
	return errors.Is(err, errMissing)
}
