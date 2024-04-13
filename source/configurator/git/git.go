package git

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexfalkowski/go-service/file"
	"github.com/alexfalkowski/go-service/meta"
	tm "github.com/alexfalkowski/go-service/transport/meta"
	source "github.com/alexfalkowski/konfig/source/configurator"
	cerrors "github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	gclient "github.com/go-git/go-git/v5/plumbing/transport/client"
	ghttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"go.opentelemetry.io/otel/trace"
)

// NewConfigurator for git.
func NewConfigurator(cfg *Config, t trace.Tracer, client *http.Client) *Configurator {
	c := ghttp.NewClient(client)

	gclient.InstallProtocol("http", c)
	gclient.InstallProtocol("https", c)

	return &Configurator{cfg: cfg, tracer: t}
}

// Configurator for git.
type Configurator struct {
	cfg    *Config
	repo   *git.Repository
	mux    sync.Mutex
	tracer trace.Tracer
}

// GetConfig for git.
func (c *Configurator) GetConfig(ctx context.Context, params source.ConfigParams) (*source.Config, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err := c.clone(ctx); err != nil {
		meta.WithAttribute(ctx, "gitCloneError", meta.Error(err))

		return nil, err
	}

	if err := c.pull(ctx); err != nil {
		meta.WithAttribute(ctx, "gitPullError", meta.Error(err))

		return nil, err
	}

	if err := c.checkout(params.Application, params.Version); err != nil {
		meta.WithAttribute(ctx, "gitCheckoutError", meta.Error(err))

		if errors.Is(err, plumbing.ErrReferenceNotFound) {
			return nil, cerrors.ErrNotFound
		}

		return nil, err
	}

	p := c.path(params.Application, params.Environment, params.Continent, params.Country, params.Command, params.Kind)
	path := filepath.Join(c.cfg.Dir, p)

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		meta.WithAttribute(ctx, "gitFileError", meta.Error(err))

		if os.IsNotExist(err) {
			return nil, cerrors.ErrNotFound
		}

		return nil, err
	}

	return &source.Config{Kind: file.Extension(path), Data: data}, nil
}

func (c *Configurator) checkout(app, ver string) error {
	tag := fmt.Sprintf("%s/%s", app, ver)
	tree, _ := c.repo.Worktree()

	return tree.Checkout(&git.CheckoutOptions{Branch: plumbing.NewTagReferenceName(tag)})
}

func (c *Configurator) pull(ctx context.Context) error {
	ctx, span := c.tracer.Start(ctx, "pull", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToValuer(span.SpanContext().TraceID()))

	tree, _ := c.repo.Worktree()

	if err := tree.Checkout(&git.CheckoutOptions{Branch: plumbing.Master}); err != nil {
		return err
	}

	if err := tree.PullContext(ctx, &git.PullOptions{RemoteName: "origin"}); err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return err
	}

	return nil
}

func (c *Configurator) clone(ctx context.Context) error {
	if c.repo != nil {
		return nil
	}

	ctx, span := c.tracer.Start(ctx, "clone", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToValuer(span.SpanContext().TraceID()))

	if err := os.RemoveAll(c.cfg.Dir); err != nil {
		return err
	}

	opts := &git.CloneOptions{Auth: &ghttp.BasicAuth{Username: "a", Password: c.cfg.Token()}, URL: c.cfg.URL}

	r, err := git.PlainCloneContext(ctx, c.cfg.Dir, false, opts)
	if err != nil {
		return err
	}

	c.repo = r

	return nil
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
