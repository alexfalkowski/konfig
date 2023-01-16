package client

import (
	"context"

	v1 "github.com/alexfalkowski/konfig/api/konfig/v1"
)

// Client for konfig.
type Client struct {
	client v1.ServiceClient
}

// NewClient for konfig.
func NewClient(client v1.ServiceClient) *Client {
	return &Client{client: client}
}

// Config from client.
func (c *Client) Config(ctx context.Context, app, ver, env, continent, country, cmd, kind string) ([]byte, error) {
	req := &v1.GetConfigRequest{
		Application: app,
		Version:     ver,
		Environment: env,
		Continent:   continent,
		Country:     country,
		Command:     cmd,
		Kind:        kind,
	}

	resp, err := c.client.GetConfig(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Config.Data, nil
}
