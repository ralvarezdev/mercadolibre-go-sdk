package mercadolibre

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	// RetryPolicy controls automatic retries on 429 and 5xx responses (and transient
	// network errors). Set MaxAttempts to 1 to disable retries.
	RetryPolicy struct {
		MaxAttempts int           // total attempts including the first
		BaseDelay   time.Duration // initial backoff
		MaxDelay    time.Duration // cap on backoff
	}

	// Request is a typed request with a JSON body of type B. Use struct{} for B when
	// there is no body.
	Request[B any] struct {
		Query    url.Values
		Header   http.Header
		Body     *B
		Endpoint Endpoint
		Args     []any
	}
)

// DefaultRetryPolicy returns sensible defaults: 4 attempts, exponential backoff
// from 500ms capped at 10s, with jitter; honors Retry-After when present.
func DefaultRetryPolicy() RetryPolicy {
	return RetryPolicy{MaxAttempts: 4, BaseDelay: 500 * time.Millisecond, MaxDelay: 10 * time.Second}
}

// Do is the generic request core: typed body B in, typed value T out.
func Do[T, B any](ctx context.Context, c *Client, method string, r Request[B]) (T, error) {
	var zero T
	var body []byte
	if r.Body != nil {
		b, err := json.Marshal(r.Body)
		if err != nil {
			return zero, fmt.Errorf("mercadolibre: encoding request body: %w", err)
		}
		body = b
	}
	out := new(T)
	if err := c.do(ctx, method, r.Endpoint.Path(r.Args...), r.Query, body, r.Header, out); err != nil {
		return zero, err
	}
	return *out, nil
}

// Get performs a GET and decodes the JSON response into T.
func Get[T any](ctx context.Context, c *Client, ep Endpoint, args ...any) (T, error) {
	return Do[T, struct{}](ctx, c, http.MethodGet, Request[struct{}]{Endpoint: ep, Args: args})
}

// GetQ is Get with a query string.
func GetQ[T any](ctx context.Context, c *Client, ep Endpoint, q url.Values, args ...any) (T, error) {
	return Do[T, struct{}](ctx, c, http.MethodGet, Request[struct{}]{Endpoint: ep, Query: q, Args: args})
}

// Post performs a POST with a typed JSON body and decodes the response into T.
func Post[T, B any](ctx context.Context, c *Client, ep Endpoint, body B, args ...any) (T, error) {
	return Do[T, B](ctx, c, http.MethodPost, Request[B]{Endpoint: ep, Body: &body, Args: args})
}

// Put performs a PUT with a typed JSON body and decodes the response into T.
func Put[T, B any](ctx context.Context, c *Client, ep Endpoint, body B, args ...any) (T, error) {
	return Do[T, B](ctx, c, http.MethodPut, Request[B]{Endpoint: ep, Body: &body, Args: args})
}

// Delete performs a DELETE and decodes any response into T.
func Delete[T any](ctx context.Context, c *Client, ep Endpoint, args ...any) (T, error) {
	return Do[T, struct{}](ctx, c, http.MethodDelete, Request[struct{}]{Endpoint: ep, Args: args})
}

// GetRaw performs a GET and returns the raw response body without JSON
// decoding. Use this for endpoints that return binary or non-JSON content
// (e.g. billing report downloads).
func GetRaw(ctx context.Context, c *Client, ep Endpoint, q url.Values, args ...any) ([]byte, error) {
	return doRaw(ctx, c, http.MethodGet, ep.Path(args...), q)
}

