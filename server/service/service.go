package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexfalkowski/konfig/provider"
	source "github.com/alexfalkowski/konfig/source/configurator"
	ce "github.com/alexfalkowski/konfig/source/configurator/errors"
)

var (
	// ErrNotFound for service.
	ErrNotFound = errors.New("not found")

	// ErrInvalidArgument for service.
	ErrInvalidArgument = errors.New("invalid argument")
)

// IsNotFoundError for service.
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsInvalidArgument for service.
func IsInvalidArgument(err error) bool {
	return errors.Is(err, ErrInvalidArgument)
}

type (
	// Config for a specific application.
	Config struct {
		application string
		version     string
		environment string
		continent   string
		country     string
		command     string
		kind        string
	}

	// Service for the different transports.
	Service struct {
		provider *provider.Transformer
		config   source.Configurator
		source   *source.Transformer
	}
)

// NewService for the different transports.
func NewService(provider *provider.Transformer, config source.Configurator, source *source.Transformer) *Service {
	return &Service{provider: provider, config: config, source: source}
}

// GetConfig for service.
func (s *Service) GetConfig(ctx context.Context, cfg *Config) ([]byte, error) {
	d, err := s.config.GetConfig(ctx, cfg.application, cfg.version,
		cfg.environment, cfg.continent, cfg.country,
		cfg.command, cfg.kind)
	if err != nil {
		if ce.IsNotFoundError(err) {
			return nil, fmt.Errorf("%s: %w", cfg, ErrNotFound)
		}

		return nil, err
	}

	return s.source.Transform(ctx, cfg.Kind(), d)
}

// GetSecrets for service.
func (s *Service) GetSecrets(ctx context.Context, secs map[string]string) (map[string][]byte, error) {
	secrets := make(map[string][]byte, len(secs))

	for n, v := range secs {
		t, err := s.provider.Transform(ctx, v)
		if err != nil {
			return nil, err
		}

		switch c := t.(type) {
		case string:
			secrets[n] = []byte(c)
		default:
			secrets[n] = []byte(v)
		}
	}

	return secrets, nil
}

// NewConfig for service.
func NewConfig(app, ver, env, continent, country, cmd, kind string) (*Config, error) {
	if continent == "" {
		continent = "*"
	}

	if country == "" {
		country = "*"
	}

	if kind == "" {
		kind = "yaml"
	}

	if app == "" {
		return nil, fmt.Errorf("application: %w", ErrInvalidArgument)
	}

	if ver == "" {
		return nil, fmt.Errorf("version: %w", ErrInvalidArgument)
	}

	if env == "" {
		return nil, fmt.Errorf("environment: %w", ErrInvalidArgument)
	}

	if cmd == "" {
		return nil, fmt.Errorf("command: %w", ErrInvalidArgument)
	}

	c := &Config{
		application: app, version: ver, environment: env,
		continent: continent, country: country,
		command: cmd, kind: kind,
	}

	return c, nil
}

// Application for config.
func (c *Config) Application() string {
	return c.application
}

// Version for config.
func (c *Config) Version() string {
	return c.version
}

// Environment for config.
func (c *Config) Environment() string {
	return c.environment
}

// Continent for config.
func (c *Config) Continent() string {
	return c.continent
}

// Country for config.
func (c *Config) Country() string {
	return c.country
}

// Command for config.
func (c *Config) Command() string {
	return c.command
}

// Kind for config.
func (c *Config) Kind() string {
	return c.kind
}

// String of config.
func (c *Config) String() string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", c.application, c.version, c.environment, c.continent, c.country, c.command, c.kind)
}
