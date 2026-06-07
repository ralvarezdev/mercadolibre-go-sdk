package mercadolibre

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// Item is a MercadoLibre listing (GET /items/{id}).
//
// Some fields are owner-only: initial_quantity, available_quantity and
// sold_quantity are omitted when the item is read without the seller's token.
//
// Note: condition is being deprecated in favor of the item_condition attribute;
// and on public Items/search resources available_quantity is referential
// (bucketed ranges), per the docs.
type Item struct {
	ID                        string         `json:"id"`
	SiteID                    string         `json:"site_id"`
	Title                     string         `json:"title"`
	Subtitle                  *string        `json:"subtitle,omitempty"`
	SellerID                  int64          `json:"seller_id"`
	CategoryID                string         `json:"category_id"`
	UserProductID             string         `json:"user_product_id,omitempty"`
	OfficialStoreID           *int64         `json:"official_store_id,omitempty"`
	Price                     float64        `json:"price"`
	BasePrice                 float64        `json:"base_price,omitempty"`
	OriginalPrice             *float64       `json:"original_price,omitempty"`
	InventoryID               *string        `json:"inventory_id,omitempty"`
	CurrencyID                string         `json:"currency_id"`
	InitialQuantity           int            `json:"initial_quantity,omitempty"`
	AvailableQuantity         int            `json:"available_quantity,omitempty"`
	SoldQuantity              int            `json:"sold_quantity,omitempty"`
	SaleTerms                 []SaleTerm     `json:"sale_terms,omitempty"`
	BuyingMode                string         `json:"buying_mode,omitempty"`
	ListingTypeID             string         `json:"listing_type_id,omitempty"`
	StartTime                 Time           `json:"start_time,omitempty"`
	StopTime                  Time           `json:"stop_time,omitempty"`
	EndTime                   Time           `json:"end_time,omitempty"`
	ExpirationTime            Time           `json:"expiration_time,omitempty"`
	HistoricalStartTime       Time           `json:"historical_start_time,omitempty"`
	Condition                 string         `json:"condition,omitempty"`
	Permalink                 string         `json:"permalink,omitempty"`
	ThumbnailID               string         `json:"thumbnail_id,omitempty"`
	Thumbnail                 string         `json:"thumbnail,omitempty"`
	Pictures                  []Picture      `json:"pictures,omitempty"`
	VideoID                   *string        `json:"video_id,omitempty"`
	Descriptions              []ItemDescRef  `json:"descriptions,omitempty"`
	AcceptsMercadoPago        bool           `json:"accepts_mercadopago,omitempty"`
	NonMercadoPagoPayMethods  []any          `json:"non_mercado_pago_payment_methods,omitempty"`
	Shipping                  *Shipping      `json:"shipping,omitempty"`
	InternationalDeliveryMode string         `json:"international_delivery_mode,omitempty"`
	SellerAddress             *SellerAddress `json:"seller_address,omitempty"`
	Geolocation               *GeoLocation   `json:"geolocation,omitempty"`
	CoverageAreas             []any          `json:"coverage_areas,omitempty"`
	Attributes                []Attribute    `json:"attributes,omitempty"`
	Variations                []Variation    `json:"variations,omitempty"`
	Status                    string         `json:"status,omitempty"`
	SubStatus                 []string       `json:"sub_status,omitempty"`
	Tags                      []string       `json:"tags,omitempty"`
	Warranty                  string         `json:"warranty,omitempty"`
	CatalogProductID          *string        `json:"catalog_product_id,omitempty"`
	DomainID                  *string        `json:"domain_id,omitempty"`
	SellerCustomField         *string        `json:"seller_custom_field,omitempty"`
	ParentItemID              *string        `json:"parent_item_id,omitempty"`
	DealIDs                   []string       `json:"deal_ids,omitempty"`
	AutomaticRelist           bool           `json:"automatic_relist,omitempty"`
	DateCreated               Time           `json:"date_created,omitempty"`
	LastUpdated               Time           `json:"last_updated,omitempty"`
	Health                    *float64       `json:"health,omitempty"`
}

