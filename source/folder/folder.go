package folder

// NewConfigurator for folder.
func NewConfigurator(cfg *Config) *Configurator {
	return &Configurator{cfg: cfg}
}
