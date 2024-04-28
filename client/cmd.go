package client

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/marshaller"
	v1 "github.com/alexfalkowski/konfig/client/v1/config"
	"go.uber.org/fx"
)

// OutputFlag for client.
var OutputFlag string

// OutputConfig for client.
type OutputConfig struct {
	*cmd.Config
}

// NewOutputConfig for client.
func NewOutputConfig(factory *marshaller.Factory) (*OutputConfig, error) {
	c, err := cmd.NewConfig(OutputFlag, factory)

	return &OutputConfig{Config: c}, err
}

// RunCommandParams for client.
type RunCommandParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	Client       *Client
	OutputConfig *OutputConfig
	Config       *v1.Config
}

// RunCommand for client.
func RunCommand(params RunCommandParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			d, err := params.Client.Config(ctx)
			if err != nil {
				return err
			}

			return params.OutputConfig.Write(d, fs.FileMode(params.Config.Mode))
		},
	})
}
