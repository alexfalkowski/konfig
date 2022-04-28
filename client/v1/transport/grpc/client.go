package grpc

import (
	"bytes"
	"context"

	"github.com/alexfalkowski/go-service/config"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/client"
)

const configFile = "APP_CONFIG_FILE"

// Client for v1.
type Client struct {
	client v1.ServiceClient
	cfg    *client.Config
}

// Perform getting config.
func (c *Client) Perform(ctx context.Context) error {
	req := &v1.GetConfigRequest{
		Application: c.cfg.Application,
		Version:     c.cfg.Version,
		Environment: c.cfg.Environment,
		Cluster:     c.cfg.Cluster,
		Command:     c.cfg.Command,
	}

	resp, err := c.client.GetConfig(ctx, req)
	if err != nil {
		return err
	}

	data, err := config.ReadFileFromEnv(configFile)
	if err != nil || bytes.Compare(data, resp.Config.Data) != 0 {
		return config.WriteFileToEnv(configFile, resp.Config.Data)
	}

	return nil
}
