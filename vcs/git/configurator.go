package git

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/konfig/vcs/errors"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

const buffSize = 256

// Configurator for git.
type Configurator struct {
	cfg  *Config
	repo *git.Repository
	mux  sync.Mutex
}

// GetConfig for git.
func (c *Configurator) GetConfig(ctx context.Context, app, ver, env, cmd string) ([]byte, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err := c.clone(); err != nil {
		meta.WithAttribute(ctx, "git.clone_error", err.Error())

		return nil, errors.ErrNotFound
	}

	tree, _ := c.repo.Worktree()

	err := tree.Checkout(&git.CheckoutOptions{Branch: plumbing.NewTagReferenceName(fmt.Sprintf("%s/%s", app, ver))})
	if err != nil {
		meta.WithAttribute(ctx, "git.checkout_error", err.Error())

		return nil, errors.ErrNotFound
	}

	file, err := tree.Filesystem.Open(fmt.Sprintf("%s/%s/%s.config.yml", app, env, cmd))
	if err != nil {
		meta.WithAttribute(ctx, "git.open_error", err.Error())

		return nil, errors.ErrNotFound
	}

	data := make([]byte, 0)
	buf := make([]byte, buffSize)

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}

		data = append(data, buf[:n]...)
	}

	return data, nil
}

func (c *Configurator) clone() error {
	if c.repo != nil {
		return nil
	}

	r, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{Auth: &http.BasicAuth{Username: "a", Password: c.cfg.Token}, URL: c.cfg.URL})
	if err != nil {
		return err
	}

	c.repo = r

	return nil
}
