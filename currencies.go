package mercadolibre

import (
	"context"
	"net/url"
)

type (
	// Currency is a currency (GET /currencies, /currencies/{id}).
	Currency struct {
		ID            string `json:"id"`
		Description   string `json:"description"`
		Symbol        string `json:"symbol"`
		DecimalPlaces int    `json:"decimal_places"`
	}

	// CurrencyConversion is the result of /currency_conversions/search.
	CurrencyConversion struct {
		CreationDate Time    `json:"creation_date,omitempty"`
		ValidUntil   Time    `json:"valid_until,omitempty"`
		Ratio        float64 `json:"ratio"`
		Rate         float64 `json:"rate"`
	}

	// CurrenciesService accesses currencies and conversions.
	CurrenciesService struct{ c *Client }
)

// List returns all currencies (GET /currencies).
func (s *CurrenciesService) List(ctx context.Context) ([]Currency, error) {
	return Get[[]Currency](ctx, s.c, EPCurrencies)
}

// Get returns a currency by ID, e.g. "ARS" (GET /currencies/{id}).
func (s *CurrenciesService) Get(ctx context.Context, currencyID string) (*Currency, error) {
	return Get[*Currency](ctx, s.c, EPCurrencyByID, currencyID)
}

// Convert returns the conversion between two currencies
// (GET /currency_conversions/search?from=&to=).
func (s *CurrenciesService) Convert(ctx context.Context, from, to string) (*CurrencyConversion, error) {
	q := url.Values{string(QueryParamFrom): {from}, string(QueryParamTo): {to}}
	return GetQ[*CurrencyConversion](ctx, s.c, EPCurrencyConversions, q)
}
