package s3

// Config for s3.
type Config struct {
	Access string `yaml:"access"`
	Secret string `yaml:"secret"`
	Region string `yaml:"region"`
	Bucket string `yaml:"bucket"`
	URL    string `yaml:"url"`
}
