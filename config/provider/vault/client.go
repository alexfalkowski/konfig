package vault

import "github.com/hashicorp/vault/api"

// NewConfig for vault.
func NewConfig() *api.Config {
	return api.DefaultConfig()
}

// NewClient for vault.
func NewClient(cfg *api.Config) (*api.Client, error) {
	return api.NewClient(cfg)
}
