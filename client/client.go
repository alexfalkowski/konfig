package client

import (
	"context"

	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
	v1c "github.com/alexfalkowski/konfig/client/v1/config"
)

// Client for konfig.
type Client struct {
	client v1.ServiceClient
	config *v1c.Config
}

// NewClient for konfig.
func NewClient(client v1.ServiceClient, config *v1c.Config) *Client {
	return &Client{client: client, config: config}
}

// Config from client.
func (c *Client) Config(ctx context.Context) ([]byte, error) {
	req := &v1.GetConfigRequest{
		Application: c.config.Application,
		Version:     c.config.Version,
		Environment: c.config.Environment,
		Continent:   c.config.Continent,
		Country:     c.config.Country,
		Command:     c.config.Command,
		Kind:        c.config.Kind,
	}

	resp, err := c.client.GetConfig(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Config.Data, nil
}
