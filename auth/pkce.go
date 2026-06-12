package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const pkceMethod = "S256"

// PKCE holds a code verifier and its derived S256 challenge.
type PKCE struct {
	Verifier  string
	Challenge string
	Method    string // always "S256"
}

// GeneratePKCE creates a cryptographically random code verifier and its SHA-256
// (S256) challenge, per RFC 7636. Use Challenge in AuthCodeURL (WithPKCE) and
// Verifier in Exchange (WithCodeVerifier).
func GeneratePKCE() (PKCE, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return PKCE{}, err
	}
	verifier := base64.RawURLEncoding.EncodeToString(b)
	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sum[:])
	return PKCE{Verifier: verifier, Challenge: challenge, Method: pkceMethod}, nil
}
