package client

// Config for client.
type Config struct {
	Host        string `yaml:"host"`
	Application string `yaml:"application"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
	Command     string `yaml:"command"`
}