// ItemDescRef references an item description (the body lives at
// /items/{id}/description).
type ItemDescRef struct {
	ID string `json:"id"`
}

// Variation is an item variation (size/color/etc.).
type Variation struct {
	ID                int64       `json:"id"`
	Price             float64     `json:"price,omitempty"`
	AvailableQuantity int         `json:"available_quantity,omitempty"`
	SoldQuantity      int         `json:"sold_quantity,omitempty"`
	AttributeCombos   []Attribute `json:"attribute_combinations,omitempty"`
	Attributes        []Attribute `json:"attributes,omitempty"`
	PictureIDs        []string    `json:"picture_ids,omitempty"`
	SellerCustomField *string     `json:"seller_custom_field,omitempty"`
}

// ItemDescription is the plain/text description of an item.
type ItemDescription struct {
	Text      string `json:"text,omitempty"`
	PlainText string `json:"plain_text,omitempty"`
}

// CreateItemRequest is the body of POST /items (publish a new listing).
//
// Only the fields MercadoLibre requires for the target category need to be
// set. Refer to GET /categories/{id}/attributes for required attributes.
type CreateItemRequest struct {
	Title             string      `json:"title"`
	CategoryID        string      `json:"category_id"`
	Price             float64     `json:"price"`
	CurrencyID        string      `json:"currency_id"`
	AvailableQuantity int         `json:"available_quantity"`
	BuyingMode        string      `json:"buying_mode"` // BuyingModeBuyItNow etc.
	Condition         string      `json:"condition"`   // ItemConditionNew etc.
	ListingTypeID     string      `json:"listing_type_id"`
	Description       *struct {
		PlainText string `json:"plain_text"`
	} `json:"description,omitempty"`
	Warranty    string      `json:"warranty,omitempty"`
	SaleTerms   []SaleTerm  `json:"sale_terms,omitempty"`
	Attributes  []Attribute `json:"attributes,omitempty"`
	Variations  []Variation `json:"variations,omitempty"`
	Pictures    []Picture   `json:"pictures,omitempty"` // set Picture.Source to a public URL
	VideoID     *string     `json:"video_id,omitempty"`
	Shipping    *Shipping   `json:"shipping,omitempty"`
}

// UpdateItemRequest is the body of PUT /items/{id}. Only set the fields you
// want to change; zero/nil values are omitted and left unchanged by the API.
type UpdateItemRequest struct {
	Title             *string     `json:"title,omitempty"`
	CategoryID        *string     `json:"category_id,omitempty"`
	Price             *float64    `json:"price,omitempty"`
	OriginalPrice     *float64    `json:"original_price,omitempty"`
	BasePrice         *float64    `json:"base_price,omitempty"`
	AvailableQuantity *int        `json:"available_quantity,omitempty"`
	ListingTypeID     *string     `json:"listing_type_id,omitempty"`
	Condition         *string     `json:"condition,omitempty"`
	Status            *string     `json:"status,omitempty"` // ItemStatusActive/Paused/Closed
	Warranty          *string     `json:"warranty,omitempty"`
	Pictures          []Picture   `json:"pictures,omitempty"` // replaces all pictures when non-empty
	Attributes        []Attribute `json:"attributes,omitempty"`
	SaleTerms         []SaleTerm  `json:"sale_terms,omitempty"`
	Variations        []Variation `json:"variations,omitempty"`
	Shipping          *Shipping   `json:"shipping,omitempty"`
}

// UploadPictureRequest is the body of POST /pictures.
type UploadPictureRequest struct {
	Source string `json:"source"` // publicly accessible URL of the image
}

// ItemsService accesses item listings, search and descriptions.
type ItemsService struct{ c *Client }

