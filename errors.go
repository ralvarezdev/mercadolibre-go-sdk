package mercadolibre

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type (
	// Cause is one structured cause inside an API error response. MercadoLibre
	// populates it inconsistently (sometimes empty, sometimes objects), so callers
	// should treat it as best-effort detail.
	Cause struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}

	// APIError represents a non-2xx response from the MercadoLibre API. It decodes
	// the three error shapes observed across the docs into a single type:
	//
	//	{ "error":"invalid_grant", "error_description":"…", "status":400, "cause":[] }
	//	{ "status":403, "error":"Invalid scopes", "code":"FORBIDDEN" }
	//	{ "status":403, "error":"access_denied", "message":"…", "code":"FORBIDDEN" }
	APIError struct {
		Header     http.Header
		Code       string
		ErrName    string
		Message    string
		Cause      []Cause
		Raw        []byte
		StatusCode int
	}
)

func (e *APIError) Error() string {
	msg := e.Message
	if msg == "" {
		msg = e.ErrName
	}
	switch {
	case e.Code != "" && msg != "":
		return fmt.Sprintf("mercadolibre: %d %s: %s", e.StatusCode, e.Code, msg)
	case msg != "":
		return fmt.Sprintf("mercadolibre: %d: %s", e.StatusCode, msg)
	default:
		return fmt.Sprintf("mercadolibre: %d %s", e.StatusCode, http.StatusText(e.StatusCode))
	}
}

// IsRateLimited reports whether the request was throttled (HTTP 429).
func (e *APIError) IsRateLimited() bool { return e.StatusCode == http.StatusTooManyRequests }

// IsForbidden reports whether the request was forbidden (HTTP 403). See the
// docs page error-403 for the common causes (scopes, blocked IP, inactive user).
func (e *APIError) IsForbidden() bool { return e.StatusCode == http.StatusForbidden }

// IsUnauthorized reports whether the access token was missing/invalid (HTTP 401).
func (e *APIError) IsUnauthorized() bool { return e.StatusCode == http.StatusUnauthorized }

// IsNotFound reports whether the resource was not found (HTTP 404).
func (e *APIError) IsNotFound() bool { return e.StatusCode == http.StatusNotFound }

// IsInvalidGrant reports whether this is an OAuth invalid_grant error, meaning
// the authorization code or refresh token expired or was already used and the
// user must re-authorize.
func (e *APIError) IsInvalidGrant() bool { return e.ErrName == "invalid_grant" }

// IsAPIError extracts an *APIError from err, if present.
func IsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}

// parseAPIError builds an *APIError from a response body. It never fails: if the
// body is not the expected JSON shape, the raw body is preserved.
func parseAPIError(status int, body []byte, header http.Header) *APIError {
	e := &APIError{StatusCode: status, Raw: body, Header: header}

	// Decode into an intermediate so a malformed "cause" never breaks parsing.
	var raw struct {
		Description string          `json:"error_description"`
		Error       string          `json:"error"`
		Code        string          `json:"code"`
		Message     string          `json:"message"`
		Cause       json.RawMessage `json:"cause"`
		Status      int             `json:"status"`
	}
	if err := json.Unmarshal(body, &raw); err != nil {
		return e
	}

	e.ErrName = raw.Error
	e.Code = raw.Code
	if raw.Message != "" {
		e.Message = raw.Message
	} else {
		e.Message = raw.Description
	}
	if raw.Status != 0 {
		e.StatusCode = raw.Status
	}
	if len(raw.Cause) == 0 {
		return e
	}

	// Try []Cause first, then []string.
	var cs []Cause
	if json.Unmarshal(raw.Cause, &cs) == nil {
		e.Cause = cs
		return e
	}
	var ss []string
	if json.Unmarshal(raw.Cause, &ss) != nil {
		return e
	}
	for _, s := range ss {
		e.Cause = append(e.Cause, Cause{Message: s})
	}
	return e
}
