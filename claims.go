package mercadolibre

import (
	"context"
	"net/url"
)

// Claim is a post-purchase claim/return/change/mediation
// (GET /post-purchase/v1/claims/{id}).
type Claim struct {
	ID              int64            `json:"id"`
	ResourceID      int64            `json:"resource_id"`
	Resource        string           `json:"resource"` // shipment | payment | order | purchase
	Status          string           `json:"status"`   // opened | closed
	Type            string           `json:"type"`     // mediations | returns | change | ...
	Stage           string           `json:"stage"`    // claim | dispute | recontact | ...
	ParentID        *int64           `json:"parent_id"`
	ReasonID        string           `json:"reason_id"`
	Fulfilled       bool             `json:"fulfilled"`
	QuantityType    *string          `json:"quantity_type"`
	ClaimedQuantity int              `json:"claimed_quantity,omitempty"`
	ClaimVersion    float64          `json:"claim_version,omitempty"`
	Players         []ClaimPlayer    `json:"players"`
	Resolution      *ClaimResolution `json:"resolution"`
	SiteID          string           `json:"site_id"`
	DateCreated     Time             `json:"date_created"`
	LastUpdated     Time             `json:"last_updated"`
	RelatedEntities []string         `json:"related_entities,omitempty"`
}

// ClaimPlayer is a participant in a claim and the actions available to them.
type ClaimPlayer struct {
	Role             string         `json:"role"` // complainant | respondent
	Type             string         `json:"type"` // buyer | seller | ...
	UserID           int64          `json:"user_id"`
	AvailableActions []ClaimAction  `json:"available_actions,omitempty"`
}

// ClaimAction is an action a player may take.
type ClaimAction struct {
	Action    string  `json:"action"`
	Mandatory bool    `json:"mandatory"`
	DueDate   *string `json:"due_date"`
}

// ClaimResolution is the outcome of a closed claim.
type ClaimResolution struct {
	Reason         string   `json:"reason"`
	DateCreated    Time     `json:"date_created"`
	Benefited      []string `json:"benefited"`
	ClosedBy       string   `json:"closed_by"`
	AppliedCoverage bool    `json:"applied_coverage"`
}

// ClaimSearchResponse is the response of /post-purchase/v1/claims/search.
type ClaimSearchResponse struct {
	Paging Paging  `json:"paging"`
	Data   []Claim `json:"data"`
}

// ClaimsService accesses claims, returns and changes.
type ClaimsService struct{ c *Client }

// Get returns a claim by ID (GET /post-purchase/v1/claims/{id}).
func (s *ClaimsService) Get(ctx context.Context, claimID int64) (*Claim, error) {
	return Get[*Claim](ctx, s.c, EPClaimByID, claimID)
}

// Search queries claims (GET /post-purchase/v1/claims/search). At least one real
// filter is required (e.g. players.user_id + players.role, resource +
// resource_id, order_id, pack_id, status, type, stage). offset/limit/sort alone
// return HTTP 400.
func (s *ClaimsService) Search(ctx context.Context, params url.Values) (*ClaimSearchResponse, error) {
	return GetQ[*ClaimSearchResponse](ctx, s.c, EPClaimsSearch, params)
}

// SearchByPlayer is a convenience for the common "claims where user has role"
// query (e.g. role "respondent" for a seller).
func (s *ClaimsService) SearchByPlayer(ctx context.Context, userID int64, role, status string, opts ListOptions) (*ClaimSearchResponse, error) {
	q := opts.Values()
	q.Set("players.user_id", itoa(userID))
	q.Set("players.role", role)
	if status != "" {
		q.Set("status", status)
	}
	return s.Search(ctx, q)
}

// ClaimHistoryEntry is one event in a claim's history
// (GET /post-purchase/v1/claims/{id}/history).
type ClaimHistoryEntry struct {
	Date    Time   `json:"date"`
	Action  string `json:"action,omitempty"`
	Role    string `json:"role,omitempty"`
	UserID  int64  `json:"user_id,omitempty"`
	Detail  string `json:"detail,omitempty"`
	NewData any    `json:"new_data,omitempty"`
}

// TakeActionRequest is the body of POST /post-purchase/v1/claims/{id}/actions.
type TakeActionRequest struct {
	Action string `json:"action"`
	SiteID string `json:"site_id,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// History returns the event history of a claim
// (GET /post-purchase/v1/claims/{id}/history).
func (s *ClaimsService) History(ctx context.Context, claimID int64) ([]ClaimHistoryEntry, error) {
	return Get[[]ClaimHistoryEntry](ctx, s.c, EPClaimHistory, claimID)
}

// TakeAction submits an action on a claim
// (POST /post-purchase/v1/claims/{id}/actions).
// Typical actions: "RETURN", "REFUND", "NOT_RECEIVED", "CLOSED_BY_BUYER".
// Check ClaimPlayer.AvailableActions for what is allowed in the current state.
func (s *ClaimsService) TakeAction(ctx context.Context, claimID int64, req TakeActionRequest) (map[string]any, error) {
	return Post[map[string]any, TakeActionRequest](ctx, s.c, EPClaimActions, req, claimID)
}
