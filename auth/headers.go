package auth

const (
	HeaderAccept      = "Accept"
	HeaderContentType = "Content-Type"
)

type ContentType string

const (
	ContentTypeJSON           ContentType = "application/json"
	ContentTypeFormURLEncoded ContentType = "application/x-www-form-urlencoded"
)
