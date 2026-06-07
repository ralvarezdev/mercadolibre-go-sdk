package mercadolibre

import (
	"context"
	"net/url"
	"strconv"
)

// Site IDs for the markets where MercadoLibre operates.
const (
	SiteArgentina  = "MLA"
	SiteBrazil     = "MLB"
	SiteMexico     = "MLM"
	SiteChile      = "MLC"
	SiteColombia   = "MCO"
	SiteUruguay    = "MLU"
	SitePeru       = "MPE"
	SiteVenezuela  = "MLV"
	SiteEcuador    = "MEC"
	SiteBolivia    = "MBO"
	SiteParaguay   = "MPY"
	SiteGuatemala  = "MGT"
	SitePanama     = "MPA"
	SiteDominicana = "MRD"
	SiteCostaRica  = "MCR"
	SiteHonduras   = "MHN"
	SiteSalvador   = "MSV"
	SiteNicaragua  = "MNI"
)

// Site is an entry from GET /sites.
type Site struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DefaultLang  string `json:"default_language"`
	CountryID    string `json:"country_id"`
	SaleFormats  []IDName `json:"sale_format,omitempty"`
}

// SiteDetail is the full response of GET /sites/{id}.
type SiteDetail struct {
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	DefaultCurrencyID  string      `json:"default_currency_id"`
	CountryID          string      `json:"country_id"`
	Currencies         []IDName    `json:"currencies,omitempty"`
	Categories         []IDName    `json:"categories,omitempty"`
	SaleFormats        []IDName    `json:"sale_format,omitempty"`
	ListingTypes       []IDName    `json:"listing_types,omitempty"`
}

// IDName is the ubiquitous {id, name} pair used across the API.
type IDName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListingType describes a listing type (gold_special, gold_pro, etc.).
type ListingType struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	SiteID        string  `json:"site_id,omitempty"`
	CategoryID    string  `json:"category_id,omitempty"`
	DurationDays  *int    `json:"duration_days,omitempty"`
	BuyItNow      bool    `json:"buy_it_now,omitempty"`
	Auction       bool    `json:"auction,omitempty"`
	SaleFeeAmount float64 `json:"sale_fee_amount,omitempty"`
}

// ListingPrice is the cost of publishing an item with a given listing type.
type ListingPrice struct {
	ListingTypeID   string  `json:"listing_type_id"`
	ListingTypeName string  `json:"listing_type_name,omitempty"`
	CategoryID      string  `json:"category_id,omitempty"`
	CurrencyID      string  `json:"currency_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	FreeRelistings  int     `json:"free_relistings,omitempty"`
}

// CategoryPrediction is one candidate category returned by the predictor.
type CategoryPrediction struct {
	ID             string  `json:"id"`
	Name           string  `json:"name,omitempty"`
	Probability    float64 `json:"probability,omitempty"`
	PredictionType string  `json:"prediction_type,omitempty"`
}

// SitesService accesses sites, categories listings and currencies.
type SitesService struct{ c *Client }

// List returns all MercadoLibre sites (GET /sites).
func (s *SitesService) List(ctx context.Context) ([]Site, error) {
	return Get[[]Site](ctx, s.c, EPSites)
}

// Get returns details for a single site (GET /sites/{id}).
func (s *SitesService) Get(ctx context.Context, siteID string) (*SiteDetail, error) {
	return Get[*SiteDetail](ctx, s.c, EPSiteByID, siteID)
}

// ListingTypes returns all listing types available on a site
// (GET /sites/{site_id}/listing_types).
func (s *SitesService) ListingTypes(ctx context.Context, siteID string) ([]ListingType, error) {
	return Get[[]ListingType](ctx, s.c, EPSiteListingTypes, siteID)
}

// ListingType returns a single listing type
// (GET /sites/{site_id}/listing_types/{type_id}).
func (s *SitesService) ListingType(ctx context.Context, siteID, typeID string) (*ListingType, error) {
	return Get[*ListingType](ctx, s.c, EPSiteListingTypeByID, siteID, typeID)
}

// ListingPrices returns the publishing cost for each listing type for a given
// price and optional category (GET /sites/{site_id}/listing_prices).
// price > 0 is required; categoryID may be empty.
func (s *SitesService) ListingPrices(ctx context.Context, siteID string, price float64, categoryID string) ([]ListingPrice, error) {
	q := url.Values{"price": {strconv.FormatFloat(price, 'f', -1, 64)}}
	if categoryID != "" {
		q.Set("category_id", categoryID)
	}
	return GetQ[[]ListingPrice](ctx, s.c, EPSiteListingPrices, q, siteID)
}

// PredictCategory uses ML's category predictor to suggest categories for a
// query string (GET /sites/{site_id}/category_predictor/predict?q=...).
func (s *SitesService) PredictCategory(ctx context.Context, siteID, query string) ([]CategoryPrediction, error) {
	return GetQ[[]CategoryPrediction](ctx, s.c, EPSiteCategoryPredictor, url.Values{"q": {query}}, siteID)
}
