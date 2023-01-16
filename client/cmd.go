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
	if err != nil {
		return nil, err
	}

	return &OutputConfig{Config: c}, nil
}

// RunCommandParams for client.
type RunCommandParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	Config       *v1.Config
	Client       *Client
	OutputConfig *OutputConfig
}

// RunCommand for client.
func RunCommand(params RunCommandParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			c := params.Config
			d, err := params.Client.Config(ctx, c.Application, c.Version, c.Environment, c.Continent, c.Country, c.Command, c.Kind)
			if err != nil {
				return err
			}

			return params.OutputConfig.Write(d, fs.FileMode(c.Mode))
		},
	})
}
