package mercadolibre

import (
	"context"
	"net/url"
	"strconv"
)

// MissedFeed is one entry of the notifications history
// (GET /missed_feeds?app_id=...). It records a webhook the platform attempted to
// deliver to your callback.
type MissedFeed struct {
	ID            string `json:"_id"`
	Resource      string `json:"resource"`
	UserID        int64  `json:"user_id"`
	Topic         string `json:"topic"`
	ApplicationID int64  `json:"application_id"`
	Attempts      int    `json:"attempts"`
	Sent          Time   `json:"sent"`
	Received      Time   `json:"received"`
}

type missedFeedsResponse struct {
	Messages []MissedFeed `json:"messages"`
}

// MissedFeedsOptions filters the notifications history.
type MissedFeedsOptions struct {
	AppID  int64
	Topic  string // optional, e.g. "payments"
	Offset int
	Limit  int
}

// MissedFeeds returns the notifications the platform tried to deliver to your
// callback (GET /missed_feeds). Use it to recover from missed webhooks.
func (c *Client) MissedFeeds(ctx context.Context, opts MissedFeedsOptions) ([]MissedFeed, error) {
	q := url.Values{"app_id": {strconv.FormatInt(opts.AppID, 10)}}
	if opts.Topic != "" {
		q.Set("topic", opts.Topic)
	}
	if opts.Offset > 0 {
		q.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		q.Set("limit", strconv.Itoa(opts.Limit))
	}
	resp, err := GetQ[missedFeedsResponse](ctx, c, EPMissedFeeds, q)
	if err != nil {
		return nil, err
	}
	return resp.Messages, nil
}
