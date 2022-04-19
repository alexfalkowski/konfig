package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/config"
	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	"github.com/alexfalkowski/konfig/client"
)

type clt struct {
	client v1.ConfiguratorServiceClient
	cfg    *client.Config
}

// Perform default.
func (c *clt) Perform(ctx context.Context) error {
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

	return config.WriteFileToEnv("APP_CONFIG_FILE", resp.Data)
}
