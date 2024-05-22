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
	cfg := c.config.Configuration
	req := &v1.GetConfigRequest{
		Application: cfg.Application,
		Version:     cfg.Version,
		Environment: cfg.Environment,
		Continent:   cfg.Continent,
		Country:     cfg.Country,
		Command:     cfg.Command,
		Kind:        cfg.Kind,
	}

	resp, err := c.client.GetConfig(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetConfig().GetData(), nil
}

// Secrets from client.
func (c *Client) Secrets(ctx context.Context) (map[string][]byte, error) {
	req := &v1.GetSecretsRequest{Secrets: c.config.Secrets.Files}

	resp, err := c.client.GetSecrets(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetSecrets(), nil
}
