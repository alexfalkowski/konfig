package cmd

import (
	"github.com/alexfalkowski/go-service/compressor"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/cmd/config"
	kc "github.com/alexfalkowski/konfig/config"
	"go.uber.org/fx"
)

// ConfigOptions for cmd.
var ConfigOptions = []fx.Option{
	compressor.Module, marshaller.Module,
	telemetry.Module, metrics.Module,
	config.Module, client.Module,
	kc.Module, config.Module, Module,
}
