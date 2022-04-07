package vault

import (
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/hashicorp/vault/api"
	"go.uber.org/zap"
)

// NewConfig for vault.
func NewConfig(cfg *http.Config, logger *zap.Logger) *api.Config {
	client := http.NewClient(cfg, logger)
	config := api.DefaultConfig()

	config.HttpClient = client

	return config
}

// NewClient for vault.
func NewClient(cfg *api.Config) (*api.Client, error) {
	return api.NewClient(cfg)
}
