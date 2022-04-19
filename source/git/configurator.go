package git

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/alexfalkowski/go-service/meta"
	serrors "github.com/alexfalkowski/konfig/source/errors"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

const buffSize = 8192

// Configurator for git.
type Configurator struct {
	cfg  *Config
	repo *git.Repository
	mux  sync.Mutex
}

// GetConfig for git.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, cluster, cmd string) ([]byte, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err := c.clone(ctx); err != nil {
		meta.WithAttribute(ctx, "git.clone_error", err.Error())

		return nil, serrors.ErrNotFound
	}

	if err := c.pull(ctx); err != nil {
		meta.WithAttribute(ctx, "git.pull_error", err.Error())

		return nil, serrors.ErrNotFound
	}

	file, err := c.file(app, ver, env, cluster, cmd)
	if err != nil {
		meta.WithAttribute(ctx, "git.file_error", err.Error())

		return nil, serrors.ErrNotFound
	}

	return c.bytes(file), nil
}

func (c *Configurator) bytes(reader io.Reader) []byte {
	data := make([]byte, 0)
	buf := make([]byte, buffSize)

	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}

		data = append(data, buf[:n]...)
	}

	return data
}

func (c *Configurator) file(app, ver, env, cluster, cmd string) (billy.File, error) {
	tree, _ := c.repo.Worktree()

	err := tree.Checkout(&git.CheckoutOptions{Branch: plumbing.NewTagReferenceName(fmt.Sprintf("%s/%s", app, ver))})
	if err != nil {
		return nil, err
	}

	var path string

	if cluster == "*" {
		path = fmt.Sprintf("%s/%s/%s.config.yml", app, env, cmd)
	} else {
		path = fmt.Sprintf("%s/%s/%s/%s.config.yml", app, env, cluster, cmd)
	}

	file, err := tree.Filesystem.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (c *Configurator) pull(ctx context.Context) error {
	tree, _ := c.repo.Worktree()

	if err := tree.Checkout(&git.CheckoutOptions{Branch: plumbing.Master}); err != nil {
		return err
	}

	if err := tree.PullContext(ctx, &git.PullOptions{RemoteName: "origin"}); err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			return nil
		}

		return err
	}

	return nil
}

func (c *Configurator) clone(ctx context.Context) error {
	if c.repo != nil {
		return nil
	}

	if err := os.RemoveAll(c.cfg.Dir); err != nil {
		return err
	}

	opts := &git.CloneOptions{Auth: &http.BasicAuth{Username: "a", Password: c.cfg.Token}, URL: c.cfg.URL}

	r, err := git.PlainCloneContext(ctx, c.cfg.Dir, false, opts)
	if err != nil {
		return err
	}

	c.repo = r

	return nil
}
