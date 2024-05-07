package git

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/alexfalkowski/go-service/file"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	tm "github.com/alexfalkowski/go-service/transport/meta"
	source "github.com/alexfalkowski/konfig/source/configurator"
	ce "github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	gc "github.com/go-git/go-git/v5/plumbing/transport/client"
	gh "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage"
	"github.com/go-git/go-git/v5/storage/memory"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// NewConfigurator for git.
func NewConfigurator(cfg *Config, t trace.Tracer, client *http.Client) *Configurator {
	c := gh.NewClient(client)

	gc.InstallProtocol("http", c)
	gc.InstallProtocol("https", c)

	gr := &errgroup.Group{}
	gr.SetLimit(1)

	cf := &Configurator{cfg: cfg, tracer: t, storage: memory.NewStorage(), fs: memfs.New(), gr: gr}
	gr.Go(func() error { return cf.clone(context.Background()) })

	return cf
}

// Configurator for git.
type Configurator struct {
	cfg     *Config
	repo    *git.Repository
	mux     sync.Mutex
	tracer  trace.Tracer
	storage storage.Storer
	fs      billy.Filesystem
	gr      *errgroup.Group
}

// GetConfig for git.
func (c *Configurator) GetConfig(ctx context.Context, params source.ConfigParams) (*source.Config, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err := c.wait(ctx); err != nil {
		return nil, err
	}

	if err := c.pull(ctx); err != nil {
		return nil, err
	}

	if err := c.checkout(ctx, params.Application, params.Version); err != nil {
		return nil, err
	}

	path := c.path(params.Application, params.Environment, params.Continent, params.Country, params.Command, params.Kind)
	data, err := c.open(ctx, path)

	return &source.Config{Kind: file.Extension(path), Data: data}, err
}

func (c *Configurator) wait(ctx context.Context) error {
	ctx, span := c.tracer.Start(ctx, operationName("wait"), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToString(span.SpanContext().TraceID()))
	tracer.Meta(ctx, span)

	return c.gr.Wait()
}

func (c *Configurator) checkout(ctx context.Context, app, ver string) error {
	ctx, span := c.tracer.Start(ctx, operationName("checkout"), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToString(span.SpanContext().TraceID()))
	tracer.Meta(ctx, span)

	tag := fmt.Sprintf("%s/%s", app, ver)
	tree, _ := c.repo.Worktree()

	if err := tree.Checkout(&git.CheckoutOptions{Branch: plumbing.NewTagReferenceName(tag)}); err != nil {
		if errors.Is(err, plumbing.ErrReferenceNotFound) {
			meta.WithAttribute(ctx, "gitCheckoutError", meta.Error(err))

			return ce.ErrNotFound
		}

		return err
	}

	return nil
}

func (c *Configurator) pull(ctx context.Context) error {
	ctx, span := c.tracer.Start(ctx, operationName("pull"), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToString(span.SpanContext().TraceID()))
	tracer.Meta(ctx, span)

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
	ctx, span := c.tracer.Start(ctx, operationName("clone"), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToString(span.SpanContext().TraceID()))
	tracer.Meta(ctx, span)

	opts := &git.CloneOptions{Auth: &gh.BasicAuth{Username: "a", Password: c.cfg.Token()}, URL: c.cfg.URL}

	r, err := git.CloneContext(ctx, c.storage, c.fs, opts)
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

func (c *Configurator) open(ctx context.Context, path string) ([]byte, error) {
	ctx, span := c.tracer.Start(ctx, operationName("open"), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToString(span.SpanContext().TraceID()))
	tracer.Meta(ctx, span)

	f, err := c.fs.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			meta.WithAttribute(ctx, "gitFileError", meta.Error(err))

			return nil, ce.ErrNotFound
		}

		return nil, err
	}

	return io.ReadAll(f)
}

func operationName(name string) string {
	return tracer.OperationName("git", name)
}
