package source

import (
	"context"
)

// Configurator for source.
type Configurator interface {
	GetConfig(ctx context.Context, app, ver, env, continent, country, cmd, kind string) ([]byte, error)
}
