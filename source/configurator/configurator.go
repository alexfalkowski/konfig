package configurator

import (
	"context"
)

// Configurator for configurator.
type Configurator interface {
	GetConfig(ctx context.Context, app, ver, env, continent, country, cmd string) ([]byte, error)
}
