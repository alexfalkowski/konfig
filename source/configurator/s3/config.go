package s3

// Config for s3.
type Config struct {
	Access string `yaml:"access" json:"access"`
	Secret string `yaml:"secret" json:"secret"`
	Region string `yaml:"region" json:"region"`
	Bucket string `yaml:"bucket" json:"bucket"`
	URL    string `yaml:"url" json:"url"`
}
