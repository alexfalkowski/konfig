package s3

// Config for s3.
type Config struct {
	Bucket string `yaml:"bucket,omitempty" json:"bucket,omitempty" toml:"bucket,omitempty"`
}
