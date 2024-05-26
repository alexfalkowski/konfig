package token

import (
	"github.com/alexfalkowski/go-service/security/token"
)

// NewVerifier for token.
func NewVerifier(tkn token.Tokenizer) token.Verifier {
	return tkn
}
