package mercadolibre

import (
	"context"
	"net/url"
	"strconv"
)

type SiteID string

// Site IDs for the markets where MercadoLibre operates.
const (
	SiteArgentina  SiteID = "MLA"
	SiteBrazil     SiteID = "MLB"
	SiteMexico     SiteID = "MLM"
	SiteChile      SiteID = "MLC"
	SiteColombia   SiteID = "MCO"
	SiteUruguay    SiteID = "MLU"
	SitePeru       SiteID = "MPE"
	SiteVenezuela  SiteID = "MLV"
	SiteEcuador    SiteID = "MEC"
	SiteBolivia    SiteID = "MBO"
	SiteParaguay   SiteID = "MPY"
	SiteGuatemala  SiteID = "MGT"
	SitePanama     SiteID = "MPA"
	SiteDominicana SiteID = "MRD"
	SiteCostaRica  SiteID = "MCR"
	SiteHonduras   SiteID = "MHN"
	SiteSalvador   SiteID = "MSV"
	SiteNicaragua  SiteID = "MNI"
)

type (
	// Site is an entry from GET /sites.
	Site struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		DefaultLang string   `json:"default_language"`
		CountryID   string   `json:"country_id"`
		SaleFormats []IDName `json:"sale_format,omitempty"`
	}

	// SiteDetail is the full response of GET /sites/{id}.
	SiteDetail struct {
		ID                string   `json:"id"`
		Name              string   `json:"name"`
		DefaultCurrencyID string   `json:"default_currency_id"`
		CountryID         string   `json:"country_id"`
		Currencies        []IDName `json:"currencies,omitempty"`
		Categories        []IDName `json:"categories,omitempty"`
		SaleFormats       []IDName `json:"sale_format,omitempty"`
		ListingTypes      []IDName `json:"listing_types,omitempty"`
	}

	// IDName is the ubiquitous {id, name} pair used across the API.
	IDName struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	// ListingType describes a listing type (gold_special, gold_pro, etc.).
	ListingType struct {
		DurationDays  *int    `json:"duration_days,omitempty"`
		ID            string  `json:"id"`
		Name          string  `json:"name"`
		SiteID        string  `json:"site_id,omitempty"`
		CategoryID    string  `json:"category_id,omitempty"`
		SaleFeeAmount float64 `json:"sale_fee_amount,omitempty"`
		BuyItNow      bool    `json:"buy_it_now,omitempty"`
		Auction       bool    `json:"auction,omitempty"`
	}

	// ListingPrice is the cost of publishing an item with a given listing type.
	ListingPrice struct {
		ListingTypeID   string  `json:"listing_type_id"`
		ListingTypeName string  `json:"listing_type_name,omitempty"`
		CategoryID      string  `json:"category_id,omitempty"`
		CurrencyID      string  `json:"currency_id,omitempty"`
		Amount          float64 `json:"amount,omitempty"`
		FreeRelistings  int     `json:"free_relistings,omitempty"`
	}

	// CategoryPrediction is one candidate category returned by the predictor.
	CategoryPrediction struct {
		ID             string  `json:"id"`
		Name           string  `json:"name,omitempty"`
		PredictionType string  `json:"prediction_type,omitempty"`
		Probability    float64 `json:"probability,omitempty"`
	}

	// SitesService accesses sites, categories listings and currencies.
	SitesService struct{ c *Client }
)

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
func (s *SitesService) ListingPrices(
	ctx context.Context,
	siteID string,
	price float64,
	categoryID string,
) ([]ListingPrice, error) {
	q := url.Values{string(QueryParamPrice): {strconv.FormatFloat(price, 'f', -1, 64)}}
	if categoryID != "" {
		q.Set(string(QueryParamCategoryID), categoryID)
	}
	return GetQ[[]ListingPrice](ctx, s.c, EPSiteListingPrices, q, siteID)
}

// PredictCategory uses ML's category predictor to suggest categories for a
// query string (GET /sites/{site_id}/category_predictor/predict?q=...).
func (s *SitesService) PredictCategory(ctx context.Context, siteID, query string) ([]CategoryPrediction, error) {
	return GetQ[[]CategoryPrediction](
		ctx,
		s.c,
		EPSiteCategoryPredictor,
		url.Values{string(QueryParamQ): {query}},
		siteID,
	)
}
