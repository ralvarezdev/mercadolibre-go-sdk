package mercadolibre

import (
	"context"
	"net/url"
	"strings"
)

// Shipment is a shipment (GET /shipments/{id}). Pass the x-format-new header on
// some related sub-resources per the docs; the core resource is modeled here.
type Shipment struct {
	ID             int64            `json:"id"`
	Status         string           `json:"status"`
	Substatus      *string          `json:"substatus,omitempty"`
	Mode           string           `json:"mode,omitempty"`
	LogisticType   string           `json:"logistic_type,omitempty"`
	OrderID        int64            `json:"order_id,omitempty"`
	TrackingNumber string           `json:"tracking_number,omitempty"`
	TrackingMethod string           `json:"tracking_method,omitempty"`
	DateCreated    Time             `json:"date_created,omitempty"`
	LastUpdated    Time             `json:"last_updated,omitempty"`
	ReceiverAddress *ShipmentAddress `json:"receiver_address,omitempty"`
	SenderAddress   *ShipmentAddress `json:"sender_address,omitempty"`
	ShippingItems   []ShipmentItem   `json:"shipping_items,omitempty"`
	StatusHistory   *ShipmentStatusHistory `json:"status_history,omitempty"`
}

// ShipmentAddress is a sender/receiver address on a shipment.
type ShipmentAddress struct {
	ID           int64   `json:"id,omitempty"`
	AddressLine  string  `json:"address_line,omitempty"`
	StreetName   string  `json:"street_name,omitempty"`
	StreetNumber string  `json:"street_number,omitempty"`
	ZipCode      string  `json:"zip_code,omitempty"`
	City         *IDName `json:"city,omitempty"`
	State        *IDName `json:"state,omitempty"`
	Country      *IDName `json:"country,omitempty"`
	ReceiverName string  `json:"receiver_name,omitempty"`
	ReceiverPhone string `json:"receiver_phone,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
}

// ShipmentItem is an item included in a shipment.
type ShipmentItem struct {
	ID          string `json:"id"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description,omitempty"`
}

// ShipmentStatusHistory is the embedded latest status/substatus snapshot.
type ShipmentStatusHistory struct {
	Status    string `json:"status,omitempty"`
	Substatus string `json:"substatus,omitempty"`
}

// ShipmentHistoryEntry is one entry of GET /shipments/{id}/history.
type ShipmentHistoryEntry struct {
	Status    string `json:"status"`
	Substatus string `json:"substatus"`
	Date      Time   `json:"date"`
}

// ShipmentPayment is one entry of GET /shipments/{id}/payments.
type ShipmentPayment struct {
	PaymentID int64   `json:"payment_id"`
	UserID    int64   `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
}

// ShipmentsService accesses shipments and their sub-resources.
type ShipmentsService struct{ c *Client }

// Get returns a shipment (GET /shipments/{id}).
func (s *ShipmentsService) Get(ctx context.Context, shipmentID int64) (*Shipment, error) {
	return Get[*Shipment](ctx, s.c, EPShipmentByID, shipmentID)
}

// Items returns the items in a shipment (GET /shipments/{id}/items).
func (s *ShipmentsService) Items(ctx context.Context, shipmentID int64) ([]ShipmentItem, error) {
	return Get[[]ShipmentItem](ctx, s.c, EPShipmentItems, shipmentID)
}

// History returns the status/substatus history (GET /shipments/{id}/history).
func (s *ShipmentsService) History(ctx context.Context, shipmentID int64) ([]ShipmentHistoryEntry, error) {
	return Get[[]ShipmentHistoryEntry](ctx, s.c, EPShipmentHistory, shipmentID)
}

// Payments returns the payments of a shipment (GET /shipments/{id}/payments).
func (s *ShipmentsService) Payments(ctx context.Context, shipmentID int64) ([]ShipmentPayment, error) {
	return Get[[]ShipmentPayment](ctx, s.c, EPShipmentPayments, shipmentID)
}

// Costs returns the shipping costs breakdown (GET /shipments/{id}/costs).
func (s *ShipmentsService) Costs(ctx context.Context, shipmentID int64) (map[string]any, error) {
	return Get[map[string]any](ctx, s.c, EPShipmentCosts, shipmentID)
}

// Carrier returns the carrier information for a shipment
// (GET /shipments/{id}/carrier). The shape varies by logistic_type; decoded as
// map for flexibility.
func (s *ShipmentsService) Carrier(ctx context.Context, shipmentID int64) (map[string]any, error) {
	return Get[map[string]any](ctx, s.c, EPShipmentCarrier, shipmentID)
}

// SLA returns the service-level agreement data for a shipment
// (GET /shipments/{id}/sla).
func (s *ShipmentsService) SLA(ctx context.Context, shipmentID int64) (map[string]any, error) {
	return Get[map[string]any](ctx, s.c, EPShipmentSLA, shipmentID)
}

// LeadTime returns the lead-time details for a shipment
// (GET /shipments/{id}/lead_time).
func (s *ShipmentsService) LeadTime(ctx context.Context, shipmentID int64) (map[string]any, error) {
	return Get[map[string]any](ctx, s.c, EPShipmentLeadTime, shipmentID)
}

// Labels returns the raw label PDF for the given shipment IDs
// (GET /shipment_labels?shipment_ids=id1,id2,...).
// The response is binary (typically application/pdf).
func (s *ShipmentsService) Labels(ctx context.Context, shipmentIDs ...int64) ([]byte, error) {
	ids := make([]string, len(shipmentIDs))
	for i, id := range shipmentIDs {
		ids[i] = itoa(id)
	}
	q := url.Values{"shipment_ids": {strings.Join(ids, ",")}}
	return GetRaw(ctx, s.c, EPShipmentLabels, q)
}
