package env

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/strings"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
)

var errMissing = errors.New("missing value")

// Transformer for env.
type Transformer struct {
	tracer *tracer.Tracer
}

// NewTransformer for env.
func NewTransformer(tracer *tracer.Tracer) *Transformer {
	return &Transformer{tracer: tracer}
}

// Transform for env.
func (t *Transformer) Transform(ctx context.Context, value string) (string, error) {
	ctx, span := t.tracer.StartClient(ctx, operationName("transform"))
	defer span.End()

	tracer.Meta(ctx, span)

	if v := os.Getenv(value); !strings.IsEmpty(v) {
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
