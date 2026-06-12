package mercadolibre

import (
	"context"
	"net/url"
	"strconv"
)

type (
	// MissedFeed is one entry of the notifications history
	// (GET /missed_feeds?app_id=...). It records a webhook the platform attempted to
	// deliver to your callback.
	MissedFeed struct {
		Received      Time   `json:"received"`
		Sent          Time   `json:"sent"`
		ID            string `json:"_id"`
		Resource      string `json:"resource"`
		Topic         string `json:"topic"`
		ApplicationID int64  `json:"application_id"`
		UserID        int64  `json:"user_id"`
		Attempts      int    `json:"attempts"`
	}

	missedFeedsResponse struct {
		Messages []MissedFeed `json:"messages"`
	}

	// MissedFeedsOptions filters the notifications history.
	MissedFeedsOptions struct {
		Topic  string // optional, e.g. "payments"
		AppID  int64
		Offset int
		Limit  int
	}
)

// MissedFeeds returns the notifications the platform tried to deliver to your
// callback (GET /missed_feeds). Use it to recover from missed webhooks.
func (c *Client) MissedFeeds(ctx context.Context, opts MissedFeedsOptions) ([]MissedFeed, error) {
	q := url.Values{string(QueryParamAppID): {strconv.FormatInt(opts.AppID, 10)}}
	if opts.Topic != "" {
		q.Set(string(QueryParamTopic), opts.Topic)
	}
	if opts.Offset > 0 {
		q.Set(string(QueryParamOffset), strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		q.Set(string(QueryParamLimit), strconv.Itoa(opts.Limit))
	}
	resp, err := GetQ[missedFeedsResponse](ctx, c, EPMissedFeeds, q)
	if err != nil {
		return nil, err
	}
	return resp.Messages, nil
}
