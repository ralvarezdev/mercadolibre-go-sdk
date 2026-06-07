// Package auth implements MercadoLibre OAuth 2.0: building the authorization
// URL, exchanging an authorization code for a token, refreshing tokens (which
// rotate and are single-use), PKCE, and pluggable token sources/stores.
//
// It is self-contained and does not import the root mercadolibre package, so it
// can be used standalone (e.g. in an auth service) and avoids an import cycle.
package auth

import "time"

// Config holds the OAuth application credentials and flow parameters.
type Config struct {
	ClientID     string // APP ID
	ClientSecret string // Secret Key
	RedirectURI  string // must match the app config exactly (no variable parts)
	Site         string // e.g. "MLA"; selects the authorization host (see hosts.go)
}

// Token is an OAuth token response plus a computed absolute Expiry.
type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	UserID       int64  `json:"user_id"`
	RefreshToken string `json:"refresh_token"`

	// ExpiresIn is the lifetime in seconds as returned by the API.
	ExpiresIn int `json:"expires_in"`
	// Expiry is the absolute expiry time, computed from ExpiresIn at receipt.
	Expiry time.Time `json:"expiry,omitempty"`
}

// Valid reports whether the access token is non-empty and not expired, applying
// leeway so the token is refreshed slightly before it actually expires.
func (t *Token) Valid(leeway time.Duration) bool {
	if t == nil || t.AccessToken == "" {
		return false
	}
	if t.Expiry.IsZero() {
		return true // unknown expiry: assume usable
	}
	return time.Now().Add(leeway).Before(t.Expiry)
}

// setExpiry computes Expiry from ExpiresIn relative to now.
func (t *Token) setExpiry(now time.Time) {
	if t.ExpiresIn > 0 {
		t.Expiry = now.Add(time.Duration(t.ExpiresIn) * time.Second)
	}
}
