package ssm

import (
	"bytes"
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/encoding/json"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"go.opentelemetry.io/otel/trace"
)

// IsNotFound for SSM.
func IsNotFound(err error) bool {
	var perr *types.ParameterNotFound

	return errors.As(err, &perr)
}

var errMissing = errors.New("missing value")

// Secret from SSM.
type Secret struct {
	Data map[string]any `json:"data"`
}

// Transformer for SSM.
type Transformer struct {
	client *ssm.Client
	json   *json.Encoder
	tracer trace.Tracer
}

// NewTransformer for SSM.
func NewTransformer(client *ssm.Client, json *json.Encoder, tracer trace.Tracer) *Transformer {
	return &Transformer{client: client, json: json, tracer: tracer}
}

// Transform for SSM.
func (t *Transformer) Transform(ctx context.Context, value string) (string, error) {
	ctx, span := t.tracer.Start(ctx, operationName("transform"), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tracer.WithTraceID(ctx, span)

	out, err := t.client.GetParameter(ctx, &ssm.GetParameterInput{Name: &value})
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		if IsNotFound(err) {
			return value, errMissing
		}

		return value, err
	}

	var sec Secret

	if err := t.json.Decode(bytes.NewReader([]byte(*out.Parameter.Value)), &sec); err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		return value, err
	}

	tracer.Meta(ctx, span)

	v, ok := sec.Data["value"].(string)
	if !ok {
		return value, errMissing
	}

	return v, nil
}

// IsMissing value for SSM.
func (t *Transformer) IsMissing(err error) bool {
	return errors.Is(err, errMissing)
}

func operationName(name string) string {
	return tracer.OperationName("ssm", name)
}
