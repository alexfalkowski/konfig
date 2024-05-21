package client

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/runtime"
	v1 "github.com/alexfalkowski/konfig/client/v1/config"
	"go.uber.org/fx"
)

// Params for client.
type Params struct {
	fx.In

	Lifecycle    fx.Lifecycle
	Client       *Client
	OutputConfig *cmd.OutputConfig
	Config       *v1.Config
}

// GetConfig for client.
func GetConfig(params Params) {
	cmd.Start(params.Lifecycle, func(ctx context.Context) {
		d, err := params.Client.Config(ctx)
		runtime.Must(err)

		err = params.OutputConfig.Write(d, fs.FileMode(params.Config.Configuration.Mode))
		runtime.Must(err)
	})
}

// WriteSecrets for client.
func WriteSecrets(params Params) {
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
