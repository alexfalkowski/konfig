package git

import (
	"context"
	"fmt"
	"io"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/konfig/git"
	ce "github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/google/go-github/v63/github"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for git.
func NewConfigurator(client *github.Client, cfg *Config, t trace.Tracer) *Configurator {
	return &Configurator{cfg: cfg, tracer: t, client: client}
}

// Configurator for git.
type Configurator struct {
	tracer trace.Tracer
	cfg    *Config
	client *github.Client
}

// GetConfig for git.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, country, cmd, kind string) ([]byte, error) {
	t, err := c.cfg.GetToken()
	if err != nil {
		return nil, err
	}

	client := c.client.WithAuthToken(t)
	path := c.path(app, env, continent, country, cmd, kind)

	ctx, span := c.span(ctx)
	defer span.End()

	tag := fmt.Sprintf("%s/%s", app, ver)
	opts := &github.RepositoryContentGetOptions{Ref: tag}

	rc, _, err := client.Repositories.DownloadContents(ctx, c.cfg.Owner, c.cfg.Repository, path, opts)
	if err != nil {
		if git.IsNotFound(err) {
			meta.WithAttribute(ctx, "gitError", meta.Error(err))

			return nil, ce.ErrNotFound
		}

		return nil, err
	}

	d, err := io.ReadAll(rc)

	return d, err
}

func (c *Configurator) path(app, env, continent, country, cmd, kind string) string {
	if continent == "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s.%s", app, env, cmd, kind)
	}

	if continent != "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s/%s.%s", app, env, continent, cmd, kind)
	}

	return fmt.Sprintf("%s/%s/%s/%s/%s.%s", app, env, continent, country, cmd, kind)
}

//nolint:spancheck
func (c *Configurator) span(ctx context.Context) (context.Context, trace.Span) {
	ctx, span := c.tracer.Start(ctx, operationName("get config"), trace.WithSpanKind(trace.SpanKindClient))
	ctx = tracer.WithTraceID(ctx, span)

	return ctx, span
}

func operationName(name string) string {
	return tracer.OperationName("git", name)
}
