package file

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
)

var errMissing = errors.New("missing value")

// Transformer for file.
type Transformer struct {
	fs     os.FileSystem
	tracer *tracer.Tracer
}

// NewTransformer for file.
func NewTransformer(fs os.FileSystem, tracer *tracer.Tracer) *Transformer {
	return &Transformer{fs: fs, tracer: tracer}
}

// Transform for file.
func (t *Transformer) Transform(ctx context.Context, value string) (string, error) {
	ctx, span := t.tracer.StartClient(ctx, operationName("transform"))
	defer span.End()

	tracer.Meta(ctx, span)

	bytes, err := t.fs.ReadFile(value)
	if err != nil {
		tracer.Error(err, span)

		if t.fs.IsNotExist(err) {
			return value, errMissing
		}

		return value, err
	}

	return string(bytes), nil
}

// IsMissing value for file.
func (t *Transformer) IsMissing(err error) bool {
	return errors.Is(err, errMissing)
}

func operationName(name string) string {
	return tracer.OperationName("secret", name)
}
