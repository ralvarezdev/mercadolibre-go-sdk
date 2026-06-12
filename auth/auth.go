package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	// tokenURL is the token endpoint (same host for every site). It is a var so
	// tests can point it at a stub server.
	tokenURL = TokenEndpoint

	// httpClient is overridable in tests.
	httpClient = http.DefaultClient
)

// Error is an OAuth error response from the token endpoint.
type Error struct {
	Code        string `json:"error"`
	Description string `json:"error_description"`
	StatusCode  int    `json:"-"`
}

func (e *Error) Error() string {
	if e.Description != "" {
		return fmt.Sprintf("mercadolibre/auth: %d %s: %s", e.StatusCode, e.Code, e.Description)
	}
	return fmt.Sprintf("mercadolibre/auth: %d %s", e.StatusCode, e.Code)
}

// IsInvalidGrant reports whether the error is invalid_grant (code/refresh token
// expired or already used → the user must re-authorize).
func (e *Error) IsInvalidGrant() bool { return e.Code == "invalid_grant" }

// AuthURLOption customizes AuthCodeURL.
type AuthURLOption func(url.Values)

// WithPKCE adds the PKCE code challenge to the authorization URL.
func WithPKCE(p PKCE) AuthURLOption {
	return func(v url.Values) {
		v.Set(string(ParamCodeChallenge), p.Challenge)
		v.Set(string(ParamCodeChallengeMethod), p.Method)
	}
}

// WithScope sets the requested scopes (default: the app's configured scopes).
// Valid values: "offline_access", "read", "write".
func WithScope(scopes ...string) AuthURLOption {
	return func(v url.Values) { v.Set(string(ParamScope), strings.Join(scopes, " ")) }
}

// AuthCodeURL builds the URL to which the user is redirected to authorize the
// application. state should be a random, unguessable value you verify on
// callback.
func (c Config) AuthCodeURL(state string, opts ...AuthURLOption) string {
	v := url.Values{}
	v.Set(string(ParamResponseType), string(ResponseTypeCode))
	v.Set(string(ParamClientID), c.ClientID)
	v.Set(string(ParamRedirectURI), c.RedirectURI)
	if state != "" {
		v.Set(string(ParamState), state)
	}
	for _, opt := range opts {
		opt(v)
	}
	return (&url.URL{
		Scheme:   SchemeHTTPS,
		Host:     AuthHost(c.Site),
		Path:     AuthorizationPath,
		RawQuery: v.Encode(),
	}).String()
}

// ExchangeOption customizes Exchange.
type ExchangeOption func(url.Values)

// WithCodeVerifier supplies the PKCE code verifier matching the challenge used
// in AuthCodeURL.
func WithCodeVerifier(verifier string) ExchangeOption {
	return func(v url.Values) { v.Set(string(ParamCodeVerifier), verifier) }
}

// Exchange trades an authorization code for an access/refresh token.
func (c Config) Exchange(ctx context.Context, code string, opts ...ExchangeOption) (*Token, error) {
	v := url.Values{}
	v.Set(string(ParamGrantType), string(GrantTypeAuthorizationCode))
	v.Set(string(ParamClientID), c.ClientID)
	v.Set(string(ParamClientSecret), c.ClientSecret)
	v.Set(string(ParamCode), code)
	v.Set(string(ParamRedirectURI), c.RedirectURI)
	for _, opt := range opts {
		opt(v)
	}
	return c.token(ctx, v)
}

// Refresh exchanges a (single-use, rotating) refresh token for a new token. The
// returned Token carries a NEW refresh token that must be persisted; the old one
// is now invalid.
func (c Config) Refresh(ctx context.Context, refreshToken string) (*Token, error) {
	v := url.Values{}
	v.Set(string(ParamGrantType), string(GrantTypeRefreshToken))
	v.Set(string(ParamClientID), c.ClientID)
	v.Set(string(ParamClientSecret), c.ClientSecret)
	v.Set(string(ParamRefreshToken), refreshToken)
	return c.token(ctx, v)
}

func (c Config) token(ctx context.Context, form url.Values) (*Token, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set(HeaderAccept, string(ContentTypeJSON))
	req.Header.Set(HeaderContentType, string(ContentTypeFormURLEncoded))

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		oerr := &Error{StatusCode: resp.StatusCode}
		_ = json.Unmarshal(body, oerr)
		if oerr.Code == "" {
			oerr.Code = http.StatusText(resp.StatusCode)
		}
		return nil, oerr
	}

	var tok Token
	if err := json.Unmarshal(body, &tok); err != nil {
		return nil, fmt.Errorf("mercadolibre/auth: decoding token: %w", err)
	}
	tok.setExpiry(time.Now())
	return &tok, nil
}
