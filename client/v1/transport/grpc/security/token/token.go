package token

import (
	"github.com/alexfalkowski/go-service/security/token"
)

// NewGenerator for token.
func NewGenerator(tkn token.Tokenizer) token.Generator {
	return tkn
}
