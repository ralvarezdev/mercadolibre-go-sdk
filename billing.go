package mercadolibre

import "context"

// BillingInfoResponse is the response of
// GET /orders/billing-info/{site}/{billing_info_id}.
//
// Note: the `invoice_type` field was removed from this API; the integrator must
// derive the invoice type from the buyer's identification document (for MLA:
// CUIT → Factura A, DNI → Factura B).
type BillingInfoResponse struct {
	SiteID        string       `json:"site_id,omitempty"`
	Buyer         BillingParty `json:"buyer"`
	Seller        BillingParty `json:"seller"`
	BuyerID       int64        `json:"buyer_id,omitempty"`
	SellerID      int64        `json:"seller_id,omitempty"`
	IsOrderOrigin bool         `json:"is_order_origin,omitempty"`
}

// BillingParty is a buyer or seller in a billing-info response.
type BillingParty struct {
	CustID      FlexID       `json:"cust_id"`
	BillingInfo *BillingInfo `json:"billing_info,omitempty"`
}

// BillingInfo is the fiscal data used to invoice a buyer.
type BillingInfo struct {
	Name           string          `json:"name,omitempty"`
	LastName       string          `json:"last_name,omitempty"`
	Identification *Identification `json:"identification,omitempty"`
	Taxes          *BillingTaxes   `json:"taxes,omitempty"`
	Address        *BillingAddress `json:"address,omitempty"`
	Attributes     map[string]any  `json:"attributes,omitempty"`
}

// BillingTaxes holds tax classification for billing.
type BillingTaxes struct {
	TaxpayerType     *IDDescription `json:"taxpayer_type,omitempty"`
	EconomicActivity string         `json:"economic_activity,omitempty"`
}

// IDDescription is an {id, description} pair (e.g. taxpayer type).
type IDDescription struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

// BillingAddress is the fiscal address for billing.
type BillingAddress struct {
	StreetName   string        `json:"street_name,omitempty"`
	StreetNumber string        `json:"street_number,omitempty"`
	CityName     string        `json:"city_name,omitempty"`
	City         string        `json:"city,omitempty"`
	Neighborhood string        `json:"neighborhood,omitempty"`
	Comment      string        `json:"comment,omitempty"`
	State        *BillingState `json:"state,omitempty"`
	ZipCode      string        `json:"zip_code,omitempty"`
	CountryID    string        `json:"country_id,omitempty"`
}

// BillingState is a state with an optional fiscal code.
type BillingState struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// BillingReport is one entry in a billing report list (provisions or
// perceptions). The fields returned vary by site and report type.
type BillingReport struct {
	ID            string  `json:"id"`
	Type          string  `json:"type,omitempty"`
	Status        string  `json:"status,omitempty"`
	DateCreated   Time    `json:"date_created,omitempty"`
	DateFrom      Time    `json:"date_from,omitempty"`
	DateTo        Time    `json:"date_to,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	CurrencyID    string  `json:"currency_id,omitempty"`
	DownloadURL   string  `json:"download_url,omitempty"`
	GeneratedDate Time    `json:"generated_date,omitempty"`
}

// BillingReportsResponse is the response of a billing report list.
type BillingReportsResponse struct {
	Results []BillingReport `json:"results"`
	Paging  Paging          `json:"paging"`
}

// BillingService accesses invoicing data.
type BillingService struct{ c *Client }

// Info returns the billing info for an order's billing_info_id on a site
// (GET /orders/billing-info/{site}/{billing_info_id}). The billing_info_id is
// obtained from the order resource.
func (s *BillingService) Info(ctx context.Context, siteID, billingInfoID string) (*BillingInfoResponse, error) {
	return Get[*BillingInfoResponse](ctx, s.c, EPOrderBillingInfo, siteID, billingInfoID)
}

// Reports lists billing reports of the given type for a user
// (GET /billing/integration/reports/{report_type}?user_id=&version=).
//
// Common report_type values: "provisions", "perceptions".
// version is the report schema version (e.g. "v1_0").
func (s *BillingService) Reports(ctx context.Context, reportType string, userID int64, version string, opts ListOptions) (*BillingReportsResponse, error) {
	q := opts.Values()
	q.Set("user_id", itoa(userID))
	if version != "" {
		q.Set("version", version)
	}
	return GetQ[*BillingReportsResponse](ctx, s.c, EPBillingReports, q, reportType)
}

// DownloadReport downloads the raw content of a billing report document
// (GET /billing/integration/reports/{report_type}/download/{doc_id}).
//
// Returns the raw bytes of the document (typically CSV or PDF).
func (s *BillingService) DownloadReport(ctx context.Context, reportType, docID string) ([]byte, error) {
	return GetRaw(ctx, s.c, EPBillingDownload, nil, reportType, docID)
}
