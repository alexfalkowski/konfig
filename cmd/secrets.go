package cmd

import (
	"github.com/alexfalkowski/go-service/compressor"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	kc "github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/config"
	"go.uber.org/fx"
)

// SecretsOptions for cmd.
var SecretsOptions = []fx.Option{
	compressor.Module, marshaller.Module,
	telemetry.Module, metrics.Module,
	config.Module, kc.SecretsModule, Module,
}
