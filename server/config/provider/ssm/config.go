package ssm

// Config for SSM.
type Config struct {
	Access string `yaml:"access" json:"access" toml:"access"`
	Secret string `yaml:"secret" json:"secret" toml:"secret"`
	Region string `yaml:"region" json:"region" toml:"region"`
	URL    string `yaml:"url" json:"url" toml:"url"`
}
