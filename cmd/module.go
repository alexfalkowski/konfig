package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto"
	"github.com/alexfalkowski/go-service/security"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	security.Module,
	crypto.Module,
	cmd.Module,
	fx.Provide(NewVersion),
)
