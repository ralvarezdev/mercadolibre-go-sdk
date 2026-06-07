package mercadolibre

import (
	"context"
	"net/url"
)

// Order is a sale (GET /orders/{id}).
type Order struct {
	ID             int64        `json:"id"`
	Status         string       `json:"status"`
	StatusDetail   *string      `json:"status_detail,omitempty"`
	DateCreated    Time         `json:"date_created"`
	DateClosed     Time         `json:"date_closed,omitempty"`
	ExpirationDate Time         `json:"expiration_date,omitempty"`
	LastUpdated    Time         `json:"last_updated,omitempty"`
	OrderItems     []OrderItem  `json:"order_items"`
	TotalAmount    float64      `json:"total_amount"`
	PaidAmount     float64      `json:"paid_amount,omitempty"`
	CurrencyID     string       `json:"currency_id"`
	Buyer          *OrderParty  `json:"buyer,omitempty"`
	Seller         *OrderParty  `json:"seller,omitempty"`
	Payments       []OrderPayment `json:"payments,omitempty"`
	Feedback       *OrderFeedback `json:"feedback,omitempty"`
	Context        *OrderContext  `json:"context,omitempty"`
	Shipping       *OrderShipping `json:"shipping,omitempty"`
	Tags           []string       `json:"tags,omitempty"`
}

// OrderItem is one purchased line of an order.
type OrderItem struct {
	Item          OrderLineItem `json:"item"`
	Quantity      int           `json:"quantity"`
	UnitPrice     float64       `json:"unit_price"`
	FullUnitPrice float64       `json:"full_unit_price,omitempty"`
	GrossPrice    float64       `json:"gross_price,omitempty"`
	SaleFee       float64       `json:"sale_fee,omitempty"`
	CurrencyID    string        `json:"currency_id"`
	ListingTypeID string        `json:"listing_type_id,omitempty"`
}

// OrderLineItem identifies the listing within an order item.
type OrderLineItem struct {
	ID                  string      `json:"id"`
	Title               string      `json:"title"`
	CategoryID          string      `json:"category_id,omitempty"`
	VariationID         *int64      `json:"variation_id,omitempty"`
	VariationAttributes []Attribute `json:"variation_attributes,omitempty"`
	SellerCustomField   *string     `json:"seller_custom_field,omitempty"`
}

// OrderParty is the buyer or seller summary on an order.
type OrderParty struct {
	ID        int64  `json:"id"`
	Nickname  string `json:"nickname,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// OrderPayment is a payment associated with an order.
type OrderPayment struct {
	ID                int64   `json:"id"`
	TransactionAmount float64 `json:"transaction_amount"`
	TotalPaidAmount   float64 `json:"total_paid_amount,omitempty"`
	CurrencyID        string  `json:"currency_id"`
	Status            string  `json:"status"`
	StatusDetail      string  `json:"status_detail,omitempty"`
	PaymentMethodID   string  `json:"payment_method_id,omitempty"`
	PaymentType       string  `json:"payment_type,omitempty"`
	DateCreated       Time    `json:"date_created,omitempty"`
	DateLastModified  Time    `json:"date_last_modified,omitempty"`
}

// OrderFeedback holds references to purchase/sale feedback IDs on an order.
type OrderFeedback struct {
	Purchase *int64 `json:"purchase"`
	Sale     *int64 `json:"sale"`
}

// FeedbackRequest is the body of POST /orders/{id}/feedback.
//
// Fulfilled true means the order was completed as expected. Rating values are
// "positive", "neutral", or "negative". Message is a free-text comment
// visible to the counterpart.
type FeedbackRequest struct {
	Fulfilled *bool            `json:"fulfilled,omitempty"`
	Ratings   []FeedbackRating `json:"ratings,omitempty"`
	Message   string           `json:"message,omitempty"`
}

// FeedbackRating is one scored dimension in a feedback submission.
type FeedbackRating struct {
	Type  string `json:"type"`
	Value string `json:"value"` // "positive", "neutral", "negative"
}

// FeedbackResponse is returned by a successful feedback submission.
type FeedbackResponse struct {
	ID          int64  `json:"id"`
	OrderID     int64  `json:"order_id"`
	Fulfilled   bool   `json:"fulfilled"`
	Status      string `json:"status"`
	DateCreated Time   `json:"date_created,omitempty"`
}

// OrderNoteRequest is the body of POST /orders/{id}/notes (seller-only note).
type OrderNoteRequest struct {
	Note string `json:"note"`
}

// OrderContext describes how the order originated.
type OrderContext struct {
	Channel string   `json:"channel,omitempty"`
	Site    string   `json:"site,omitempty"`
	Flows   []string `json:"flows,omitempty"`
}

// OrderShipping references the shipment of an order.
type OrderShipping struct {
	ID int64 `json:"id"`
}

// OrderSearchResponse is the response of GET /orders/search.
type OrderSearchResponse struct {
	Query          string  `json:"query,omitempty"`
	Results        []Order `json:"results"`
	Paging         Paging  `json:"paging"`
	Sort           *Sort   `json:"sort,omitempty"`
	AvailableSorts []Sort  `json:"available_sorts,omitempty"`
}

// OrdersService accesses sales orders.
type OrdersService struct{ c *Client }

// Get returns a single order (GET /orders/{id}).
func (s *OrdersService) Get(ctx context.Context, orderID int64) (*Order, error) {
	return Get[*Order](ctx, s.c, EPOrderByID, orderID)
}

// Search lists orders (GET /orders/search). Typical params: seller, buyer,
// order.status, order.date_created.from/to, q, sort.
func (s *OrdersService) Search(ctx context.Context, params url.Values) (*OrderSearchResponse, error) {
	return GetQ[*OrderSearchResponse](ctx, s.c, EPOrdersSearch, params)
}

// SearchBySeller is a convenience for the common "orders of a seller" query.
func (s *OrdersService) SearchBySeller(ctx context.Context, sellerID int64, opts ListOptions, extra url.Values) (*OrderSearchResponse, error) {
	q := mergeValues(opts.Values(), extra)
	q.Set("seller", itoa(sellerID))
	return s.Search(ctx, q)
}

// PostFeedback submits seller/buyer feedback for an order
// (POST /orders/{id}/feedback).
func (s *OrdersService) PostFeedback(ctx context.Context, orderID int64, req FeedbackRequest) (*FeedbackResponse, error) {
	return Post[*FeedbackResponse, FeedbackRequest](ctx, s.c, EPOrderFeedback, req, orderID)
}

// AddNote attaches a private seller note to an order
// (POST /orders/{id}/notes). Notes are not visible to the buyer.
func (s *OrdersService) AddNote(ctx context.Context, orderID int64, note string) error {
	_, err := Post[map[string]any, OrderNoteRequest](ctx, s.c, EPOrderNotes, OrderNoteRequest{Note: note}, orderID)
	return err
}
