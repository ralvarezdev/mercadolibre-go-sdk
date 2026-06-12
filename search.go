package mercadolibre

import (
	"context"
	"net/url"
)

// SearchItem is an item as returned by the public site search
// (GET /sites/{id}/search). It is a leaner projection than Item.
type SearchItem struct {
	Seller            *IDNameUser `json:"seller,omitempty"`
	Shipping          *Shipping   `json:"shipping,omitempty"`
	CurrencyID        string      `json:"currency_id"`
	Thumbnail         string      `json:"thumbnail"`
	Permalink         string      `json:"permalink"`
	Condition         string      `json:"condition"`
	ID                string      `json:"id"`
	Title             string      `json:"title"`
	CategoryID        string      `json:"category_id"`
	Attributes        []Attribute `json:"attributes,omitempty"`
	Price             float64     `json:"price"`
	SoldQuantity      int         `json:"sold_quantity"`
	AvailableQuantity int         `json:"available_quantity"`
	SellerID          int64       `json:"seller_id,omitempty"`
}

// IDNameUser is a seller projection embedded in search results.
type IDNameUser struct {
	Nickname string `json:"nickname,omitempty"`
	ID       int64  `json:"id"`
}

// SiteSearchResponse is the response of GET /sites/{id}/search.
type SiteSearchResponse struct {
	Sort             *Sort        `json:"sort,omitempty"`
	SiteID           string       `json:"site_id"`
	Query            string       `json:"query,omitempty"`
	AvailableFilters []Filter     `json:"available_filters,omitempty"`
	AvailableSorts   []Sort       `json:"available_sorts,omitempty"`
	Filters          []Filter     `json:"filters,omitempty"`
	Results          []SearchItem `json:"results"`
	Paging           Paging       `json:"paging"`
}

// Search runs a public search over a site's listings (GET /sites/{id}/search).
// params carries the query (q, seller_id, nickname, category, sort, filters…)
// and pagination (offset, limit).
func (s *SitesService) Search(ctx context.Context, siteID string, params url.Values) (*SiteSearchResponse, error) {
	return GetQ[*SiteSearchResponse](ctx, s.c, EPSiteSearch, params, siteID)
}
