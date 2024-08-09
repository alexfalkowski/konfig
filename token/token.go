package token

import (
	"bytes"
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/security/token"
)

// ErrInvalid for token.
var ErrInvalid = errors.New("invalid token")

// NewVerifier for konfig.
func NewVerifier(cfg *Config) token.Verifier {
	return &Verifier{cfg: cfg}
}

// Verifier for konfig.
type Verifier struct {
	cfg *Config
}

// // Verify a token or error.
func (v *Verifier) Verify(ctx context.Context, token []byte) (context.Context, error) {
	d, err := os.ReadBase64File(v.cfg.Key)
	if err != nil {
		return ctx, err
	}

	if !bytes.Equal([]byte(d), token) {
		return ctx, ErrInvalid
	}

	return ctx, nil
}
