package opentracing

import (
	"context"

	"github.com/alexfalkowski/go-service/cache/redis"
	copentracing "github.com/alexfalkowski/go-service/cache/trace/opentracing"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/time"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

// Server for redis.
type Server struct {
	cfg *redis.Config
	v1.ConfiguratorServiceServer
}

// GetConfig for redis.
func (s *Server) GetConfig(ctx context.Context, req *v1.GetConfigRequest) (*v1.GetConfigResponse, error) {
	start := time.Now().UTC()
	opts := []opentracing.StartSpanOption{
		opentracing.Tag{Key: "server.start_time", Value: start.Format(time.RFC3339)},
		opentracing.Tag{Key: "server.application", Value: req.Application},
		opentracing.Tag{Key: "server.version", Value: req.Application},
		opentracing.Tag{Key: "server.environment", Value: req.Application},
		opentracing.Tag{Key: "server.command", Value: req.Application},
	}

	ctx, span := copentracing.StartSpanFromContext(ctx, "get", s.cfg.Host, opts...)
	defer span.Finish()

	resp, err := s.ConfiguratorServiceServer.GetConfig(ctx, req)

	for k, v := range meta.Attributes(ctx) {
		span.SetTag(k, v)
	}

	span.SetTag("server.duration_ms", time.ToMilliseconds(time.Since(start)))

	if err != nil {
		ext.Error.Set(span, true)
		span.LogFields(log.String("event", "error"), log.String("message", err.Error()))
	}

	return resp, err
}
