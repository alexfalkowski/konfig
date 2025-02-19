package git

import (
	"context"
	"fmt"
	"io"

	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/konfig/internal/git"
	"github.com/alexfalkowski/konfig/internal/source/errors"
	"github.com/google/go-github/v68/github"
)

// NewConfigurator for git.
func NewConfigurator(client *github.Client, config *Config, tracer *tracer.Tracer) *Configurator {
	return &Configurator{config: config, tracer: tracer, client: client}
}

// Configurator for git.
type Configurator struct {
	tracer *tracer.Tracer
	config *Config
	client *github.Client
}

// GetConfig for git.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, country, cmd, kind string) ([]byte, error) {
	ctx, span := c.tracer.StartClient(ctx, operationName("get config"))
	defer span.End()

	t, err := c.config.GetToken()
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		return nil, err
	}

	client := c.client.WithAuthToken(t)
	path := c.path(app, env, continent, country, cmd, kind)

	tag := fmt.Sprintf("%s/%s", app, ver)
	opts := &github.RepositoryContentGetOptions{Ref: tag}

	rc, _, err := client.Repositories.DownloadContents(ctx, c.config.Owner, c.config.Repository, path, opts)
	if err != nil {
		tracer.Meta(ctx, span)
		tracer.Error(err, span)

		if git.IsNotFound(err) {
			return nil, fmt.Errorf("%w: %w", err, errors.ErrNotFound)
		}

		return nil, err
	}

	tracer.Meta(ctx, span)

	return io.ReadAll(rc)
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

func operationName(name string) string {
	return tracer.OperationName("git", name)
}