// Create publishes a new item listing (POST /items).
// Returns the created item with its assigned ID.
func (s *ItemsService) Create(ctx context.Context, req CreateItemRequest) (*Item, error) {
	return Post[*Item, CreateItemRequest](ctx, s.c, EPItems, req)
}

// Update modifies an existing item (PUT /items/{id}).
// Only fields set (non-nil pointers / non-empty slices) in req are sent.
func (s *ItemsService) Update(ctx context.Context, itemID string, req UpdateItemRequest) (*Item, error) {
	return Put[*Item, UpdateItemRequest](ctx, s.c, EPItemByID, req, itemID)
}

// UpdateDescription replaces a listing's plain-text description
// (PUT /items/{id}/description).
func (s *ItemsService) UpdateDescription(ctx context.Context, itemID, plainText string) (*ItemDescription, error) {
	body := ItemDescription{PlainText: plainText}
	return Put[*ItemDescription, ItemDescription](ctx, s.c, EPItemDescription, body, itemID)
}

// UpdateStatus pauses, activates, or closes a listing.
// Use ItemStatusActive, ItemStatusPaused, or ItemStatusClosed.
func (s *ItemsService) UpdateStatus(ctx context.Context, itemID, status string) (*Item, error) {
	return s.Update(ctx, itemID, UpdateItemRequest{Status: &status})
}

// UpdatePrice changes an item's price.
func (s *ItemsService) UpdatePrice(ctx context.Context, itemID string, price float64) (*Item, error) {
	return s.Update(ctx, itemID, UpdateItemRequest{Price: &price})
}

// UpdateStock changes an item's available_quantity.
func (s *ItemsService) UpdateStock(ctx context.Context, itemID string, qty int) (*Item, error) {
	return s.Update(ctx, itemID, UpdateItemRequest{AvailableQuantity: &qty})
}

// UploadPictureFromURL registers a remote image URL with MercadoLibre and
// returns the assigned Picture (with ID). The URL must be publicly accessible;
// ML downloads and caches the image during this call.
//
// Attach the returned Picture.ID to an item by calling Update with a Pictures
// slice containing Picture{ID: id}.
func (s *ItemsService) UploadPictureFromURL(ctx context.Context, sourceURL string) (*Picture, error) {
	body := UploadPictureRequest{Source: sourceURL}
	return Post[*Picture, UploadPictureRequest](ctx, s.c, EPPictures, body)
}

// Get returns a single item (GET /items/{id}).
func (s *ItemsService) Get(ctx context.Context, itemID string) (*Item, error) {
	return Get[*Item](ctx, s.c, EPItemByID, itemID)
}

// GetMulti fetches up to 20 items in one call (GET /items?ids=...). Each result
// carries its own status code (verbose multiget format).
func (s *ItemsService) GetMulti(ctx context.Context, itemIDs ...string) ([]MultiGetResult[Item], error) {
	q := url.Values{"ids": {strings.Join(itemIDs, ",")}}
	return GetQ[[]MultiGetResult[Item]](ctx, s.c, EPItems, q)
}

// GetMultiFields is like GetMulti but selects only the given fields
// (GET /items?ids=...&attributes=...).
func (s *ItemsService) GetMultiFields(ctx context.Context, fields []string, itemIDs ...string) ([]MultiGetResult[Item], error) {
	q := url.Values{
		"ids":        {strings.Join(itemIDs, ",")},
		"attributes": {strings.Join(fields, ",")},
	}
	return GetQ[[]MultiGetResult[Item]](ctx, s.c, EPItems, q)
}

// Description returns an item's description (GET /items/{id}/description).
func (s *ItemsService) Description(ctx context.Context, itemID string) (*ItemDescription, error) {
	return Get[*ItemDescription](ctx, s.c, EPItemDescription, itemID)
}

