package health

import (
	"time"
)

// Config for health.
type Config struct {
	Duration time.Duration `yaml:"duration"`
	Timeout  time.Duration `yaml:"timeout"`
}
