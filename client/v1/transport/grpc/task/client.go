package task

import (
	"context"
)

// Client for client.
type Client interface {
	Perform(context.Context) error
}