// doRaw is like do but returns the raw response bytes instead of decoding JSON.
func doRaw(ctx context.Context, c *Client, method, path string, q url.Values) ([]byte, error) {
	u := *c.baseURL
	u.Path = strings.TrimRight(u.Path, "/") + path
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var token string
	if c.tokens != nil {
		t, err := c.tokens.Token(ctx)
		if err != nil {
			return nil, fmt.Errorf("mercadolibre: obtaining access token: %w", err)
		}
		token = t
	}

	attempts := c.retry.MaxAttempts
	if attempts < 1 {
		attempts = 1
	}

	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		if c.limiter != nil {
			if err := c.limiter.Wait(ctx); err != nil {
				return nil, err
			}
		}

		req, err := http.NewRequestWithContext(ctx, method, u.String(), http.NoBody)
		if err != nil {
			return nil, err
		}
		req.Header.Set(HeaderAccept, AcceptAll)
		if c.userAgent != "" {
			req.Header.Set(HeaderUserAgent, c.userAgent)
		}
		if token != "" {
			req.Header.Set(HeaderAuthorization, AuthorizationBearer+token)
		}

		resp, err := c.http.Do(req)
		if err != nil {
			lastErr = err
			if attempt < attempts && ctx.Err() == nil {
				if werr := wait(ctx, c.retry.backoff(attempt, 0)); werr != nil {
					return nil, werr
				}
				continue
			}
			return nil, err
		}

		body, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			return nil, readErr
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return body, nil
		}

		apiErr := parseAPIError(resp.StatusCode, body, resp.Header)
		lastErr = apiErr
		if attempt < attempts && retryable(resp.StatusCode) {
			if werr := wait(ctx, c.retry.backoff(attempt, retryAfter(resp.Header))); werr != nil {
				return nil, werr
			}
			continue
		}
		return nil, apiErr
	}
	return nil, lastErr
}

// do executes a request with auth injection, rate limiting, retries/backoff and
// error decoding, unmarshaling a 2xx JSON body into out (when non-nil).
func (c *Client) do(
	ctx context.Context,
	method, path string,
	q url.Values,
	body []byte,
	header http.Header,
	out any,
) error {
	u := *c.baseURL
	u.Path = strings.TrimRight(u.Path, "/") + path
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var token string
	if c.tokens != nil {
		t, err := c.tokens.Token(ctx)
		if err != nil {
			return fmt.Errorf("mercadolibre: obtaining access token: %w", err)
		}
		token = t
	}

	attempts := c.retry.MaxAttempts
	if attempts < 1 {
		attempts = 1
	}

	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		if c.limiter != nil {
			if err := c.limiter.Wait(ctx); err != nil {
				return err
			}
		}

		req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewReader(body))
		if err != nil {
			return err
		}
		req.Header.Set(HeaderAccept, string(ContentTypeJSON))
		if len(body) > 0 {
			req.Header.Set(HeaderContentType, string(ContentTypeJSON))
		}
		if c.userAgent != "" {
			req.Header.Set(HeaderUserAgent, c.userAgent)
		}
		if token != "" {
			req.Header.Set(HeaderAuthorization, AuthorizationBearer+token)
		}
		for k, vs := range header {
			for _, v := range vs {
				req.Header.Add(k, v)
			}
		}

		resp, err := c.http.Do(req)
		if err != nil {
			// Transient network error: retry within policy.
			lastErr = err
			if attempt < attempts && ctx.Err() == nil {
				if werr := wait(ctx, c.retry.backoff(attempt, 0)); werr != nil {
					return werr
				}
				continue
			}
			return err
		}

		respBody, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			return readErr
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			if out != nil && len(respBody) > 0 {
				if err := json.Unmarshal(respBody, out); err != nil {
					return fmt.Errorf("mercadolibre: decoding response: %w", err)
				}
			}
			return nil
		}

		apiErr := parseAPIError(resp.StatusCode, respBody, resp.Header)
		lastErr = apiErr
		if attempt < attempts && retryable(resp.StatusCode) {
			if werr := wait(ctx, c.retry.backoff(attempt, retryAfter(resp.Header))); werr != nil {
				return werr
			}
			continue
		}
		return apiErr
	}
	return lastErr
}

func retryable(status int) bool {
	return status == http.StatusTooManyRequests || status >= 500
}

// backoff returns the delay before the next attempt. If retryAfter > 0 it takes
// precedence; otherwise exponential backoff with full jitter, capped at MaxDelay.
func (p RetryPolicy) backoff(attempt int, retryAfter time.Duration) time.Duration {
	if retryAfter > 0 {
		return retryAfter
	}
	d := float64(p.BaseDelay) * math.Pow(2, float64(attempt-1))
	maxD := float64(p.MaxDelay)
	if d > maxD {
		d = maxD
	}
	return time.Duration(rand.Int63n(int64(d) + 1))
}

func retryAfter(h http.Header) time.Duration {
	v := h.Get("Retry-After")
	if v == "" {
		return 0
	}
	if secs, err := strconv.Atoi(v); err == nil {
		return time.Duration(secs) * time.Second
	}
	if t, err := http.ParseTime(v); err == nil {
		if d := time.Until(t); d > 0 {
			return d
		}
	}
	return 0
}

func wait(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
