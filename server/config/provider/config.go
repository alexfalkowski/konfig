package provider

import (
	"github.com/alexfalkowski/konfig/server/config/provider/ssm"
)

// Config for provider.
type Config struct {
	SSM ssm.Config `yaml:"ssm" json:"ssm" toml:"ssm"`
}
