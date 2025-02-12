package health

// Config for health.
type Config struct {
	Duration string `yaml:"duration,omitempty" json:"duration,omitempty" toml:"duration,omitempty"`
	Timeout  string `yaml:"timeout,omitempty" json:"timeout,omitempty" toml:"timeout,omitempty"`
}
