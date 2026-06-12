// Package mercadolibre is a Go client for the MercadoLibre marketplace APIs
// (https://api.mercadolibre.com).
//
// Create a client with a token source, then call typed services
// (c.Users, c.Items, …) or the generic helpers (Get[T], Post[T,B], …) for any
// endpoint that does not yet have a typed wrapper.
package mercadolibre

import (
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"

	"github.com/ralvarezdev/mercadolibre-go-sdk/auth"
)

const (
	// DefaultBaseURL is the MercadoLibre REST API root.
	DefaultBaseURL = "https://api.mercadolibre.com"

	// DefaultUserAgent is sent unless overridden with WithUserAgent.
	DefaultUserAgent = "mercadolibre-go-sdk"
)

type (
	// Client is a MercadoLibre API client. It is safe for concurrent use.
	Client struct {
		tokens     auth.TokenSource
		Messaging  *MessagingService
		Items      *ItemsService
		Categories *CategoriesService
		Domains    *DomainsService
		Currencies *CurrenciesService
		Locations  *LocationsService
		Orders     *OrdersService
		Shipments  *ShipmentsService
		Users      *UsersService
		Questions  *QuestionsService
		Promotions *PromotionsService
		Metrics    *MetricsService
		Claims     *ClaimsService
		Billing    *BillingService
		http       *http.Client
		Sites      *SitesService
		limiter    *rate.Limiter
		baseURL    *url.URL
		userAgent  string
		retry      RetryPolicy
	}

	// Option configures a Client.
	Option func(*Client)
)

// WithHTTPClient sets the underlying *http.Client (timeouts, transport, proxy).
func WithHTTPClient(h *http.Client) Option { return func(c *Client) { c.http = h } }

// WithTokenSource sets how access tokens are obtained and refreshed.
func WithTokenSource(ts auth.TokenSource) Option { return func(c *Client) { c.tokens = ts } }

// WithStaticToken is a shortcut for WithTokenSource(auth.StaticTokenSource(tok)).
func WithStaticToken(accessToken string) Option {
	return func(c *Client) { c.tokens = auth.StaticTokenSource(accessToken) }
}

// WithUserAgent overrides the User-Agent header.
func WithUserAgent(ua string) Option { return func(c *Client) { c.userAgent = ua } }

// WithBaseURL overrides the API base URL (primarily for tests).
func WithBaseURL(raw string) Option {
	return func(c *Client) {
		if u, err := url.Parse(raw); err == nil {
			c.baseURL = u
		}
	}
}

// WithRateLimiter sets a client-side rate limiter. MercadoLibre applies limits
// primarily per application (Client ID), so one limiter per Client is the right
// granularity.
func WithRateLimiter(l *rate.Limiter) Option { return func(c *Client) { c.limiter = l } }

// WithRequestsPerSecond is a convenience for WithRateLimiter with the given
// steady rate and burst.
func WithRequestsPerSecond(rps float64, burst int) Option {
	return func(c *Client) { c.limiter = rate.NewLimiter(rate.Limit(rps), burst) }
}

// WithRetryPolicy customizes retry/backoff behavior on 429 and 5xx responses.
func WithRetryPolicy(p RetryPolicy) Option { return func(c *Client) { c.retry = p } }

// NewClient creates a Client. Without WithTokenSource/WithStaticToken, only
// public (unauthenticated) endpoints will work.
func NewClient(opts ...Option) *Client {
	base, _ := url.Parse(DefaultBaseURL)
	c := &Client{
		baseURL:   base,
		http:      &http.Client{Timeout: 30 * time.Second},
		userAgent: DefaultUserAgent,
		retry:     DefaultRetryPolicy(),
	}
	for _, opt := range opts {
		opt(c)
	}
	c.Users = &UsersService{c: c}
	c.Sites = &SitesService{c: c}
	c.Items = &ItemsService{c: c}
	c.Categories = &CategoriesService{c: c}
	c.Domains = &DomainsService{c: c}
	c.Currencies = &CurrenciesService{c: c}
	c.Locations = &LocationsService{c: c}
	c.Orders = &OrdersService{c: c}
	c.Shipments = &ShipmentsService{c: c}
	c.Questions = &QuestionsService{c: c}
	c.Messaging = &MessagingService{c: c}
	c.Claims = &ClaimsService{c: c}
	c.Metrics = &MetricsService{c: c}
	c.Promotions = &PromotionsService{c: c}
	c.Billing = &BillingService{c: c}
	return c
}
