package mercadolibre

const (
	HeaderAccept         = "Accept"
	HeaderContentType    = "Content-Type"
	HeaderUserAgent      = "User-Agent"
	HeaderAuthorization  = "Authorization"
	HeaderAcceptLanguage = "Accept-Language"

	ContentTypeJSON           = "application/json"
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"

	AcceptAll           = "*/*"
	AuthorizationBearer = "Bearer "
)

type ContentType string
