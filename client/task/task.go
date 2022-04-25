package task

import (
	"context"
)

// Task to be performed.
type Task interface {
	Perform(context.Context) error
}
