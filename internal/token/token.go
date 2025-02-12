package token

import (
	"github.com/alexfalkowski/go-service/token"
)

// NewVerifier for token.
func NewVerifier(token *token.Token) token.Verifier {
	return token
}
