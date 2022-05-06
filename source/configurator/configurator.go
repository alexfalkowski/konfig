package configurator

import (
	"context"
	"fmt"
)

// Configurator for source.
type Configurator interface {
	fmt.Stringer

	GetConfig(ctx context.Context, app, ver, env, cluster, cmd string) ([]byte, error)
}
