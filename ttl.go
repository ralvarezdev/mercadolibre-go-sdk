package mercadolibre

import "time"

// TTL is a documented MercadoLibre lifetime. These are reference/default values
// taken from the developer docs; at runtime, token expiry MUST come from the
// token response's expires_in field rather than TTLAccessToken (the docs text
// says 6h while older examples report expires_in=10800).
type TTL = time.Duration

const (
	// TTLAccessToken is the documented access_token lifetime.
	TTLAccessToken TTL = 6 * time.Hour
	// TTLRefreshToken is the approximate refresh_token lifetime (~6 months).
	TTLRefreshToken TTL = 180 * 24 * time.Hour
	// TTLAuthCode is the short lifetime of an authorization code.
	TTLAuthCode TTL = 10 * time.Minute
	// TTLAppInactivity is the idle period after which a token is invalidated
	// (no API call in this window).
	TTLAppInactivity TTL = 4 * 30 * 24 * time.Hour
	// TTLScroll is the approximate scan/scroll_id window. Consume scroll pages
	// continuously; an expired scroll yields errors / 429.
	TTLScroll TTL = 5 * time.Minute
	// TTLTokenLeeway is how early a token is refreshed before its expiry.
	TTLTokenLeeway TTL = 5 * time.Minute
)
