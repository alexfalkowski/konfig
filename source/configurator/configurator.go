package configurator

import (
	"context"
	"fmt"
)

// ConfigParams for configurator.
type ConfigParams struct {
	Application string
	Version     string
	Environment string
	Continent   string
	Country     string
	Command     string
	Kind        string
}

// String of params.
func (p ConfigParams) String() string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", p.Application, p.Version, p.Environment, p.Continent, p.Country, p.Command, p.Kind)
}

// Config for configurator.
type Config struct {
	Kind string
	Data []byte
}

// Configurator for configurator.
type Configurator interface {
	GetConfig(ctx context.Context, params ConfigParams) (*Config, error)
}
