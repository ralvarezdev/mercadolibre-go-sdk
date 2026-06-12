package mercadolibre

import (
	"context"
	"net/url"
)

type (
	// Order is a sale (GET /orders/{id}).
	Order struct {
		DateClosed     Time           `json:"date_closed,omitempty"`
		DateCreated    Time           `json:"date_created"`
		LastUpdated    Time           `json:"last_updated,omitempty"`
		ExpirationDate Time           `json:"expiration_date,omitempty"`
		Seller         *OrderParty    `json:"seller,omitempty"`
		StatusDetail   *string        `json:"status_detail,omitempty"`
		Context        *OrderContext  `json:"context,omitempty"`
		Shipping       *OrderShipping `json:"shipping,omitempty"`
		Feedback       *OrderFeedback `json:"feedback,omitempty"`
		Buyer          *OrderParty    `json:"buyer,omitempty"`
		CurrencyID     string         `json:"currency_id"`
		Status         string         `json:"status"`
		Payments       []OrderPayment `json:"payments,omitempty"`
		Tags           []string       `json:"tags,omitempty"`
		OrderItems     []OrderItem    `json:"order_items"`
		TotalAmount    float64        `json:"total_amount"`
		PaidAmount     float64        `json:"paid_amount,omitempty"`
		ID             int64          `json:"id"`
	}

	// OrderItem is one purchased line of an order.
	OrderItem struct {
		CurrencyID    string        `json:"currency_id"`
		ListingTypeID string        `json:"listing_type_id,omitempty"`
		Item          OrderLineItem `json:"item"`
		FullUnitPrice float64       `json:"full_unit_price,omitempty"`
		GrossPrice    float64       `json:"gross_price,omitempty"`
		SaleFee       float64       `json:"sale_fee,omitempty"`
		UnitPrice     float64       `json:"unit_price"`
		Quantity      int           `json:"quantity"`
	}

	// OrderLineItem identifies the listing within an order item.
	OrderLineItem struct {
		VariationID         *int64      `json:"variation_id,omitempty"`
		SellerCustomField   *string     `json:"seller_custom_field,omitempty"`
		ID                  string      `json:"id"`
		Title               string      `json:"title"`
		CategoryID          string      `json:"category_id,omitempty"`
		VariationAttributes []Attribute `json:"variation_attributes,omitempty"`
	}

	// OrderParty is the buyer or seller summary on an order.
	OrderParty struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Nickname  string `json:"nickname,omitempty"`
		ID        int64  `json:"id"`
	}

	// OrderPayment is a payment associated with an order.
	OrderPayment struct {
		DateCreated       Time    `json:"date_created,omitempty"`
		DateLastModified  Time    `json:"date_last_modified,omitempty"`
		CurrencyID        string  `json:"currency_id"`
		Status            string  `json:"status"`
		StatusDetail      string  `json:"status_detail,omitempty"`
		PaymentMethodID   string  `json:"payment_method_id,omitempty"`
		PaymentType       string  `json:"payment_type,omitempty"`
		TransactionAmount float64 `json:"transaction_amount"`
		TotalPaidAmount   float64 `json:"total_paid_amount,omitempty"`
		ID                int64   `json:"id"`
	}

	// OrderFeedback holds references to purchase/sale feedback IDs on an order.
	OrderFeedback struct {
		Sale     *int64 `json:"sale"`
		Purchase *int64 `json:"purchase"`
	}

	// FeedbackRequest is the body of POST /orders/{id}/feedback.
	//
	// Fulfilled true means the order was completed as expected. Rating values are
	// "positive", "neutral", or "negative". Message is a free-text comment
	// visible to the counterpart.
	FeedbackRequest struct {
		Fulfilled *bool            `json:"fulfilled,omitempty"`
		Message   string           `json:"message,omitempty"`
		Ratings   []FeedbackRating `json:"ratings,omitempty"`
	}

	// FeedbackRating is one scored dimension in a feedback submission.
	FeedbackRating struct {
		Type  string `json:"type"`
		Value string `json:"value"` // "positive", "neutral", "negative"
	}

	// FeedbackResponse is returned by a successful feedback submission.
	FeedbackResponse struct {
		DateCreated Time   `json:"date_created,omitempty"`
		Status      string `json:"status"`
		ID          int64  `json:"id"`
		OrderID     int64  `json:"order_id"`
		Fulfilled   bool   `json:"fulfilled"`
	}

	// OrderNoteRequest is the body of POST /orders/{id}/notes (seller-only note).
	OrderNoteRequest struct {
		Note string `json:"note"`
	}

	// OrderContext describes how the order originated.
	OrderContext struct {
		Site    string   `json:"site,omitempty"`
		Channel string   `json:"channel,omitempty"`
		Flows   []string `json:"flows,omitempty"`
	}

	// OrderShipping references the shipment of an order.
	OrderShipping struct {
		ID int64 `json:"id"`
	}

	// OrderSearchResponse is the response of GET /orders/search.
	OrderSearchResponse struct {
		Sort           *Sort   `json:"sort,omitempty"`
		Query          string  `json:"query,omitempty"`
		AvailableSorts []Sort  `json:"available_sorts,omitempty"`
		Results        []Order `json:"results"`
		Paging         Paging  `json:"paging"`
	}

	// OrdersService accesses sales orders.
	OrdersService struct{ c *Client }
)

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
func (s *OrdersService) SearchBySeller(
	ctx context.Context,
	sellerID int64,
	opts ListOptions,
	extra url.Values,
) (*OrderSearchResponse, error) {
	q := mergeValues(opts.Values(), extra)
	q.Set(string(QueryParamSeller), itoa(sellerID))
	return s.Search(ctx, q)
}

// PostFeedback submits seller/buyer feedback for an order
// (POST /orders/{id}/feedback).
func (s *OrdersService) PostFeedback(
	ctx context.Context,
	orderID int64,
	req FeedbackRequest,
) (*FeedbackResponse, error) {
	return Post[*FeedbackResponse, FeedbackRequest](ctx, s.c, EPOrderFeedback, req, orderID)
}

// AddNote attaches a private seller note to an order
// (POST /orders/{id}/notes). Notes are not visible to the buyer.
func (s *OrdersService) AddNote(ctx context.Context, orderID int64, note string) error {
	_, err := Post[map[string]any, OrderNoteRequest](ctx, s.c, EPOrderNotes, OrderNoteRequest{Note: note}, orderID)
	return err
}
