package configurator

import (
	"context"
)

// Configurator for configurator.
type Configurator interface {
	GetConfig(ctx context.Context, app, ver, env, continent, country, cmd, kind string) ([]byte, error)
}
