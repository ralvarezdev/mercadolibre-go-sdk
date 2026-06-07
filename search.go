package mercadolibre

import (
	"context"
	"net/url"
)

// SearchItem is an item as returned by the public site search
// (GET /sites/{id}/search). It is a leaner projection than Item.
type SearchItem struct {
	ID                string      `json:"id"`
	Title             string      `json:"title"`
	CategoryID        string      `json:"category_id"`
	Price             float64     `json:"price"`
	CurrencyID        string      `json:"currency_id"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Condition         string      `json:"condition"`
	Permalink         string      `json:"permalink"`
	Thumbnail         string      `json:"thumbnail"`
	SellerID          int64       `json:"seller_id,omitempty"`
	Seller            *IDNameUser `json:"seller,omitempty"`
	Shipping          *Shipping   `json:"shipping,omitempty"`
	Attributes        []Attribute `json:"attributes,omitempty"`
}

// IDNameUser is a seller projection embedded in search results.
type IDNameUser struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname,omitempty"`
}

// SiteSearchResponse is the response of GET /sites/{id}/search.
type SiteSearchResponse struct {
	SiteID           string       `json:"site_id"`
	Query            string       `json:"query,omitempty"`
	Paging           Paging       `json:"paging"`
	Results          []SearchItem `json:"results"`
	Sort             *Sort        `json:"sort,omitempty"`
	AvailableSorts   []Sort       `json:"available_sorts,omitempty"`
	Filters          []Filter     `json:"filters,omitempty"`
	AvailableFilters []Filter     `json:"available_filters,omitempty"`
}

// Search runs a public search over a site's listings (GET /sites/{id}/search).
// params carries the query (q, seller_id, nickname, category, sort, filters…)
// and pagination (offset, limit).
func (s *SitesService) Search(ctx context.Context, siteID string, params url.Values) (*SiteSearchResponse, error) {
	return GetQ[*SiteSearchResponse](ctx, s.c, EPSiteSearch, params, siteID)
}
