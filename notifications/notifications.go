// Package notifications models MercadoLibre webhook notifications and provides
// an http.Handler to receive them.
//
// MercadoLibre POSTs a small JSON payload to your application's callback URL for
// each subscribed topic. Your endpoint must respond quickly with 2xx; otherwise
// the platform retries and the event ends up in the missed_feeds history.
package notifications

import (
	"encoding/json"
	"net/http"
	"time"
)

// Topic identifies the kind of event a notification refers to.
type Topic string

const (
	TopicOrders         Topic = "orders"
	TopicOrdersV2       Topic = "orders_v2"
	TopicOrdersFeedback Topic = "orders_feedback"
	TopicItems          Topic = "items"
	TopicQuestions      Topic = "questions"
	TopicMessages       Topic = "messages"
	TopicPayments       Topic = "payments"
	TopicShipments      Topic = "shipments"
	TopicPrices         Topic = "prices"
	TopicCatalog        Topic = "catalog_item_competition_status"
	TopicPromotions     Topic = "public_offers"
	TopicVISLeads       Topic = "vis_leads"
	TopicPostPurchase   Topic = "post_purchase"
	TopicStock          Topic = "stock-locations"
)

// Notification is the webhook payload pushed to your callback URL.
type Notification struct {
	Sent          time.Time `json:"sent"`
	Received      time.Time `json:"received"`
	Resource      string    `json:"resource"`
	Topic         Topic     `json:"topic"`
	UserID        int64     `json:"user_id"`
	ApplicationID int64     `json:"application_id"`
	Attempts      int       `json:"attempts"`
}

// Parse decodes a notification from a JSON body.
func Parse(body []byte) (Notification, error) {
	var n Notification
	err := json.Unmarshal(body, &n)
	return n, err
}

// Handler returns an http.Handler that decodes incoming notifications and passes
// each to fn. It responds 200 on success. If fn returns an error the handler
// responds 500 so MercadoLibre retries; return nil once the event is safely
// enqueued/persisted. For lowest latency, have fn enqueue and return quickly
// rather than processing inline.
func Handler(fn func(r *http.Request, n Notification) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()
		var n Notification
		if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if err := fn(r, n); err != nil {
			http.Error(w, "processing error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
