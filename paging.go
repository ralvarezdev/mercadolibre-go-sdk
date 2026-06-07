package mercadolibre

import (
	"net/url"
	"strconv"
)

// Paging mirrors the "paging" object returned by offset/limit list endpoints.
type Paging struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// ListOptions are the common offset/limit pagination parameters. Do not combine
// offset/limit pagination with scan/scroll on the same request.
type ListOptions struct {
	Offset int
	Limit  int
}

// Values renders the options as query parameters (omitting zero values).
func (o ListOptions) Values() url.Values {
	v := url.Values{}
	o.apply(v)
	return v
}

func (o ListOptions) apply(v url.Values) {
	if o.Offset > 0 {
		v.Set("offset", strconv.Itoa(o.Offset))
	}
	if o.Limit > 0 {
		v.Set("limit", strconv.Itoa(o.Limit))
	}
}

func itoa(n int64) string { return strconv.FormatInt(n, 10) }
