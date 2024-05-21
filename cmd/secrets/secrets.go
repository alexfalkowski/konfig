package secrets

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/konfig/client"
	v1 "github.com/alexfalkowski/konfig/client/v1/config"
	"go.uber.org/fx"
)

// Params for secrets.
type Params struct {
	fx.In

	Lifecycle    fx.Lifecycle
	Client       *client.Client
	OutputConfig *cmd.OutputConfig
	Config       *v1.Config
}

// Start for secrets.
func Start(params Params) {
	cmd.Start(params.Lifecycle, func(ctx context.Context) {
		secrets, err := params.Client.Secrets(ctx)
		runtime.Must(err)

		cfg := params.Config.Secrets

		for n, v := range secrets {
			p := filepath.Join(cfg.Path, n)

			err := os.WriteFile(p, v, fs.FileMode(cfg.Mode))
			runtime.Must(err)
		}
	})
}
