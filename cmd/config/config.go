package config

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/konfig/client"
	v1 "github.com/alexfalkowski/konfig/client/v1/config"
	"go.uber.org/fx"
)

// Params for config.
type Params struct {
	fx.In

	Lifecycle    fx.Lifecycle
	Client       *client.Client
	OutputConfig *cmd.OutputConfig
	Config       *v1.Config
}

// Start for config.
func Start(params Params) {
	cmd.Start(params.Lifecycle, func(ctx context.Context) {
		d, err := params.Client.Config(ctx)
		runtime.Must(err)

		err = params.OutputConfig.Write(d, fs.FileMode(params.Config.Configuration.Mode))
		runtime.Must(err)
	})
}
