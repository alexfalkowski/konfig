package git

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexfalkowski/go-service/meta"
	cerrors "github.com/alexfalkowski/konfig/source/configurator/errors"
	"github.com/alexfalkowski/konfig/source/configurator/git/opentracing"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	gclient "github.com/go-git/go-git/v5/plumbing/transport/client"
	ghttp "github.com/go-git/go-git/v5/plumbing/transport/http"
)

// NewConfigurator for git.
func NewConfigurator(cfg Config, tracer opentracing.Tracer, client *http.Client) *Configurator {
	c := ghttp.NewClient(client)

	gclient.InstallProtocol("http", c)
	gclient.InstallProtocol("https", c)

	return &Configurator{cfg: cfg, tracer: tracer}
}

// Configurator for git.
type Configurator struct {
	cfg    Config
	repo   *git.Repository
	mux    sync.Mutex
	tracer opentracing.Tracer
}

// GetConfig for git.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, continent, country, cmd string) ([]byte, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err := c.clone(ctx); err != nil {
		meta.WithAttribute(ctx, "git.clone_error", err.Error())

		return nil, err
	}

	if err := c.pull(ctx); err != nil {
		meta.WithAttribute(ctx, "git.pull_error", err.Error())

		return nil, err
	}

	if err := c.checkout(ctx, app, ver); err != nil {
		meta.WithAttribute(ctx, "git.checkout_error", err.Error())

		if errors.Is(err, plumbing.ErrReferenceNotFound) {
			return nil, cerrors.ErrNotFound
		}

		return nil, err
	}

	path := filepath.Join(c.cfg.Dir, c.path(app, env, continent, country, cmd))

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		meta.WithAttribute(ctx, "git.file_error", err.Error())

		if os.IsNotExist(err) {
			return nil, cerrors.ErrNotFound
		}

		return nil, err
	}

	return data, nil
}

func (c *Configurator) checkout(ctx context.Context, app, ver string) error {
	tag := fmt.Sprintf("%s/%s", app, ver)

	_, span := opentracing.StartSpanFromContext(ctx, c.tracer, "checkout", tag)
	defer span.Finish()

	tree, _ := c.repo.Worktree()

	return tree.Checkout(&git.CheckoutOptions{Branch: plumbing.NewTagReferenceName(tag)})
}

func (c *Configurator) pull(ctx context.Context) error {
	ctx, span := opentracing.StartSpanFromContext(ctx, c.tracer, "pull", plumbing.Master.String())
	defer span.Finish()

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
	ctx, span := opentracing.StartSpanFromContext(ctx, c.tracer, "clone", c.cfg.URL)
	defer span.Finish()

	if c.repo != nil {
		return nil
	}

	if err := os.RemoveAll(c.cfg.Dir); err != nil {
		return err
	}

	opts := &git.CloneOptions{Auth: &ghttp.BasicAuth{Username: "a", Password: c.cfg.GetToken()}, URL: c.cfg.URL}

	r, err := git.PlainCloneContext(ctx, c.cfg.Dir, false, opts)
	if err != nil {
		return err
	}

	c.repo = r

	return nil
}

func (c *Configurator) path(app, env, continent, country, cmd string) string {
	if continent == "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s.config.yml", app, env, cmd)
	}

	if continent != "*" && country == "*" {
		return fmt.Sprintf("%s/%s/%s/%s.config.yml", app, env, continent, cmd)
	}

	return fmt.Sprintf("%s/%s/%s/%s/%s.config.yml", app, env, continent, country, cmd)
}
