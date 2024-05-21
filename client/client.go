package client

import (
	"context"

	"github.com/alexfalkowski/go-service/time"
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
	ctx, cancel := context.WithTimeout(ctx, time.MustParseDuration(c.config.Timeout))
	defer cancel()

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
	ctx, cancel := context.WithTimeout(ctx, time.MustParseDuration(c.config.Timeout))
	defer cancel()

	req := &v1.GetSecretsRequest{Secrets: c.config.Secrets.Files}

	resp, err := c.client.GetSecrets(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetSecrets(), nil
}
