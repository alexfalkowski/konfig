package configurator

import (
	"context"
)

// Configurator for source.
type Configurator interface {
	GetConfig(ctx context.Context, app, ver, env, cluster, cmd string) ([]byte, error)
}
