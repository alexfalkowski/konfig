package token

import (
	"github.com/alexfalkowski/go-service/token"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	token.Module,
	fx.Provide(NewVerifier),
)
