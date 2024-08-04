package cmd

import (
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/cmd/config"
	kc "github.com/alexfalkowski/konfig/config"
	"go.uber.org/fx"
)

// ConfigOptions for cmd.
var ConfigOptions = []fx.Option{
	compress.Module, encoding.Module,
	telemetry.Module, config.Module,
	client.Module, kc.Module,
	config.Module, Module,
}
