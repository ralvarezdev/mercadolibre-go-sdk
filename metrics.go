package mercadolibre

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

type (
	// VisitsDetail breaks visits down by company/source.
	VisitsDetail struct {
		Company  string `json:"company"`
		Quantity int    `json:"quantity"`
	}

	// ItemVisits is total item visits over a date range (GET /items/visits).
	ItemVisits struct {
		DateFrom     Time           `json:"date_from"`
		DateTo       Time           `json:"date_to"`
		ItemID       string         `json:"item_id"`
		VisitsDetail []VisitsDetail `json:"visits_detail"`
		TotalVisits  int            `json:"total_visits"`
	}

	// UserVisits is total visits across a user's items over a date range
	// (GET /users/{id}/items_visits).
	UserVisits struct {
		DateFrom     Time           `json:"date_from"`
		DateTo       Time           `json:"date_to"`
		VisitsDetail []VisitsDetail `json:"visits_detail"`
		UserID       int64          `json:"user_id"`
		TotalVisits  int            `json:"total_visits"`
	}

	// VisitsBucket is one time-interval bucket in a time-window visits response.
	VisitsBucket struct {
		Date         Time           `json:"date"`
		VisitsDetail []VisitsDetail `json:"visits_detail"`
		Total        int            `json:"total"`
	}

	// TimeWindowVisits is the response of the visits/time_window endpoints.
	TimeWindowVisits struct {
		DateFrom    Time           `json:"date_from"`
		DateTo      Time           `json:"date_to"`
		ItemID      string         `json:"item_id,omitempty"`
		Unit        string         `json:"unit"`
		Results     []VisitsBucket `json:"results"`
		UserID      int64          `json:"user_id,omitempty"`
		TotalVisits int            `json:"total_visits"`
		Last        int            `json:"last"`
	}

	// Trend is a trending search keyword (GET /trends/{site}[/{category}]).
	Trend struct {
		Keyword string `json:"keyword"`
		URL     string `json:"url"`
	}

	// MetricsService accesses visits, trends and seller reputation.
	MetricsService struct{ c *Client }
)

const dateOnly = "2006-01-02"

// ItemVisits returns an item's total visits between two dates
// (GET /items/visits?ids={id}&date_from=&date_to=).
func (s *MetricsService) ItemVisits(ctx context.Context, itemID string, from, to time.Time) (*ItemVisits, error) {
	q := url.Values{
		string(QueryParamIDs):      {itemID},
		string(QueryParamDateFrom): {from.Format(dateOnly)},
		string(QueryParamDateTo):   {to.Format(dateOnly)},
	}
	return GetQ[*ItemVisits](ctx, s.c, EPItemsVisits, q)
}

// UserVisits returns total visits to a user's items between two dates
// (GET /users/{id}/items_visits?date_from=&date_to=).
func (s *MetricsService) UserVisits(ctx context.Context, userID int64, from, to time.Time) (*UserVisits, error) {
	q := url.Values{
		string(QueryParamDateFrom): {from.Format(dateOnly)},
		string(QueryParamDateTo):   {to.Format(dateOnly)},
	}
	return GetQ[*UserVisits](ctx, s.c, EPUserItemsVisits, q, userID)
}

// ItemVisitsTimeWindow returns an item's visits bucketed over the last N units
// (GET /items/{id}/visits/time_window?last=&unit=&ending=). unit is e.g. "day".
func (s *MetricsService) ItemVisitsTimeWindow(
	ctx context.Context,
	itemID string,
	last int,
	unit string,
	ending time.Time,
) (*TimeWindowVisits, error) {
	q := url.Values{
		string(QueryParamLast): {strconv.Itoa(last)},
		string(QueryParamUnit): {unit},
	}
	if !ending.IsZero() {
		q.Set(string(QueryParamEnding), ending.Format(dateOnly))
	}
	return GetQ[*TimeWindowVisits](ctx, s.c, EPItemVisitsTimeWindow, q, itemID)
}

// SiteTrends returns trending keywords for a site (GET /trends/{site}).
func (s *MetricsService) SiteTrends(ctx context.Context, siteID string) ([]Trend, error) {
	return Get[[]Trend](ctx, s.c, EPTrendsSite, siteID)
}

// CategoryTrends returns trending keywords for a site category
// (GET /trends/{site}/{category}).
func (s *MetricsService) CategoryTrends(ctx context.Context, siteID, categoryID string) ([]Trend, error) {
	return Get[[]Trend](ctx, s.c, EPTrendsCategory, siteID, categoryID)
}

// Reputation returns a seller's reputation by reading the user resource
// (GET /users/{id}).
func (s *MetricsService) Reputation(ctx context.Context, userID int64) (*SellerReputation, error) {
	u, err := s.c.Users.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return u.SellerReputation, nil
}
