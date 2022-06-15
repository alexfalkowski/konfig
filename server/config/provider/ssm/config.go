package ssm

// Config for SSM.
type Config struct {
	Access string `yaml:"access"`
	Secret string `yaml:"secret"`
	Region string `yaml:"region"`
	URL    string `yaml:"url"`
}
