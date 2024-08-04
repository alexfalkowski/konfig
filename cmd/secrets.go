package cmd

import (
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/cmd/secrets"
	"github.com/alexfalkowski/konfig/config"
	"go.uber.org/fx"
)

// SecretsOptions for cmd.
var SecretsOptions = []fx.Option{
	compress.Module, encoding.Module,
	telemetry.Module, client.Module,
	secrets.Module, config.Module, Module,
}