// UserItemsSearchResponse is the response of /users/{id}/items/search. Results
// are item IDs; hydrate them with GetMulti.
type UserItemsSearchResponse struct {
	SellerID         string   `json:"seller_id"`
	Query            *string  `json:"query"`
	Paging           Paging   `json:"paging"`
	Results          []string `json:"results"`
	Filters          []Filter `json:"filters,omitempty"`
	AvailableFilters []Filter `json:"available_filters,omitempty"`
	AvailableOrders  []Sort   `json:"available_orders,omitempty"`
	ScrollID         string   `json:"scroll_id,omitempty"`
}

// SearchUserItems lists a seller's item IDs with offset/limit paging
// (GET /users/{id}/items/search). Use extra to add filters such as
// status=active, listing_type_id=gold_pro, sku=..., orders=start_time_desc.
func (s *ItemsService) SearchUserItems(ctx context.Context, userID int64, opts ListOptions, extra url.Values) (*UserItemsSearchResponse, error) {
	q := mergeValues(opts.Values(), extra)
	return GetQ[*UserItemsSearchResponse](ctx, s.c, EPUserItemsSearch, q, userID)
}

// ScanUserItems returns an iterator over ALL of a seller's item IDs using
// search_type=scan (scroll). Use this for more than 1000 results; do not combine
// with offset/limit. Consume continuously — the scroll expires (~5 min).
func (s *ItemsService) ScanUserItems(userID int64, limit int, extra url.Values) *ItemScanIterator {
	q := url.Values{}
	for k, v := range extra {
		q[k] = v
	}
	q.Set("search_type", "scan")
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	return &ItemScanIterator{c: s.c, userID: userID, q: q}
}

// ItemScanIterator walks scan/scroll pages of seller item IDs.
type ItemScanIterator struct {
	c        *Client
	userID   int64
	q        url.Values
	scrollID string
	done     bool
}

// Next returns the next page of item IDs. ok is false when iteration is
// complete (the API returns an empty result set / null scroll).
func (it *ItemScanIterator) Next(ctx context.Context) (ids []string, ok bool, err error) {
	if it.done {
		return nil, false, nil
	}
	q := url.Values{}
	for k, v := range it.q {
		q[k] = v
	}
	if it.scrollID != "" {
		q.Set("scroll_id", it.scrollID)
	}
	resp, err := GetQ[*UserItemsSearchResponse](ctx, it.c, EPUserItemsSearch, q, it.userID)
	if err != nil {
		return nil, false, err
	}
	it.scrollID = resp.ScrollID
	if len(resp.Results) == 0 {
		it.done = true
		return nil, false, nil
	}
	return resp.Results, true, nil
}

// GetVariations returns all variations of an item
// (GET /items/{id}/variations).
func (s *ItemsService) GetVariations(ctx context.Context, itemID string) ([]Variation, error) {
	return Get[[]Variation](ctx, s.c, EPItemVariations, itemID)
}

// CreateVariation adds a new variation to an item
// (POST /items/{id}/variations).
func (s *ItemsService) CreateVariation(ctx context.Context, itemID string, v Variation) (*Variation, error) {
	return Post[*Variation, Variation](ctx, s.c, EPItemVariations, v, itemID)
}

// UpdateVariation modifies a variation (PUT /items/{id}/variations/{variation_id}).
func (s *ItemsService) UpdateVariation(ctx context.Context, itemID string, variationID int64, v Variation) (*Variation, error) {
	return Put[*Variation, Variation](ctx, s.c, EPItemVariationByID, v, itemID, variationID)
}

// DeleteVariation removes a variation (DELETE /items/{id}/variations/{variation_id}).
func (s *ItemsService) DeleteVariation(ctx context.Context, itemID string, variationID int64) error {
	_, err := Delete[struct{}](ctx, s.c, EPItemVariationByID, itemID, variationID)
	return err
}

func mergeValues(a, b url.Values) url.Values {
	out := url.Values{}
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}
