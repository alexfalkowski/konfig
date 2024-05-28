package git

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/alexfalkowski/go-service/file"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/konfig/git"
	source "github.com/alexfalkowski/konfig/source/configurator"
	ce "github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/google/go-github/v62/github"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for git.
func NewConfigurator(cfg *Config, t trace.Tracer, client *http.Client) *Configurator {
	cl := github.NewClient(client)

	return &Configurator{cfg: cfg, tracer: t, client: cl}
}

// Configurator for git.
type Configurator struct {
	tracer trace.Tracer
	cfg    *Config
	client *github.Client
}

// GetConfig for git.
func (c *Configurator) GetConfig(ctx context.Context, params source.ConfigParams) (*source.Config, error) {
	t, err := c.cfg.GetToken()
	if err != nil {
		return nil, err
	}

	client := c.client.WithAuthToken(t)
	path := c.path(params.Application, params.Environment, params.Continent, params.Country, params.Command, params.Kind)

	ctx, span := c.span(ctx)
	defer span.End()

	tag := fmt.Sprintf("%s/%s", params.Application, params.Version)
	opts := &github.RepositoryContentGetOptions{Ref: tag}

	rc, _, err := client.Repositories.DownloadContents(ctx, c.cfg.Owner, c.cfg.Repository, path, opts)
	if err != nil {
		if git.IsNotFoundError(err) {
			meta.WithAttribute(ctx, "gitError", meta.Error(err))

			return nil, ce.ErrNotFound
		}

		return nil, err
	}

	d, err := io.ReadAll(rc)

	return &source.Config{Kind: file.Extension(path), Data: d}, err
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

func (c *Configurator) span(ctx context.Context) (context.Context, trace.Span) {
	ctx, span := c.tracer.Start(ctx, operationName("get config"), trace.WithSpanKind(trace.SpanKindClient))
	ctx = tracer.WithTraceID(ctx, span)

	return ctx, span
}

func operationName(name string) string {
	return tracer.OperationName("git", name)
}
