package git

// Config for git.
type Config struct {
	URL   string `yaml:"url"`
	Dir   string `yaml:"dir"`
	Token string `yaml:"token"`
}
