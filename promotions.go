package mercadolibre

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
)

// Promotion is a seller promotion/campaign invitation
// (GET /seller-promotions/users/{id}).
type Promotion struct {
	ID           string            `json:"id"`
	Type         string            `json:"type"` // MARKETPLACE_CAMPAIGN | VOLUME | DEAL | ...
	Status       string            `json:"status"`
	Name         string            `json:"name,omitempty"`
	StartDate    Time              `json:"start_date,omitempty"`
	FinishDate   Time              `json:"finish_date,omitempty"`
	DeadlineDate Time              `json:"deadline_date,omitempty"`
	Benefits     *PromotionBenefit `json:"benefits,omitempty"`
}

// PromotionBenefit describes the discount mechanics of a promotion.
type PromotionBenefit struct {
	Type                string  `json:"type,omitempty"`
	Name                string  `json:"name,omitempty"`
	MeliPercent         float64 `json:"meli_percent,omitempty"`
	SellerPercent       float64 `json:"seller_percent,omitempty"`
	BuyQuantity         int     `json:"buy_quantity,omitempty"`
	PayQuantity         int     `json:"pay_quantity,omitempty"`
	ItemDiscountPercent float64 `json:"item_discount_percent,omitempty"`
}

// PromotionsResponse is the response of listing a user's promotions.
type PromotionsResponse struct {
	Results []Promotion `json:"results"`
	Paging  Paging      `json:"paging"`
}

// PromotionItem is an item participating in (or candidate for) a promotion.
type PromotionItem struct {
	ID            string  `json:"id"`
	Status        string  `json:"status"`
	Price         float64 `json:"price,omitempty"`
	OriginalPrice float64 `json:"original_price,omitempty"`
}

// PromotionItemsResponse is the response of listing items of a promotion. The
// API returns paging as an object, {} or [] depending on the case, so it is
// decoded leniently.
type PromotionItemsResponse struct {
	Results []PromotionItem `json:"results"`
	Paging  loosePaging     `json:"paging"`
}

// loosePaging decodes a paging object, tolerating the empty-array form the
// promotions API sometimes returns.
type loosePaging struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (p *loosePaging) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || b[0] == '[' { // [] form: nothing to decode
		return nil
	}
	type alias loosePaging
	return json.Unmarshal(b, (*alias)(p))
}

// PromotionsService accesses seller promotions. All calls use app_version=v2.
type PromotionsService struct{ c *Client }

func promoQuery(extra url.Values) url.Values {
	q := url.Values{"app_version": {"v2"}}
	for k, v := range extra {
		q[k] = v
	}
	return q
}

// Get returns a promotion by ID
// (GET /seller-promotions/promotions/{id}?app_version=v2).
func (s *PromotionsService) Get(ctx context.Context, promotionID string) (*Promotion, error) {
	return GetQ[*Promotion](ctx, s.c, EPSellerPromotionByID, promoQuery(nil), promotionID)
}

// User lists a seller's promotion invitations
// (GET /seller-promotions/users/{id}?app_version=v2).
func (s *PromotionsService) User(ctx context.Context, userID int64, opts ListOptions) (*PromotionsResponse, error) {
	return GetQ[*PromotionsResponse](ctx, s.c, EPSellerPromotionsUser, promoQuery(opts.Values()), userID)
}

// Items lists the items of a promotion
// (GET /seller-promotions/promotions/{id}/items?promotion_type=...&app_version=v2).
// extra may add status, status_item, item_id, limit, search_after.
func (s *PromotionsService) Items(ctx context.Context, promotionID, promotionType string, extra url.Values) (*PromotionItemsResponse, error) {
	q := promoQuery(extra)
	q.Set("promotion_type", promotionType)
	return GetQ[*PromotionItemsResponse](ctx, s.c, EPSellerPromotionItems, q, promotionID)
}

// ItemPromotions lists the promotions an item is eligible for or participating in
// (GET /seller-promotions/items/{item_id}?app_version=v2).
func (s *PromotionsService) ItemPromotions(ctx context.Context, itemID string) ([]Promotion, error) {
	return GetQ[[]Promotion](ctx, s.c, EPSellerPromotionByItem, promoQuery(nil), itemID)
}

// EnrollItemRequest is the body of POST /seller-promotions/promotions/{id}/items.
type EnrollItemRequest struct {
	ItemID        string  `json:"item_id"`
	Price         float64 `json:"price,omitempty"`
	OriginalPrice float64 `json:"original_price,omitempty"`
}

// EnrollItemResponse is returned when an item is enrolled in a promotion.
type EnrollItemResponse struct {
	ItemID string `json:"item_id"`
	Status string `json:"status"`
}

// EnrollItem adds an item to an active promotion
// (POST /seller-promotions/promotions/{id}/items?app_version=v2).
func (s *PromotionsService) EnrollItem(ctx context.Context, promotionID string, req EnrollItemRequest) (*EnrollItemResponse, error) {
	return Do[*EnrollItemResponse, EnrollItemRequest](ctx, s.c, "POST", Request[EnrollItemRequest]{
		Endpoint: EPSellerPromotionItems,
		Args:     []any{promotionID},
		Query:    promoQuery(nil),
		Body:     &req,
	})
}

// UnenrollItem removes an item from a promotion
// (DELETE /seller-promotions/promotions/{id}/items/{item_id}?app_version=v2).
func (s *PromotionsService) UnenrollItem(ctx context.Context, promotionID, itemID string) error {
	_, err := Do[map[string]any, struct{}](ctx, s.c, "DELETE", Request[struct{}]{
		Endpoint: EPSellerPromotionItemByID,
		Args:     []any{promotionID, itemID},
		Query:    promoQuery(nil),
	})
	return err
}
