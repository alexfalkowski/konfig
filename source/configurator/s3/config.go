package s3

// Config for s3.
type Config struct {
	Access string `yaml:"access" json:"access" toml:"access"`
	Secret string `yaml:"secret" json:"secret" toml:"secret"`
	Region string `yaml:"region" json:"region" toml:"region"`
	Bucket string `yaml:"bucket" json:"bucket" toml:"bucket"`
	URL    string `yaml:"url" json:"url" toml:"url"`
}
