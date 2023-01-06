package health

import (
	"time"
)

// Config for health.
type Config struct {
	Duration time.Duration `yaml:"duration" json:"duration"`
	Timeout  time.Duration `yaml:"timeout" json:"timeout"`
}
