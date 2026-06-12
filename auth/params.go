package auth

type ParamName string

// Query parameter keys for OAuth/auth endpoints.
const (
	ParamCodeChallenge       ParamName = "code_challenge"
	ParamCodeChallengeMethod ParamName = "code_challenge_method"
	ParamScope               ParamName = "scope"
	ParamResponseType        ParamName = "response_type"
	ParamClientID            ParamName = "client_id"
	ParamRedirectURI         ParamName = "redirect_uri"
	ParamState               ParamName = "state"
	ParamCodeVerifier        ParamName = "code_verifier"
	ParamGrantType           ParamName = "grant_type"
	ParamClientSecret        ParamName = "client_secret"
	ParamCode                ParamName = "code"
	ParamRefreshToken        ParamName = "refresh_token"
)

type GrantType string

// Grant types for OAuth token requests.
const (
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeRefreshToken      GrantType = "refresh_token"
)

type ResponseType string

// Response types for OAuth authorization requests.
const (
	ResponseTypeCode ResponseType = "code"
)

// URL paths and schemes for OAuth endpoints.
const (
	SchemeHTTPS       = "https"
	AuthorizationPath = "/authorization"
	TokenEndpoint     = "https://api.mercadolibre.com/oauth/token"
)
