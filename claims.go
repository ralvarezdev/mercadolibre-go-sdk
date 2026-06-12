package mercadolibre

import (
	"context"
	"net/url"
)

type (
	// Claim is a post-purchase claim/return/change/mediation
	// (GET /post-purchase/v1/claims/{id}).
	Claim struct {
		DateCreated     Time             `json:"date_created"`
		LastUpdated     Time             `json:"last_updated"`
		ParentID        *int64           `json:"parent_id"`
		QuantityType    *string          `json:"quantity_type"`
		Resolution      *ClaimResolution `json:"resolution"`
		Type            string           `json:"type"`
		Resource        string           `json:"resource"`
		Status          string           `json:"status"`
		Stage           string           `json:"stage"`
		ReasonID        string           `json:"reason_id"`
		SiteID          string           `json:"site_id"`
		RelatedEntities []string         `json:"related_entities,omitempty"`
		Players         []ClaimPlayer    `json:"players"`
		ClaimVersion    float64          `json:"claim_version,omitempty"`
		ID              int64            `json:"id"`
		ResourceID      int64            `json:"resource_id"`
		ClaimedQuantity int              `json:"claimed_quantity,omitempty"`
		Fulfilled       bool             `json:"fulfilled"`
	}

	// ClaimPlayer is a participant in a claim and the actions available to them.
	ClaimPlayer struct {
		Role             string        `json:"role"`
		Type             string        `json:"type"`
		AvailableActions []ClaimAction `json:"available_actions,omitempty"`
		UserID           int64         `json:"user_id"`
	}

	// ClaimAction is an action a player may take.
	ClaimAction struct {
		DueDate   *string `json:"due_date"`
		Action    string  `json:"action"`
		Mandatory bool    `json:"mandatory"`
	}

	// ClaimResolution is the outcome of a closed claim.
	ClaimResolution struct {
		DateCreated     Time     `json:"date_created"`
		Reason          string   `json:"reason"`
		ClosedBy        string   `json:"closed_by"`
		Benefited       []string `json:"benefited"`
		AppliedCoverage bool     `json:"applied_coverage"`
	}

	// ClaimSearchResponse is the response of /post-purchase/v1/claims/search.
	ClaimSearchResponse struct {
		Data   []Claim `json:"data"`
		Paging Paging  `json:"paging"`
	}

	// ClaimHistoryEntry is one event in a claim's history
	// (GET /post-purchase/v1/claims/{id}/history).
	ClaimHistoryEntry struct {
		Date    Time   `json:"date"`
		NewData any    `json:"new_data,omitempty"`
		Action  string `json:"action,omitempty"`
		Role    string `json:"role,omitempty"`
		Detail  string `json:"detail,omitempty"`
		UserID  int64  `json:"user_id,omitempty"`
	}

	// TakeActionRequest is the body of POST /post-purchase/v1/claims/{id}/actions.
	TakeActionRequest struct {
		Action string `json:"action"`
		SiteID string `json:"site_id,omitempty"`
		Reason string `json:"reason,omitempty"`
	}

	// ClaimsService accesses claims, returns and changes.
	ClaimsService struct{ c *Client }
)

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
func (s *ClaimsService) SearchByPlayer(
	ctx context.Context,
	userID int64,
	role, status string,
	opts ListOptions,
) (*ClaimSearchResponse, error) {
	q := opts.Values()
	q.Set(string(QueryParamPlayersUserID), itoa(userID))
	q.Set(string(QueryParamPlayersRole), role)
	if status != "" {
		q.Set(string(QueryParamStatus), status)
	}
	return s.Search(ctx, q)
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
