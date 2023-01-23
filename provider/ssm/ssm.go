package ssm

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/alexfalkowski/konfig/provider/ssm/trace/opentracing"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

// Secret from SSM.
type Secret struct {
	Data map[string]any `json:"data"`
}

// Transformer for SSM.
type Transformer struct {
	client *ssm.Client
	tracer opentracing.Tracer
}

// NewTransformer for SSM.
func NewTransformer(client *ssm.Client, tracer opentracing.Tracer) *Transformer {
	return &Transformer{client: client, tracer: tracer}
}

// Transform for SSM.
func (t *Transformer) Transform(ctx context.Context, value string) (any, error) {
	ctx, span := opentracing.StartSpanFromContext(ctx, t.tracer, "transform", value)
	defer span.Finish()

	out, err := t.client.GetParameter(ctx, &ssm.GetParameterInput{Name: &value})
	if err != nil {
		var perr *types.ParameterNotFound
		if errors.As(err, &perr) {
			return value, nil
		}

		return value, err
	}

	var sec Secret

	if err := json.Unmarshal([]byte(*out.Parameter.Value), &sec); err != nil {
		return value, err
	}

	v := sec.Data["value"]
	if v == nil {
		return value, nil
	}

	return v, nil
}
